default: run

run:
	go run main.go

test:
	ginkgo common windows menu

watch_test:
	ginkgo watch --notify common windows menu

godep:
	godep save ./...