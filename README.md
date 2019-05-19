# gostats

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/gostats)](https://goreportcard.com/report/github.com/andygeiss/gostats)

Collect and print statistics about a specific Golang source directory.

**Features**

* Complexity (cyclomatic): Measures the number of linearly independent paths rough a source code.

* Nodes (syntax tree): Measures the number of nodes created in the abstract syntax tree of a source code.

* Node Complexity Rate: Measures the number of nodes per complexity.

##### Table of Contents

- [Installation](README.md#installation)
    * [From Source](README.md#from-source)
- [Usage](README.md#usage)

## Installation

### From Source

    go get -u github.com/andygeiss/gostats
    go install github.com/andygeiss/gostats

## Usage

    gostats help
    
```
Print statistics about a specific Golang source directory.
Complete documentation is available at https://www.github.com/andygeiss/gostats

Usage:
  gostats [flags]
  gostats [command]

Available Commands:
  complexity  Measures the number of linearly independent paths rough a source code
  help        Help about any command
  ncr         Measures the number of nodes per complexity
  nodes       Measures the number of nodes created in the abstract syntax tree of a source code
  table       Print stats about a given Golang source directory as a ASCII table
  version     Print the version identifier

Flags:
  -h, --help   help for gostats

Use "gostats [command] --help" for more information about a command.

```
