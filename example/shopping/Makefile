.PHONY: proto build

proto:
	for d in pb/**.proto; do \
        protoc -I pb --go_out=plugins=grpc:pb $$d; \
		echo compiled: $$d; \
	done

lint:
	./bin/lint.sh

build:
	./bin/build.sh

clear:
	./bin/clear.sh