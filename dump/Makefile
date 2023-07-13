default_exec := build/dump


.PHONY: build
build :
	go build -o $(default_exec) -tags production -v .

.PHONY: run
run :
	$(default_exec)

.PHONY: clean
clean :
	go clean -x
	rm build/*


