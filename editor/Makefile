.PHONY: build

usage:
	@echo "Commands     : Description"
	@echo "------------ : -----------"
	@echo "make install : Install the packages"
	@echo "make watch   : Watch the file changes and rebuild and reload them"
	@echo "make bindata : Generate bindata"
	@echo "make build   : Build the editor"
	@echo "make clean   : Clean the build and asset files"

install:
	@npm install

watch:
	@npm run dev

bindata: clean build
	@cd dist && $(GOPATH)/bin/go-bindata -o editor.go -pkg editor -ignore="\\.DS_Store|\\.gitignore|\\.gitkeep" ./...
	@mv dist/editor.go editor.go

build:
	@npm run build

clean:
	@rm -rf dist
	@rm -rf editor.go
