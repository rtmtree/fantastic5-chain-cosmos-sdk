mock-expected-keepers:
	mockgen -source=x/fantasfive/types/expected_keepers.go \
		-package testutil \
		-destination=x/fantasfive/testutil/expected_keepers_mocks.go 

install-protoc-gen-ts:
	mkdir -p scripts/protoc
	cd scripts && npm install
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip -o scripts/protoc/protoc.zip
	cd scripts/protoc && unzip -o protoc.zip
	rm scripts/protoc/protoc.zip
	ln -s $(pwd)/scripts/protoc/bin/protoc /usr/local/bin/protoc

cosmos-version = v0.45.4

download-cosmos-proto:
	mkdir -p proto/cosmos/base/query/v1beta1
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/proto/cosmos/base/query/v1beta1/pagination.proto -o proto/cosmos/base/query/v1beta1/pagination.proto
	mkdir -p proto/google/api
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/annotations.proto -o proto/google/api/annotations.proto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/http.proto -o proto/google/api/http.proto
	mkdir -p proto/gogoproto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/gogoproto/gogo.proto -o proto/gogoproto/gogo.proto

gen-protoc-ts: 
	mkdir -p ./client/src/types/generated/
	ls proto/fantasfive | xargs -I {} protoc \
		--plugin="./scripts/node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_out="./client/src/types/generated" \
		--proto_path="./proto" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
		fantasfive/{}
	
install-and-gen-protoc-ts: download-cosmos-proto install-protoc-gen-ts gen-protoc-ts

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/fantasfived-linux-amd64 ./cmd/fantasfived/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/fantasfived-linux-arm64 ./cmd/fantasfived/main.go

do-checksum-linux:
	cd build && sha256sum \
		fantasfived-linux-amd64 fantasfived-linux-arm64 \
		> fantasfive-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/fantasfived-darwin-amd64 ./cmd/fantasfived/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/fantasfived-darwin-arm64 ./cmd/fantasfived/main.go

build-all: build-linux build-darwin

do-checksum-darwin:
	cd build && sha256sum \
		fantasfived-darwin-amd64 fantasfived-darwin-arm64 \
		> fantasfive-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-with-checksum: build-linux-with-checksum build-darwin-with-checksum
