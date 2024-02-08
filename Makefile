run: build
	@./bin/ctl-the-kube

build:
	@go build -o bin/ctl-the-kube
