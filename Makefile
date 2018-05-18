build:
	go run main.go
	npm install --unsafe-perm

watch:
	mkdir logs
	trap 'kill %1;' SIGINT
	make watch-templates | tee logs/templates.log | sed -e 's/^/[Templates] /' & make watch-assets | tee logs/assets.log | sed -e 's/^/[Assets] /'

watch-templates:
	make build
	go get github.com/cespare/reflex
	reflex -v -R node_modules -R assets -R vendor -R logs -r '^(main\.go|renderer/.*|templates/.*\.tmpl|static/.*\.md)$$' go run main.go

watch-assets:
	npm run build
	npm run watch

.PHONY: build watch