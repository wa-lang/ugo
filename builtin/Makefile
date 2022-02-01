default:
	clang -S -emit-llvm _builtin.c
	clang -Woverride-module _builtin.ll main.ll
	./a.out || echo $$?

builtin-ll-darwin-amd64:
	clang -S -emit-llvm _builtin.c -o builtin_darwin_amd64.ll

builtin-ll-darwin-arm64:
	clang -S -emit-llvm _builtin.c -o builtin_darwin_arm64.ll

builtin-ll-linux-amd64:
	clang -S -emit-llvm _builtin.c -o builtin_linux_amd64.ll

builtin-ll-linux-arm64:
	clang -S -emit-llvm _builtin.c -o builtin_linux_arm64.ll

builtin-ll-windows-amd64:
	clang -S -emit-llvm _builtin.c -o builtin_windows_amd64.ll

clean:
