PORT=23600

NVM_SOURCE_PATH ?= $(HOME)/.nvm/nvm.sh

ifneq ("$(wildcard $(NVM_SOURCE_PATH))","")
	NVM_EXEC = source $(NVM_SOURCE_PATH) && nvm exec --
endif
NPM = $(NVM_EXEC) npm

.PHONY: all
all: audit build watch

.PHONY: audit
audit:
	dis-vulncheck
	$(NPM) audit --audit-level=high

.PHONY: build
build: deps-javascript
	go run main.go

.PHONY: deps-javascript
deps-javascript:
	$(NPM) install --unsafe-perm

.PHONY: install-prereqs
install-prereqs:
	go install github.com/fogleman/serve@latest
	go install github.com/cespare/reflex@latest

.PHONY: test
test: deps-javascript
	$(NPM) test

.PHONY: watch
watch:
	mkdir -p logs
	trap 'kill %1;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /'

.PHONY: watch-templates
watch-templates:
	make build
	reflex -d none -c ./reflex

.PHONY: watch-assets
watch-assets:
	$(NPM) run build
	$(NPM) run watch

.PHONY: watch-serve
watch-serve:
	mkdir -p logs
	trap 'kill %1; kill %2;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /' & make serve | tee logs/server.log | sed -e 's/^/[Server] /'

.PHONY: serve
serve:
	serve -port=${PORT} -dir="assets"

.PHONY: clean
clean:
	rm -rf assets/ logs/ node_modules/
