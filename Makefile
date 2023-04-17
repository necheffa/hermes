COVERLOG="./coverage.out"

all: hermes

hermes:
	cd cmd/hermes; go build

test:
	ginkgo -r -v --race --trace --coverpkg=necheff.net/hermes --coverprofile=$(COVERLOG) ./...
	go tool cover -html=$(COVERLOG)

quality:
	go vet ./...

clean:
	rm -f $(COVERLOG) cmd/hermes/hermes
