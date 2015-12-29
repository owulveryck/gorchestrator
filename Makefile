TARGET:=./dist/$(shell uname -s)/$(shell uname -p)

executor: $(TARGET)/executor/executor

orchestrator: $(TARGET)/orchestrator/orchestrator

clients: $(TARGET)/clients/web

clients: $(TARGET)/clients/tosca $(TARGET)/clients/web

$(TARGET)/generate_cert: security/util/generate_cert.go
	go build -o $(TARGET)/generate_cert security/util/generate_cert.go

$(TARGET)/clients/tosca: clients/tosca/*.go
	go build -o $(TARGET)/clients/tosca/tosca2gorch clients/tosca/*.go

$(TARGET)/clients/web: clients/web/*.go clients/web/htdocs/* clients/web/tmpl/*
	go build -o $(TARGET)/clients/web/webclient clients/web/*go
	cp -r clients/web/htdocs clients/web/tmpl $(TARGET)/clients/web/

$(TARGET)/orchestrator/orchestrator: orchestrator/*.go http/*.go
	go build -o $(TARGET)/orchestrator/orchestrator main.go

$(TARGET)/executor/executor: executor/*.go executor/cmd/*.go
	go build -o $(TARGET)/executor/executor executor/cmd/main.go

$(TARGET)/certs/orchestrator.pem: $(TARGET)/certs/orchestrator_key.pem

$(TARGET)/certs/orchestrator_key.pem: $(TARGET)/generate_cert
	mkdir -p $(TARGET)/certs && \
	cd $(TARGET)/certs && \
	../generate_cert -ca -host 127.0.0.1 -target orchestrator

$(TARGET)/certs/executor.pem: $(TARGET)/certs/executor_key.pem

$(TARGET)/certs/executor_key.pem: $(TARGET)/generate_cert
	mkdir -p $(TARGET)/certs && \
	cd $(TARGET)/certs && \
	../generate_cert -ca -host 127.0.0.1 -target executor 

certificates: $(TARGET)/certs/orchestrator_key.pem $(TARGET)/certs/executor_key.pem

install_certificates: certificates $(TARGET)/orchestrator/orchestrator $(TARGET)/executor/executor
	cp $(TARGET)/certs/orchestrator*pem $(TARGET)/certs/executor.pem $(TARGET)/orchestrator && \
	cp $(TARGET)/certs/executor*pem $(TARGET)/certs/orchestrator.pem $(TARGET)/executor 
	

dist: $(TARGET)/executor/executor $(TARGET)/orchestrator/orchestrator $(TARGET)/generate_cert $(TARGET)/clients/web install_certificates clients

clean:
	rm -rf $(TARGET)

testing: dist
	# Creating the layout
	tmux new-window -n "Gorchestrator"
	tmux split-window -h 
	tmux split-window -v
	tmux select-pane -t 0
	tmux split-window -v
	tmux select-pane -t 0
	tmux resize-pane -D 10
	tmux send-keys "cd $(TARGET)/orchestrator && ./orchestrator" 
	tmux select-pane -t 3
	tmux send-keys "cd $(TARGET)/clients/web && ./webclient"
	tmux select-pane -t 2
	tmux send-keys "cd $(TARGET)/executor && ./executor"
	tmux select-pane -t 1
	tmux send-keys "unset http_proxy && cd clients/tosca/test && cat topology2.yaml | ../../../$(TARGET)/clients/tosca/tosca2gorch | curl  -X POST -H 'Content-Type:application/json' -H 'Accept:application/json' -d@- http://localhost:8080/v1/tasks" 
