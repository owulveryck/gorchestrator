# Abstract

This library is an implementation of the TOSCA definition as described in the document written in pure GO
[TOSCA Simple Profile in YAML Version 1.0](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html)

## Normative Types
The normative types definitions are included de facto. The files are embeded using go-bindata.

# Howto

Create a `ServiceTemplateDefinition` and call `Parse(r io.Reader)` to fill it with a YAML definition.

## Example

```go
var t toscalib.ServiceTemplateDefinition
err := t.Parse(os.Stdin)
if err != nil {
    log.Fatal(err)
}
```

## Projects

* [gorchestrator](https://github.com/owulveryck/gorchestrator) is implementing thos toscalib for one of its client.

# Status

## Test
The basic tests function are taking all the examples of the standard and try to parse them.
No verification is done, but by now, I don't have any error in the parsing of any file.

### Coverage
```shell
github.com/owulveryck/toscalib/capabilities.go:14:              UnmarshalYAML           94.1%
github.com/owulveryck/toscalib/constraints.go:9:                IsValid                 0.0%
github.com/owulveryck/toscalib/constraints.go:24:               Evaluate                0.0%
github.com/owulveryck/toscalib/constraints.go:27:               UnmarshalYAML           90.9%
github.com/owulveryck/toscalib/interfaces.go:22:                UnmarshalYAML           90.9%
github.com/owulveryck/toscalib/interfaces.go:55:                UnmarshalYAML           90.9%
github.com/owulveryck/toscalib/node_template.go:29:             setRefs                 87.5%
github.com/owulveryck/toscalib/node_template.go:45:             getInterface            100.0%
github.com/owulveryck/toscalib/node_template.go:54:             fillInterface           94.1%
github.com/owulveryck/toscalib/node_template.go:82:             setName                 100.0%
github.com/owulveryck/toscalib/node_type.go:21:                 getInterface            100.0%
github.com/owulveryck/toscalib/parser.go:14:                    GetNodeTemplate         0.0%
github.com/owulveryck/toscalib/parser.go:24:                    merge                   86.0%
github.com/owulveryck/toscalib/parser.go:92:                    Parse                   67.3%
github.com/owulveryck/toscalib/properties.go:28:                UnmarshalYAML           72.2%
github.com/owulveryck/toscalib/requirements.go:12:              UnmarshalYAML           93.3%
github.com/owulveryck/toscalib/service_template.go:27:          Bytes                   0.0%
github.com/owulveryck/toscalib/service_template.go:33:          String                  0.0%
github.com/owulveryck/toscalib/tosca_namespace_alias.go:96:     UnmarshalYAML           0.0%
```

 
# API
[![GoDoc](https://godoc.org/github.com/owulveryck/toscalib?status.svg)](https://godoc.org/github.com/owulveryck/toscalib)

# Legacy

This API is in complete rewrite, for the old version, please checkout the "v1" branch
