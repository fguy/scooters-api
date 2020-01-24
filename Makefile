PKG=github.com/fguy/scooters-api

server:
	go build -o server
run: 
	go run .
test:
	go test -v -cover ./...
clean:
	rm -rf server
docker:
	docker build -t hub.docker.com/fguy/scooters-api .
mock:
	mockgen -destination=mocks/go.uber.org/fx/lifecycle.go go.uber.org/fx Lifecycle
	go generate ./...
