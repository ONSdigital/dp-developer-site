build:
	go run main.go
	npm install
watch:
	make build
	go get github.com/cespare/reflex
	reflex -r '\.(go|tmpl|md)$$' go run main.go

.PHONY: build watch