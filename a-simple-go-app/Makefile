GOCMD=go
GOBUILD=$(GOCMD) build -o
BINARY_NAME=a-simple-go-app

start: 
	@make build
	@./$(BINARY_NAME)

build: 
	@CGO_ENABLED=0 $(GOBUILD) $(BINARY_NAME) -v github.com/fredrikssongustav/a-simple-cluster/a-simple-go-app

## Clean output files.
clean:
	@rm -f $(BINARY_NAME)

## Hacky way to update our docker image
docker:
	@docker build -t a-simple-go-app .
	@docker tag a-simple-go-app gustavfredriksson/a-simple-go-app:1.0.0
	@docker push gustavfredriksson/a-simple-go-app:1.0.0
fmt:
	@go fmt ./...

