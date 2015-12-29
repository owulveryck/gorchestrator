TARGET:=./dist/$(shell uname -s)/$(shell uname -p)

executor: $(TARGET)/executor/executor

orchestrator: $(TARGET)/orchestrator/orchestrator

clients: $(TARGET)/clients/web

$(TARGET)/generate_cert: security/util/generate_cert.go
	go build -o $(TARGET)/generate_cert security/util/generate_cert.go

$(TARGET)/clients/web: clients/web/*.go clients/web/htdocs/* clients/web/tmpl/*
	go build -o $(TARGET)/clients/web/webclient clients/web/*go
	cp -r clients/web/htdocs clients/web/tmpl $(TARGET)/clients/web/

$(TARGET)/orchestrator/orchestrator: orchestrator/*.go 
	go build -o $(TARGET)/orchestrator/orchestrator main.go

$(TARGET)/executor/executor: executor/*.go executor/cmd/*.go
	go build -o $(TARGET)/executor/executor executor/cmd/main.go

$(TARGET)/certs/orchestrator.pem: $(TARGET)/certs/orchestrator_key.pem

$(TARGET)/certs/orchestrator_key.pem: $(TARGET)/generate_cert
	mkdir -p $(TARGET)/certs && \
	cd $(TARGET)/certs && \
	../generate_cert -ca -host 127.0.0.1 -target orchestrator

$(TARGET)/certs/executor.pem: $(TARGET)/certs/executor_key.pem

$(TARGET)/certs/executor_key.pem:
	mkdir -p $(TARGET)/certs && \
	cd $(TARGET)/certs && \
	../generate_cert -ca -host 127.0.0.1 -target executor 

certificates: $(TARGET)/certs/orchestrator_key.pem $(TARGET)/certs/executor_key.pem

install_certificates: certificates $(TARGET)/orchestrator/orchestrator $(TARGET)/executor/executor
	cp $(TARGET)/certs/orchestrator*pem $(TARGET)/certs/executor.pem $(TARGET)/orchestrator && \
	cp $(TARGET)/certs/executor*pem $(TARGET)/certs/orchestrator.pem $(TARGET)/executor 
	

dist: $(TARGET)/executor/executor $(TARGET)/orchestrator/orchestrator $(TARGET)/generate_cert $(TARGET)/clients/web install_certificates

clean:
	rm -rf $(TARGET)
