default: run

run:
	go run main.go

test:
	ginkgo watch --notify common windows

godep:
	godep save ./...