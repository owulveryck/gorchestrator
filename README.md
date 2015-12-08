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

