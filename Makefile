.PHONY: help update
.DEFAULT_GOAL := help

# update all deps with workspaces
update:
	@echo "Updating all dependencies with workspaces..."
	@go get -u ./... && go mod tidy
	@cd encrypt && go get -u ./... && go mod tidy
	@cd fiber && go get -u ./... && go mod tidy
	@cd fiber_boot && go get -u ./... && go mod tidy
	@cd polymorphism && go get -u ./... && go mod tidy
	@cd sqlite3_demo && go get -u ./... && go mod tidy
	@go work sync

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
