# we will put our integration testing in this path
INTEGRATION_TEST_PATH?=./routes

# set of env variables that you need for testing
ENV_LOCAL_TEST=\
  DATABASE_DRIVER=memory

# this command will start a docker components that we set in docker-compose.yml
docker.start.components:
  ./docker/docker-compose up -d --remove-orphans;

# shutting down docker components
docker.stop:
  ./docker/docker-compose down;

# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
  $(ENV_LOCAL_TEST) \
  go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -run=$(INTEGRATION_TEST_PATH)

# this command will trigger integration test with verbose mode
test.integration.debug:
  $(ENV_LOCAL_TEST) \
  go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_PATH)