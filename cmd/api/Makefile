.PHONY: init
init:
	hz new \
	-module github.com/SIN5t/CancerCellDetectSys/cmd/api \
	-service "upload" \
	-proto_path ../../ \
	-idl ../../idl/upload.proto
	&& go mod tidy

.PHONY: update
update:
	hz update \
	-proto_path ../../ \
	-idl ../../idl/upload.proto
.PHONY: run
run:
	go build -o hertz_test && ./hertz_test