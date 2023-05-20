build:
	@go build -o ./facts/catfact

run: build
	./facts/catfact

