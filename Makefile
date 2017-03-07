.PHONY: vendor

TARGETS = darwin/amd64 darwin/386 linux/amd64 linux/386 windows/amd64 windows/386

usage:
	@echo
	@echo "Command      : Description"
	@echo "------------ : ---------------"
	@echo "make develop : Install the develop tools"
	@echo "make vendor  : Install the vendors"
	@echo "make watch   : Watch the changes and rebuild the go-contix application"
	@echo "make clean   : Clean up the build files and reset assets"
	@echo "make release : Generate binaries for all supported OSes"
	@echo "make windows : Generate assets for windows environment like cmder, fonts"
	@echo

develop:
	@go get github.com/Unknwon/bra

vendor:
	@glide install

watch:
	@$(GOPATH)/bin/bra run

clean:
	@rm -rf ./bin/*
	@rm -rf ./go-contix

clean-windows:
	@rm -rf ./bin/windows

release: clean
	@echo "Downloading gox"
	@go get github.com/mitchellh/gox

	@echo "Building binaries..."
	@$(GOPATH)/bin/gox \
		-osarch "$(TARGETS)" \
		-output "./bin/contix_{{.OS}}_{{.Arch}}"

	@echo "Building ARM binaries..."
	@GOOS=linux GOARCH=arm GOARM=5 go build -o "./bin/contix_linux_arm_v5"

	@echo "Generating assets"
	@mkdir -p ./bin/data
	@touch ./bin/data/proxy-{fetch,pool}.txt
	@cp ./data/mail-notification.txt ./bin/data/mail-notification.txt
	@cp ./data/cron-task.yaml.example ./bin/data/cron-task.yaml

	@echo "Done"

windows: clean-windows
	@echo "Creating windows assets"
	@mkdir -p ./bin/windows

	@echo "Downloading cmder"
	@wget https://github.com/cmderdev/cmder/releases/download/v1.3.2/cmder_mini.zip \
		-O ./bin/windows/cmder.zip

	@echo "Downloading fonts"
	@wget https://excellmedia.dl.sourceforge.net/project/dejavu/dejavu/2.37/dejavu-fonts-ttf-2.37.zip \
		-O ./bin/windows/fonts.zip

	@echo "Repacking cmder and fonts"
	@cd ./bin/windows && unzip cmder.zip -d cmder
	@cd ./bin/windows && unzip fonts.zip && mkdir fonts && cp dejavu-fonts-ttf-*/ttf/DejaVuSansMono.ttf fonts/DejaVuSansMono.ttf
	@cd ./bin/windows && rm -rf ./cmder.zip ./fonts.zip ./dejavu-fonts-ttf-*/

	@echo "Done"
