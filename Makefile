TARGET:=./dist/$(shell uname -s)/$(shell uname -p)

executor: $(TARGET)/executor

orchestrator: $(TARGET)/orchestrator

clients: $(TARGET)/clients/web

$(TARGET)/clients/web: clients/web/*.go clients/web/htdocs/* clients/web/tmpl/*
	go build -o $(TARGET)/clients/web/webclient clients/web/*go
	cp -r clients/web/htdocs clients/web/tmpl $(TARGET)/clients/web/

$(TARGET)/orchestrator: orchestrator/*.go 
	go build -o $(TARGET)/orchestrator main.go

$(TARGET)/executor: executor/*.go executor/cmd/*.go
	go build -o $(TARGET)/executor executor/cmd/main.go

dist: $(TARGET)/executor $(TARGET)/orchestrator
	mkdir -p $(TARGET)/security
	cp -r security/certs $(TARGET)/security
