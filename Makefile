env := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
fileName := aiChat

.PHONY: all build clean kill setconfig showResult

all: clean build run
	@echo "success"

build:
	$(env) go build -o $(fileName) ./*.go
	@echo "$(fileName) build success"

clean:
	@rm -rf $(fileName)
	@echo "$(fileName) clean success"

run:
	@nohup ./$(fileName) &
	@echo "$(fileName) run success"
kill:
	@-kill -9 $(word 2,$(shell ps -ef | grep $(fileName) | grep -v grep | awk '{print $2}'))

setconfig:
	@ln -s ../config config

showResult:
	@cat nohup.out