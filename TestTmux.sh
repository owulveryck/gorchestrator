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
