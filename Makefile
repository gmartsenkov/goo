default: run

run:
	go run main.go

test:
	ginkgo common windows menu common

watch_test:
	ginkgo watch --notify common windows menu

godep:
	godep save ./...