.PHONY: build
build:
	GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o flap

.PHONY: certs
certs:
	./scripts/gen_certs.sh
