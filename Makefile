expr:
	go run main.go -lex -ast -llir -mode=expr -file=expr.ugo

file:
	go run main.go -debug=false -lex -ast -llir -mode=file -file=file.ugo

clean:
