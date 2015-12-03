#
# BSD 3-Clause
# (c) 2015, thorsten.johannvorderbrueggen@t-online.de                        
#

.PHONY: all clean

all: os_args.go os_args.c
	go build os_args.go
	gcc -Wall -o os_args_c os_args.c

clean:
	rm -f *~
	rm -f .*~
	rm -f os_args
	rm -f os_args_c

run: all
	./os_args that is a long string which builds a nice but senseless test string to check the duration from string addtions by hand and using join for it
	./os_args_c that is a long string which builds a nice but senseless test string to check the duration from string addtions by hand and using join for it
