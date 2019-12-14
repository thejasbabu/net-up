# Net-up

A program using google-gopacket to set BPF filter and analyse packets at a configured network devise 

## Requirement

1. golang

## Set-up

1. `go get` the project
2. Set the network devise to start intercepting with BPF 
    ```
     export NET_FILTER="tcp and port 3000" 
     export NET_LOGLEVEL=debug 
     export NET_NETWORKDEVISE=lo0
     export NET_OUTPUT=stdout 
    ```
3. Run the executable with sudo access.
    ```
    sudo ./net-up
    ```

Currently the program just outputs the packet details to stdout, ie `NET_OUTPUT` supports `stdout`. It needs to be extended to support writing to file, etc.  