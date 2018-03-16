.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/Sberex/go-sberex/cmd/swarm
	go build -o bin/geth     github.com/Sberex/go-sberex/cmd/geth
	go build -o bin/bootnode github.com/Sberex/go-sberex/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/Sberex/go-sberex/swarm/...
