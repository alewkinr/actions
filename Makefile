build:
	docker build -t github.com/alewkinr/actions .

run:
	docker run -it --rm github.com/alewkinr/actions:latest

mod:
	go mod tidy && go mod download && go mod vendor