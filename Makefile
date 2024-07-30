PORT=23600

.PHONY: all
all: audit build watch

.PHONY: audit
audit:
	go list -json -m all | nancy sleuth
	npm audit --audit-level=high

.PHONY: build
build: deps-javascript
	go run main.go
	
.PHONY: deps-javascript
deps-javascript:
	npm install --unsafe-perm

.PHONY: install-prereqs
install-prereqs:
	go install github.com/fogleman/serve@latest

.PHONY: test
test: deps-javascript
	npm test

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
	npm run build
	npm run watch

.PHONY: watch-serve
watch-serve:
	mkdir -p logs
	trap 'kill %1; kill %2;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /' & make serve | tee logs/server.log | sed -e 's/^/[Server] /'

.PHONY: serve
serve:
	serve -port=${PORT} -dir="assets"
