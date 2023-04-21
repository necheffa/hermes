COVERLOG="./coverage.out"
VERSION=$$(cat cmd/hermes/VERSION)

all: hermes

hermes:
	cd cmd/hermes; go build

test:
	ginkgo -r -v --race --trace --coverpkg=necheff.net/hermes --coverprofile=$(COVERLOG) ./...
	go tool cover -html=$(COVERLOG)

quality:
	go vet ./...
	golangci-lint run --enable godox --enable gomnd --enable gosec --enable errorlint --enable gofmt --enable unconvert ./...

debian: hermes
	scripts/package-deb $(VERSION)-1

clean:
	rm -f $(COVERLOG) cmd/hermes/hermes hermes_$(VERSION)-1_amd64.deb
