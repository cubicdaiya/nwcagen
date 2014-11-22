
nwcagen:
	go build -o nwcagen

fmt:
	go fmt ./...

check:
	go test

clean:
	rm -rf nwcagen
