FROM golang:1.13-alpine as builder
RUN mkdir -p /net-up
COPY . /net-up
WORKDIR /net-up
RUN apk update && \
    apk add linux-headers musl-dev gcc go libpcap-dev ca-certificates git

RUN go clean && go build --ldflags '-linkmode external -extldflags "-static -s -w"' -v ./

FROM alpine
COPY --from=builder /net-up/net-up ./net-up
CMD ./net-up