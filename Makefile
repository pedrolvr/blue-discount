CMD_DIR=./cmd
DISCOUNT_DIR=$(CMD_DIR)/discount

.PHONY: run
run:
	go run $(DISCOUNT_DIR)/main.go

.PHONY: test
test:
	ginkgo -r -race -failFast -progress -cover

.PHONY: test-clear
test-clear:
	rm -rf ./coverage/*.*

.PHONY: test-cov
test-cov: test-clear
	ginkgo -r -race -failFast -progress -cover -coverprofile=coverage.out -outputdir=./coverage/
	go tool cover -html=./coverage/coverage.out
