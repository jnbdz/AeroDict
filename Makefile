include config.mk

.DEFAULT_GOAL := help

.PHONY: install
install: ## Install AeroDict
	mkdir -p $(app_path)
	mkdir -p $(app_desktop_path)
	mkdir -p $(config_path)
	cp $(CURDIR)/aerodict $(app_path)
	cp -r $(CURDIR)/.config/AeroDict/* $(config_path)

.PHONY: 

.PHONY: help
help: ## Help
        @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
