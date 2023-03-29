qa: analyze test

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

test:
	@go test -cover ./...

benchmark:
#	$ needs to be escaped by $$ in Makefiles
	@go test -bench=. -run=^$$ ./...

coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/cover.out ./...
	@go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	@xdg-open ./coverage/cover.html

clean:
	@rm -rf build/

build: clean
	$(eval VERSION=$(shell git tag --points-at HEAD))
	$(eval VERSION=$(or $(VERSION), (version unavailable)))

	go build -ldflags="-X 'git.staubwolke.org/yeldir/go-ecs/version.Version=$(VERSION)'" -o ./build/go-ecs cli/main.go

.PHONY: qa \
	analyze \
	test \
	benchmark \
	coverage \
	clean \
	build
