compose-restart:
	@sh restart.sh

test-integration:
	go test -tags integration ./internal/.../handler/... -race -coverprofile=coverage_integration.out -coverpkg=./internal/.../handler/... -covermode=atomic -v

test-benchmark:
	go test -tags "benchmark integration" ./internal/.../handler/... -race -coverprofile=coverage_integration.out  -coverpkg=./internal/.../handler/... -covermode=atomic -v

test-unit:
	go test -v ./internal/... -race -covermode=atomic -v

test-all: test-integration test-unit

generate-mocks:
	mockgen -source=./pkg/httpclient/client.go -destination=./pkg/httpclient/mocks/client_mock.go -package=mocks
	mockgen -source=./pkg/dummyclient/client.go -destination=./pkg/dummyclient/mocks/dummyclient_mock.go -package=mocks
	mockgen -source=./pkg/claimer/claimer.go -destination=./pkg/claimer/mocks/claimer_mock.go -package=mocks
	mockgen -source=./pkg/rabbit/preproducer.go -destination=./pkg/rabbit/mocks/preproducer_mock.go -package=mocks
	mockgen -source=./pkg/rabbit/producer.go -destination=./pkg/rabbit/mocks/producer_mock.go -package=mocks
	mockgen -source=./pkg/rabbit/consumermanager.go -destination=./pkg/rabbit/mocks/consumermanager_mock.go -package=mocks
	mockgen -source=./pkg/rabbit/processor.go -destination=./pkg/rabbit/mocks/processor_mock.go -package=mocks
	mockgen -source=./pkg/rabbit/client.go -destination=./pkg/rabbit/mocks/client_mock.go -package=mocks
	mockgen -source=./pkg/rabbit/clientmanager.go -destination=./pkg/rabbit/mocks/clientmanager_mock.go -package=mocks
	mockgen -source=./internal/app/dummy/service.go -destination=./internal/app/dummy/mocks/service_mock.go -package=mocks

lint:
	goimports -w . && gofmt -w .
