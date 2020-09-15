# Brownie compiles each solidity source file only when it has changed.
# If contract A depends on contract B, and we change B in an
# incompatible way, A will not be recompiled when we run 'brownie
# test' or 'brownie compile' and it will even run the tests with the
# cached compilation output of A.
#
# This Makefile makes sure we always recompile all contracts.

SRCS = $(wildcard contracts/*.sol)

compile: build/compile

build/compile: ${SRCS}
	brownie compile --all
	@touch build/compile

test: compile
	brownie test

clean:
	rm -f build/*.sol build/compile

.PHONY: compile test clean
