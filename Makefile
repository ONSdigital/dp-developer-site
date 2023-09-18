PORT=23600

.PHONY: all
all: audit build watch

.PHONY: audit
audit:
	go list -json -m all | nancy sleuth
	npm audit --audit-level=high

.PHONY: build
build:
	go run main.go
	npm install --unsafe-perm

.PHONY: test
test:
	npm test

.PHONY: watch
watch:
	mkdir -p logs
	trap 'kill %1;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /'

.PHONY: watch-templates
watch-templates:
	make build
	go install github.com/cespare/reflex
	reflex -v -R node_modules -R assets -R vendor -R logs -r '^(main\.go|renderer/.*|templates/.*\.tmpl|static/.*\.md|static/.*\.html)$$' go run main.go

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
	go install github.com/fogleman/serve
	serve -port=${PORT} -dir="assets"
