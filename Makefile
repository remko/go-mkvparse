GO_TEST_FLAGS:=-timeout 5s
ifeq ($(V),1)
GO_TEST_FLAGS:=$(GO_TEST_FLAGS) -v
endif
ifeq ($(COVERAGE),1)
GO_TEST_FLAGS:=$(GO_TEST_FLAGS) -coverprofile=coverage.out
endif

.PHONY: all
all: check examples

.PHONY: examples
examples:
	mkdir -p bin
	go build -o bin/ ./examples/...

.PHONY: matroska-test-files
matroska-test-files:
	git clone https://github.com/ietf-wg-cellar/matroska-test-files.git testdata/matroska-test-files

.PHONY: check
check: 
	go test $(GO_TEST_FLAGS) . 
ifeq ($(COVERAGE),1)
	go tool cover -html=coverage.out
endif

.PHONY: install-tools
install-tools:
	go install honnef.co/go/tools/cmd/staticcheck@2023.1.3

.PHONY: lint
lint:
	go vet ./...
	staticcheck ./...

example-live+junk.mkv:
	ffmpeg -t 1 -s 320x240 -f rawvideo -r 25 -pix_fmt rgb24 -i /dev/zero -metadata title="Live + Junk" -metadata author="John Doe" -c:v libx264 -pix_fmt yuv420p dirty.$@
	mkclean --live dirty.$@ $@
	-rm -rf dirty.$@
