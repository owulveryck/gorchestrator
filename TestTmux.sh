#!/bin/sh

# Creating the layout
tmux new-window -n "Gorchestrator"
tmux split-window -h 
tmux split-window -v


tmux select-pane -t 0
tmux split-window -v
tmux select-pane -t 0
tmux resize-pane -D 10

tmux send-keys "go run main.go" 
tmux select-pane -t 3
tmux send-keys "cd clients/web && go run *.go" 
tmux select-pane -t 2
tmux send-keys "cd executor && go run cmd/*.go" 

tmux select-pane -t 1
#tmux send-keys "cd example && curl  -X POST -H 'Content-Type:application/json' -H 'Accept:application/json' -d @example.json http://localhost:8080/v1/tasks" 
tmux send-keys "unset http_proxy && cd clients/tosca/test && cat topology2.yaml | go run ../tosca2gorch.go | curl  -X POST -H 'Content-Type:application/json' -H 'Accept:application/json' -d@- http://localhost:8080/v1/tasks" 
