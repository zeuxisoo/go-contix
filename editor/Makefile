.PHONY: build

usage:
	@echo "Commands     : Description"
	@echo "------------ : -----------"
	@echo "make install : Install the packages"
	@echo "make watch   : Watch the file changes and rebuild and reload them"
	@echo "make build   : Build the editor"
	@echo "make clean   : Clean the build and asset files"

install:
	@npm install

watch:
	@npm run dev

build:
	@npm run build

clean:
	@rm -rf dist
