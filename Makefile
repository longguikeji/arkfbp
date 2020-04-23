override BIN_DIR      := bin
override APP_NAME     := arkfbp
override CURRENT_DIR  := $(shell pwd)
override SYSTEM_ARCH  := $(shell uname)

GOBINDATA := $(BIN_DIR)/go-bindata

all: $(APP_NAME)

$(GOBINDATA):
	go get -u github.com/go-bindata/go-bindata/...

.PHONY: asset
asset: $(GOMETALINTER)
	go-bindata -pkg cmd -o cmd/asset.go ./asset/...

$(APP_NAME):
	@echo "Building $@ ..."
	@env \
		go build -o $(BIN_DIR)/$@ $^
	@echo "Build $@ finished!\n"


.PHONY: clean
clean:
	rm -f $(BIN_DIR)/$(APP_NAME)