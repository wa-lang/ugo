default:
	make builtin-ll
	clang -Wno-override-module a.out.builtin.ll main.ll
	./a.out || echo $$?

	-@make clean

builtin-ll:
	clang -S -emit-llvm _builtin.c -o a.out.builtin.ll

clean:
	-@rm *.out*
