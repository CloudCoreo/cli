OSES       := linux darwin freebsd windows
ARCHES     := 386 amd64
EXECUTABLE ?= cloudcoreo-cli

.PHONY: all
all: vet test build

.PHONY: fmt
fmt:
	govendor fmt +local

.PHONY: vet
vet:
	govendor vet +local

.PHONY: test
test:
	govendor test -v -race +local

.PHONY: build
build:
	for os in $(OSES); do \
		for arch in $(ARCHES); do \
			echo "GOOS=$$os GOARCH=$$arch go build -o ./build/$(EXECUTABLE)-$${os}-$${arch} ." && \
			GOOS=$$os GOARCH=$$arch go build -o ./build/$(EXECUTABLE)-$${os}-$${arch} .; \
		done \
	done

.PHONY: clean
clean:
	@rm -rf ./build
