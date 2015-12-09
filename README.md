# Abstract

A simple orchestrator that takes and adjacency matrix as input.

This orchestrator is a webservice.
The concurrency is implemented thanks to go routines (See this [post](http://blog.owulveryck.info/2015/12/02/orchestrate-a-digraph-with-goroutine-a-concurrent-orchestrator/) for more information about the implementation)


# Getting it up and running

```
go get github.com/owulveryck/gorchestrator
```

then 

```
go run main.go
```

It should start a http server listening on port 8080.

# API

The api is in developement and will be documented with swagger. The APIdoc will be included in the distribution.

To see the API documentation, please go to [http://localhost:8080/apidocs/](http://localhost:8080/apidocs/)

# AuthN and AuthZ

The implementation is in the roadmap, and will be based on `oauth2`

# Performances

I did a `test` file to bench the orchestrator engine with the `go` mechanism. The example, is the simple one listed one my blog post.

Here are the results on my chromebook (which is small, with only 2 Gb of RAM)

```
orchestrator git:(master)go test -bench . -cpu 1
PASS
BenchmarkRun        3000            981423 ns/op
ok      github.com/owulveryck/gorchestrator/orchestrator        3.882s
```

Which means that I can interpret and run 3000 times this digraph in 3.8s (excluding the actual execution time of the task)

__Note__ : It is simply the execution workflow as all the nodes do not perform any action.
