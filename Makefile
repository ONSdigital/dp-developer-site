PORT=23600

build:
	go run main.go
	npm install --unsafe-perm

watch:
	mkdir -p logs
	trap 'kill %1;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /'

watch-templates:
	make build
	go get github.com/cespare/reflex
	reflex -v -R node_modules -R assets -R vendor -R logs -r '^(main\.go|renderer/.*|templates/.*\.tmpl|static/.*\.md|static/.*\.html)$$' go run main.go

watch-assets:
	npm run build
	npm run watch

watch-serve:
	mkdir -p logs
	trap 'kill %1; kill %2;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /' & make serve | tee logs/server.log | sed -e 's/^/[Server] /'

serve:
	go get github.com/fogleman/serve
	serve -port=${PORT} -dir="assets"

.PHONY: build watch
