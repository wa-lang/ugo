package builtin

import _ "embed"

//go:embed builtin_darwin_amd64.ll
var llBuiltin_dawin_amd64 string

func GetBuiltinLL(goos, goarch string) string {
	switch goos {
	case "darwin":
		switch goarch {
		case "amd64":
			return llBuiltin_dawin_amd64
		}
	case "linux":
	case "windows":
	}
	panic("builtin missing")
}

const Header = `
declare i32 @ugo_builtin_println(i32)
declare i32 @ugo_builtin_exit(i32)
`

const MainMain = `
define i32 @main() {
	call i32() @ugo_main_main()
	ret i32 0
}
`
