VERSION := 0.0.1
COMMIT := $(shell git log -1 --format='%H')
COMMIT := $(shell git log -1 --format='%H')
CONTINENT ?= NA
ifeq ($(CONTINENT), NA)
	TOKEN_SYMBOL := prcna
	UTOKEN_SYMBOL := uprcna
	PREFIX := pricena
else ifeq ($(CONTINENT), SA)
	TOKEN_SYMBOL := prcsa
	UTOKEN_SYMBOL := uprcsa
	PREFIX := pricesa
else ifeq ($(CONTINENT), EU)
	TOKEN_SYMBOL := prceu
	UTOKEN_SYMBOL := uprceu
	PREFIX := priceeu
else ifeq ($(CONTINENT), AF)
	TOKEN_SYMBOL := prcaf
	UTOKEN_SYMBOL := uprcaf
	PREFIX := priceaf
else ifeq ($(CONTINENT), AS)
	TOKEN_SYMBOL := prcas
	UTOKEN_SYMBOL := uprcas
	PREFIX := priceas
else ifeq ($(CONTINENT), OC)
	TOKEN_SYMBOL := prcoc
	UTOKEN_SYMBOL := uprcoc
	PREFIX := priceoc
endif
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=cprc \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=cprcd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X github.com/PriceChain/cprc/app.AccountAddressPrefix=$(PREFIX) \
	-X github.com/PriceChain/cprc/app.CoinOneDenom=$(TOKEN_SYMBOL) \
	-X github.com/PriceChain/cprc/app.CoinPrevDenom=$(UTOKEN_SYMBOL)

BUILD_FLAGS := -ldflags '$(ldflags)'
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)

all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/cprcd-manager
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/cprcd

build: go.sum clean
	go build -mod=mod $(BUILD_FLAGS) -o build/cprcd-manager ./cmd/cprcd-manager
	go build -mod=mod $(BUILD_FLAGS) -o build/cprcd ./cmd/cprcd

build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

# devnet

devnet: clean install devnet-prepare devnet-start

devnet-prepare:
	./scripts/prepare-devnet.sh

devnet-start:
	DAEMON_NAME=cprcd DAEMON_HOME=~/.cprcd DAEMON_ALLOW_DOWNLOAD_BINARIES=true DAEMON_RESTART_AFTER_UPGRADE=true \
    cprcd start --pruning="nothing" --inv-check-period 5

# Clean up the build directory
clean:
	rm -rf build/


# Localnet

# Build nodes using Docker
build-docker:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux localnet-stop
	@if ! [ -f build/node0/.cprcd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/cprcd:Z cprc/core testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test --chain-id test ; fi
	./scripts/import-localnet-seeds.sh
	docker-compose up

# Stop testnet
localnet-stop:
	docker-compose down

localnet: clean build-linux build-docker localnet-start

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.3
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)
containerProtoGenAny=$(PROJECT_NAME)-proto-gen-any-$(protoVer)
containerProtoGenSwagger=$(PROJECT_NAME)-proto-gen-swagger-$(protoVer)
containerProtoFmt=$(PROJECT_NAME)-proto-fmt-$(protoVer)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	./scripts/protocgen.sh
	go mod tidy

# This generates the SDK's custom wrapper for google.protobuf.Any. It should only be run manually when needed
proto-gen-js:
	@echo "Generating Protobuf Typescript"
	bash ./scripts/protocgen-js.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protoc-swagger-gen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi


proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master

# Create log files
log-files:
	sudo mkdir -p /var/log/cprcd && sudo touch /var/log/cprcd/cprcd.log && sudo touch /var/log/cprcd/cprcd_error.log
