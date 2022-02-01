default:
	make -C ./builtin builtin-ll
	go run main.go > a.out.ll

	clang -Wno-override-module ./a.out.ll ./builtin/a.out.builtin.ll
	./a.out || echo $$?

	@make clean

clean:
	@make -C ./builtin clean
	-@rm *.out*
