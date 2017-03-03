TARGETS = darwin/amd64 darwin/386 linux/amd64 linux/386 windows/amd64 windows/386

usage:
	@echo
	@echo "Command      : Description"
	@echo "------------ : ---------------"
	@echo "make clean   : Clean up the build files and reset assets"
	@echo "make release : Generate binaries for all supported OSes"
	@echo

clean:
	@rm -rf ./bin/*

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
	@mkdir ./bin/data
	@touch ./bin/data/proxy-{fetch,pool}.txt
	@cp ./data/mail-notification.txt ./bin/data/mail-notification.txt
	@cp ./data/cron-task.yaml.example ./bin/data/cron-task.yaml

	@echo "Done"
