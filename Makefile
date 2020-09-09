CMD_DIR=./cmd
DISCOUNT_DIR=$(CMD_DIR)/discount

.PHONY: run
run:
	go run $(DISCOUNT_DIR)/main.go

.PHONY: test
test:
	ginkgo -r -race -failFast -progress
