# µGo 编程语言

µGo 是 Go 语言的真子集(不含标准库部分), 可以直接作为Go代码编译执行.

- 安装 ugo: `go get github.com/wa-lang/ugo`.
- 实现原理: https://github.com/wa-lang/ugo-compiler-book

## 例子

例子 ([_example/hello.ugo](_example/hello.ugo)):

```go
package main

func main() {
	for n := 2; n <= 30; n = n + 1 {
		var isPrime int = 1
		for i := 2; i*i <= n; i = i + 1 {
			if x := n % i; x == 0 {
				isPrime = 0
			}
		}
		if isPrime != 0 {
			println(n)
		}
	}
}
```

运行:

```
$ ugo run _examples/hello.ugo 
2
3
5
7
11
13
17
19
23
29
```

## 版权

个人学习目的可自由使用.
