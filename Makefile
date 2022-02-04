default:
	go run main.go run ./_examples/fib1.ugo

run:
	go run main.go run ./_examples/fib1.ugo

dev:
	go run main.go -debug run ./_examples/fib1.ugo

clean:
	-rm *.out*
