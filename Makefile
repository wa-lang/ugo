default:
	go run main.go ast ./_examples/fibonacci.ugo

run:
	go run main.go run ./_examples/prime.ugo

dev:
	go run main.go run ./_examples/prime.ugo

clean:
	-rm *.out*
