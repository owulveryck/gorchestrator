cat topology2.yaml| go run ../main.go | go run ../../gorch2dot/gorchestrator2dot.go | dot -Tpng > topology.png
