BINARY = nventory
SOURCE_FILES = $(shell find ./ -type f -name '*.go')

NVENTORY_COMMIT = $(shell git rev-parse HEAD)
NVENTORY_PKG = github.com/rclsilver/nventory

VERSION := $(shell git describe --exact-match --abbrev=0 --tags $(git rev-list --tags --max-count=1) 2> /dev/null)
ifndef VERSION
	VERSION = development
endif

define build_binary
	go build \
		-ldflags "-X ${NVENTORY_PKG}.Commit=${NVENTORY_COMMIT} -X ${NVENTORY_PKG}.Version=${VERSION}" \
		-o $(1)
endef

all: ${BINARY}

${BINARY}: $(SOURCE_FILES) go.mod
	$(call build_binary,${BINARY})

clean:
	rm -f ${BINARY}

.PHONY: all clean
