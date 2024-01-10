.PHONY: help
.DEFAULT_GOAL := help

APP_NAME := security

# build cross platform binaries: linux, windows, mac/amd64 and mac/arm64
build:
	@echo "Building binaries..."
	@echo "Building for linux/amd64..."
	@GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux-amd64
	@echo "Building for linux/arm64..."
	@GOOS=linux GOARCH=arm64 go build -o bin/$(APP_NAME)-linux-arm64
	@echo "Building for windows/amd64..."
	@GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)-windows-amd64.exe
	@echo "Building for darwin/amd64..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-darwin-amd64
	@echo "Building for darwin/arm64..."
	@GOOS=darwin GOARCH=arm64 go build -o bin/$(APP_NAME)-darwin-arm64
	@echo "Done."

# Show help
help:
	@echo ""
	@echo "Usage:"
	@echo "    make [target]"
	@echo ""
	@echo "Targets:"
	@awk '/^[a-zA-Z\-_0-9]+:/ \
	{ \
		helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} { lastLine = $$0 }' $(MAKEFILE_LIST)
