.POSIX:

lint:
	golangci-lint run --tests

lint-changes:
	base=`git merge-base HEAD origin/main`; \
	golangci-lint run --new-from-rev $${base}
