
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	go mod tidy
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/hardhat-abigen /usr/local/bin/hardhat-abigen

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/hardhat-abigen.service /etc/systemd/system/hardhat-abigen.service
	sudo systemctl daemon-reload
	@echo
	@echo "hardhat-abigen service installed."

