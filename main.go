package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

type Config struct {
	LogLevel string `default:"debug"`
	NetworkDevise string `default:"eth0"`
	Filter string `required:"true"`
	Output string `default:"stdout"`
}

func getLogLevel(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
		case "debug": return zapcore.DebugLevel
		case "info": return zapcore.InfoLevel
		case "warning": return zapcore.WarnLevel
		case "error": return zapcore.ErrorLevel
		default: return zapcore.InfoLevel
	}
}

func initLogger(config Config) {
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(getLogLevel(config.LogLevel)),
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{MessageKey: "message", LevelKey: "level", EncodeLevel: zapcore.LowercaseLevelEncoder},
	}.Build()
	zap.ReplaceGlobals(logger)
}

func main() {
	var config Config
	err := envconfig.Process("NET", &config)
	if err != nil {
		log.Fatalf("error reading environment variable: %s", err.Error())
	}
	initLogger(config)
	device := config.NetworkDevise
	zap.S().Infof("Net-up is up and ready to capture packets at devise: %s", device)

	handle, err := pcap.OpenLive(device, 1024, false, pcap.BlockForever)
	if err != nil {
		zap.S().Errorf("error opening handler to devise: %s", err.Error())
		os.Exit(1)
	}
	filter := config.Filter
	zap.S().Infof("Setting up BPF filter: %s", filter)

	err = handle.SetBPFFilter(filter)
	if err != nil {
		zap.S().Errorf("Error applying filter: %s", err.Error())
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	mode := config.Output
	switch strings.ToLower(mode) {
		case "stdout": outputToStdOutput(packetSource)
		default:
			zap.S().Fatalf("Invalid output type, currently supported `stdout`")
	}
}

func outputToStdOutput(source *gopacket.PacketSource) {
	for packet := range source.Packets() {
		zap.S().Info(packet.String())
	}

}
