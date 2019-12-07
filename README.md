# gostats

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/gostats)](https://goreportcard.com/report/github.com/andygeiss/gostats)
[![BCH compliance](https://bettercodehub.com/edge/badge/andygeiss/gostats?branch=master)](https://bettercodehub.com/)

Collect and print statistics like code complexity about a specific Golang source directory.

**Features**

* Complexity (cyclomatic): Measures the number of linearly independent paths rough a source code.

* Nodes (syntax tree): Measures the number of nodes created in the abstract syntax tree of a source code.

* Node Complexity Rate: Measures the number of nodes per complexity, which indicates good balance between control structure and simple statements.

**Important**

* Files in subdirectory "testdata" will be ignored.
* Files with suffix "_test.go" will be ignored.

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

## Example

Use the following command to create statistics about by itself:

    gostats table .

The table prints the statistics per file and in total.    
```
+-----------------------------------+-------+------------+-----+
|               IDENT               | NODES | COMPLEXITY | NCR |
+-----------------------------------+-------+------------+-----+
| cmd/complexity.go                 |  84   |     3      | 28  |
| cmd/ncr.go                        |  104  |     3      | 34  |
| cmd/nodes.go                      |  84   |     3      | 28  |
| cmd/root.go                       |  56   |     2      | 28  |
| cmd/table.go                      |  324  |     4      | 81  |
| cmd/utils.go                      |  134  |     3      | 44  |
| cmd/version.go                    |  22   |     1      | 22  |
| configs/appInfo.go                |   0   |     1      |  0  |
| internal/stats/service.default.go |  302  |     7      | 43  |
| internal/stats/service.go         |   0   |     1      |  0  |
| main.go                           |  20   |     1      | 20  |
| pkg/complexity/fromSource.go      |  324  |     11     | 29  |
| pkg/files/excludeByPrefix.go      |  78   |     3      | 26  |
| pkg/files/excludeBySuffix.go      |  78   |     3      | 26  |
| pkg/files/filterByExtension.go    |  76   |     3      | 25  |
| pkg/files/fromPath.go             |  132  |     3      | 44  |
| pkg/nodes/fromSource.go           |  200  |     6      | 33  |
+-----------------------------------+-------+------------+-----+
|               TOTAL               | 2018  |     58     | 34  |
+-----------------------------------+-------+------------+-----+
```
