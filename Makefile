default:
	go run main.go ast ./_examples/import.ugo

asm:
	go run main.go asm ./_examples/import.ugo

run:
	go run main.go run ./_examples/import.ugo

dev:
	go run main.go -debug run ./_examples/import.ugo

clean:
	-rm *.out*
