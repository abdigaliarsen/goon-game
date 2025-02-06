PROTO_DIRS = pkg/proto/wikipedia
PACKAGE_BASE_DIR = github.com/abdigaliarsen/goon-game/pkg/proto
MKFILE_PATH = $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_ROOT = $(patsubst %/,%,$(dir $(MKFILE_PATH)))

project-root:
	echo $(PROJECT_ROOT)

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

hard-restart:
	@if [ -n "$$(docker ps -aq)" ]; then docker stop $$(docker ps -aq); fi
	@if [ -n "$$(docker ps -aq)" ]; then docker rm $$(docker ps -aq); fi
	@if [ -n "$$(docker volume ls -q)" ]; then docker volume rm $$(docker volume ls -q); fi
	@if [ -n "$$(docker network ls -q --filter type=custom)" ]; then docker network rm $$(docker network ls -q --filter type=custom); fi
	@if [ -n "$$(docker images -aq)" ]; then docker rmi $$(docker images -aq); fi
	docker compose -f $(PROJECT_ROOT)/docker/docker-compose.yml up -d

run:
	docker compose -f $(PROJECT_ROOT)/docker/docker-compose.yml up -d --build

stop:
	docker compose -f $(PROJECT_ROOT)/docker/docker-compose.yml down
