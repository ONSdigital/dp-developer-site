build:
	go run main.go

watch:
	make build
	go get github.com/cespare/reflex
	reflex -r '\.(go|html|tmpl|md)$$' go run main.go

.PHONY: build watch
