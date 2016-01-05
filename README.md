[![Build Status](https://travis-ci.org/owulveryck/gorchestrator.svg?branch=master)](https://travis-ci.org/owulveryck/gorchestrator)

# Abstract

A simple orchestrator that takes and adjacency matrix as input.

This orchestrator acts as a webservice.
This means that you send a representation of your graph and nodes via an HTTP POST request to the engine, and:

* It decomposes the workflow
* Launch as many "processes" (actually goroutines) as nodes (see performances)
* Launch a conductor that acts as a communication vector for the running nodes

OaaS is micro service oriented. Which means that it does not actually run the artifact of the node. Instead, It calls another web service that acts as a proxy for the execution task. The proxy may implement drivers as needed, such as a `shell` driver, an `ansible` driver, `docker`, ...

The concurrency is implemented thanks to go routines (See this [post](http://blog.owulveryck.info/2015/12/02/orchestrate-a-digraph-with-goroutine-a-concurrent-orchestrator/) for more information about the implementation)

# Architecture

## orchestrator

The orchestrator is the main web service. It is a cloud native, stateless application. Its goal is simply to orchestrate/schedule the tasks so they are executed concurently in the correct order.

## executor(s)

The executor is a web service that actually executes a task.

## The Gorch

`gorch` is the Graphical-Orchestration representation.

It is a JSON representation of the graph.

It is composed of and adjaceny matrix and a list of nodes.

```JSON
{
    "name": "string",
    "state": 0,
    "digraph": [
        0
    ],
    "nodes": [
        {
            "id": 0,
            "state": 0,
            "name": "string",
            "engine": "string",
            "artifact": "string",
            "args": [
                "string"
            ]
        }
    ]

}
```

# Getting it up and running

The engine is written in pure go. The package is go-gettable. Assuming you have a GO environment up and running, the following tasks should be enough to enjoy the orchestrator:

* `go get github.com/owulveryck/gorchestrator`
* `cd $GOPATH/src/github.com/owulveryck/gorchestrator`
* `go run`

Then you can post a query as described in the example folder:

```shell
# curl -X POST -H 'Content-Type:application/json' -H 'Accept:application/json' -d @example.json -k http://localhost:8080/v1/tasks
```

# API

The REST API is in developement but nearly stable. It is self documented with swagger. 

## Apidoc

The api doc is viewable

* live [here](http://blog.owulveryck.info/gorchestrator/swagger/) for api documentation.
* In your own instance at [http://localhost:8080/apidocs/](http://localhost:8080/apidocs/)

# Clients
 
## Webclient

A web client is in development, see _clients/web_ for the sources:
![Screenshot](https://raw.githubusercontent.com/owulveryck/gorchestrator/master/Documentation/images/webclient.png)

## Tosca client

A tosca client using the [toscalib](https://github.com/owulveryck/toscalib) is present in _clients/tosca_.
It converts a TOSCA execution plan into a _gorch_ representation.
