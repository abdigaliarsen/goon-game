PROTO_DIRS = pkg/proto/wikipedia
PACKAGE_BASE_DIR = github.com/abdigaliarsen/goon-game/pkg/proto

generate-structs:
	@for DIR in $(PROTO_DIRS); do \
		echo "Generating structs in $$DIR..."; \
		mkdir -p $$DIR; \
		PROTO_FILE=$$(basename $$DIR).proto; \
		protoc --go_out=. --go_opt=paths=source_relative \
			$$DIR/$$PROTO_FILE || exit 1; \
	done

generate:
	@for DIR in $(PROTO_DIRS); do \
		echo "Generating code in $$DIR..."; \
		mkdir -p $$DIR; \
		PROTO_FILE=$$(basename $$DIR).proto; \
		protoc --go_out=$$DIR --go_opt=paths=import \
			--go-grpc_out=$$DIR --go-grpc_opt=paths=import \
			$$DIR/$$PROTO_FILE || exit 1; \
		mv $$DIR/$(PACKAGE_BASE_DIR)/$$(basename $$DIR)/* $$DIR 2>/dev/null || true; \
		rm -rf $$DIR/github.com; \
	done

buf: generate-structs generate

proto-plugin-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get google.golang.org/grpc
	go get google.golang.org/protobuf