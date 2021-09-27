export APP=smart-home
export LDFLAGS="-w -s"

mqtt-subscribe:
	go run -ldflags $(LDFLAGS) ./cmd/smart-home mqtt-subscribe

amqp-subscribe:
	go run -ldflags $(LDFLAGS) ./cmd/smart-home amqp-subscribe

coap-server:
	go run -ldflags $(LDFLAGS) ./cmd/smart-home coap-server

http-server:
	go run -ldflags $(LDFLAGS) ./cmd/smart-home http-server

build:
	CGO_ENABLED=1 go build -ldflags $(LDFLAGS) ./cmd/smart-home

install:
	CGO_ENABLED=1 go install -ldflags $(LDFLAGS) ./cmd/smart-home
