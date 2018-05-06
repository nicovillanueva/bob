run:
	go run *.go

build:
	go build

release:
	docker build -t nicovillanueva/bob:0.1 . && docker push nicovillanueva/bob:0.1

prepare:
	dep ensure
