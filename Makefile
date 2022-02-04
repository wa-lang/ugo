default:
	go run main.go asm ./_examples/hello.ugo

run:
	go run main.go run ./_examples/prime.ugo

dev:
	go run main.go run ./_examples/prime.ugo

clean:
