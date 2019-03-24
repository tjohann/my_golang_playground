#
# BSD 3-Clause
# (c) 2015, thorsten.johannvorderbrueggen@t-online.de
#

.PHONY: all clean

all:
	for dir in $(MODULES); do (cd $$dir && $(MAKE) $@); done
	@echo
	@echo "+-----------------------------------------------------+"
	@echo "|                                                     |"
	@echo "|                   Cheers $(USER)                    |"
	@echo "|                                                     |"
	@echo "+-----------------------------------------------------+"
	@echo

clean:
	rm -f *~
	rm -f .*~
	rm -f os_args

distclean: clean
