
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	go mod tidy
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/create-beauty-sol-address /usr/local/bin/create-beauty-sol-address

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/create-beauty-sol-address.service /etc/systemd/system/create-beauty-sol-address.service
	sudo systemctl daemon-reload
	@echo
	@echo "create-beauty-sol-address service installed."

