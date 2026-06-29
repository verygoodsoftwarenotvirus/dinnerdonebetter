# syntax=docker/dockerfile:1
FROM golang:1.26-trixie

WORKDIR /go/src/github.com/verygoodsoftwarenotvirus/dinnerdonebetter/backend
COPY . .

# to debug a specific test:
# ENTRYPOINT go test -parallel 1 -v -failfast github.com/verygoodsoftwarenotvirus/dinnerdonebetter/backend/tests/integration -run TestIntegration/TestValidPreparationInstruments_CompleteLifecycle

ENTRYPOINT ["go", "test", "-v", "github.com/verygoodsoftwarenotvirus/dinnerdonebetter/backend/tests/integration"]
