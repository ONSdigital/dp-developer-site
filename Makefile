build:
	go run main.go

watch:
	go get github.com/cespare/reflex
	reflex -r '\.(go|tmpl)$$' go run main.go

.PHONY: build watch