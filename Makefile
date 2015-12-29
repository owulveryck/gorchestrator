TARGET:=./tests/binaries

executor: $(TARGET)/executor

orchestrator: $(TARGET)/orchestrator

clients: $(TARGET)/clients/web

$(TARGET)/generate_cert: security/util/generate_cert.go
	go build -o $(TARGET)/generate_cert security/util/generate_cert.go

$(TARGET)/clients/web: clients/web/*.go clients/web/htdocs/* clients/web/tmpl/*
	go build -o $(TARGET)/clients/web/webclient clients/web/*go
	cp -r clients/web/htdocs clients/web/tmpl $(TARGET)/clients/web/

$(TARGET)/orchestrator: orchestrator/*.go 
	go build -o $(TARGET)/orchestrator main.go

$(TARGET)/executor: executor/*.go executor/cmd/*.go
	go build -o $(TARGET)/executor executor/cmd/main.go

dist: $(TARGET)/executor $(TARGET)/orchestrator $(TARGET)/generate_cert
	mkdir -p $(TARGET)/security


