help: ### Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

run: ### 1-run server
	go run main.go handler.go

test_curl:### 2-Run curl test after run server
	@echo "---- Fail:No Name Error"
	curl -X POST localhost:8080/save -H "Content-Type: application/json" -d '{"age":20}'
	@echo "---- Fail:less age than 16 Error"
	curl -X POST localhost:8080/save -H "Content-Type: application/json" -d '{"name":"Alishah","age":15}'
	@echo "---- Pass"
	curl -X POST localhost:8080/save -H "Content-Type: application/json" -d '{"name":"Alishah","age":20}'

test:### 3-Run unit test
	go test -v -race .

