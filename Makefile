default:
	go run main.go

clean:
	@make -C ./builtin clean
	-@rm *.out*
