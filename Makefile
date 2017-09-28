default: run

run:
	go run main.go

test:
	ginkgo common windows

watch_test:
	ginkgo watch --notify common windows

godep:
	godep save ./...