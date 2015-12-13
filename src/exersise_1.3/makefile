#
# BSD 3-Clause
# (c) 2015, thorsten.johannvorderbrueggen@t-online.de                        
#

TEST_STRING = that is a long string which builds a nice but senseless test string to check the duration from string addtions by hand and using join for it

.PHONY: all clean

all: os_args.go os_args.c
	go build os_args.go
	gcc -Wall -o os_args_c os_args.c
	g++ --std=c++11 -Wall -o os_args_cpp os_args.cc

clean:
	rm -f *~
	rm -f .*~
	rm -f os_args
	rm -f os_args_c
	rm -f os_args_cpp

run: all
	./os_args ${TEST_STRING}
	./os_args_c ${TEST_STRING}
	./os_args_cpp ${TEST_STRING}
