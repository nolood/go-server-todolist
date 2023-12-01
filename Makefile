PID = ./obsly.pid
GO_FILES = $(wildcard *.go)
APP = ./obsly

serve: start
	@fswatch -x -o --event Created --event Updated --event Renamed -r -e '.*' -i '\.go$$'  . | xargs -n1 make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@echo "actually do nothing"

build: $(GO_FILES)
	@go build -o $(APP)

$(APP): $(GO_FILES)
	@go build $? -o $@

start:
	@go run $(GO_FILES) & echo $$! > $(PID)

restart: kill before build start
