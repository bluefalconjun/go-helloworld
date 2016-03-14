.DEFAULT_GOAL := testmain

SHELL := /bin/bash
CC := gcc

sourcefiles := $(wildcard *.c)
objects := $(patsubst %.c,%.o,$(wildcard *.c))

testmain: $(objects)
	$(CC) -Wall -Werror -g -fdiagnostics-color=always -o testmain  $(objects)

.PHONY: clean
clean:
	-rm -f testmain

distclean:
	-rm -f testmain $(objects)
