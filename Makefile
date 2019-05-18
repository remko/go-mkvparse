check:
	go test .

lint:
	staticcheck .
	go vet .
