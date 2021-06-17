INTEGRATION_TEST_PATH?=./routes

ENV_LOCAL_TEST=DATABASE_DRIVER=memory

docker.start:
	docker-compose up -d;

docker.stop:
	docker-compose down;

test.integration:
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -run=$(INTEGRATION_TEST_PATH);