default:
	echo "1+3-2" | go run main.go
	-@rm *.out*

ll:
	clang -Wno-override-module main.ll
	-./a.out || echo $$?

	-@rm *.out

clean:
	-@rm *.out*
