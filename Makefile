default: run

run:
	go run main.go

test:
	ginkgo common windows

godep:
	godep save ./...