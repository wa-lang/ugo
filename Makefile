default:
	go run main.go run ./_examples/fib2.ugo

run:
	go run main.go run ./_examples/fib2.ugo

dev:
	go run main.go -debug run ./_examples/fib2.ugo

clean:
	-rm *.out*
