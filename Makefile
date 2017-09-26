default: run

run:
	go run main.go

test:
	ginkgo common

godep:
	godep save ./...