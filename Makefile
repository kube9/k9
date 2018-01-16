include Makefile.defs

all:
	@make build

build:
	@make k9server-build
	@make k9ctl-build

clean:
	@make k9server-clean
	@make k9ctl-clean
	rm -Rf ./bin

k9server-gen-openapi:
	rm -Rf ./$(K9SERVER_OPENAPI_GEN_DIR)
	@make k9server-gen-openapi-server
	@make k9server-gen-openapi-client

k9server-gen-openapi-server:
	mkdir -p $(K9SERVER_OPENAPI_GEN_DIR)
	swagger generate server -t $(K9SERVER_OPENAPI_GEN_DIR) -f $(K9SERVER_OPENAPI_SPEC_FILE) -s server --exclude-main

k9server-gen-openapi-client:
	mkdir -p $(K9SERVER_OPENAPI_GEN_DIR)
	swagger generate client -t $(K9SERVER_OPENAPI_GEN_DIR) -f $(K9SERVER_OPENAPI_SPEC_FILE)

k9server-build:
	go build -i $(GOBUILD) -o ./bin/k9server ./cmd/k9server/main.go

k9server-clean:

k9ctl-build:
	go build -i $(GOBUILD) -o ./bin/k9ctl ./cmd/k9ctl/main.go

k9ctl-clean: