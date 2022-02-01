default:
	clang -S -emit-llvm _builtin.c
	clang -Woverride-module _builtin.ll main.ll
	./a.out || echo $$?

builtin-ll-darwin-amd64:
	clang -S -emit-llvm _builtin.c -o builtin_darwin_amd64.ll

clean:
