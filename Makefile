default:
	make -C ./01
	make -C ./02

clean:
	make -C ./01 clean
	make -C ./02 clean
