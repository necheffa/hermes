COVERLOG="./coverage.out"
VERSION=$$(cat cmd/hermes/VERSION)
GOAMD64:=v3

all: hermes

hermes:
	cd cmd/hermes; GOAMD64=$(GOAMD64) go build -buildmode=pie

test:
	GOAMD64=$(GOAMD64) ginkgo -r -v --race --trace --coverpkg=necheff.net/hermes --coverprofile=$(COVERLOG) ./...
	go tool cover -html=$(COVERLOG)

quality:
	go vet ./...
	golangci-lint run --enable godox --enable mnd --enable gosec --enable errorlint --enable gofmt --enable unconvert \
        --enable ginkgolinter ./...

debian: hermes
	scripts/package-deb $(VERSION)

vulns:
	govulncheck -show verbose ./...

clean:
	rm -f $(COVERLOG) cmd/hermes/hermes hermes_$(VERSION)-*_amd64.deb
