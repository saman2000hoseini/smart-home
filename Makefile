export APP=smart-home
export LDFLAGS="-w -s"

subscribe:
	go run -ldflags $(LDFLAGS) ./cmd/smart-home subscribe

build:
	CGO_ENABLED=1 go build -ldflags $(LDFLAGS) ./cmd/smart-home

install:
	CGO_ENABLED=1 go install -ldflags $(LDFLAGS) ./cmd/smart-home
