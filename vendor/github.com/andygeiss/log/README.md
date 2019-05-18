# log

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/log)](https://goreportcard.com/report/github.com/andygeiss/log)

This package provides a simple log extension for the default API in <code>log</code>.

##### Table of Contents

- [Installation](README.md#installation)
    * [From Source](README.md#from-source)
- [Usage](README.md#usage)

## Installation

### From Source

    go get -u github.com/andygeiss/log

## Usage

The following example sets the Log-Level from Info to Warn and prints a few log messages to stdout.

```go
log.Level = log.LevelWarn // switch from Info to Warn
log.Debug("4 CPUs found.")
log.Info("Starting ...")
log.Warn("No parameters set!")
log.Error("Action failed! Shutdown is recommended!")
log.Fatal("Action failed! Shutting down...")
log.Warn("Never reached.")
```

The output will be:
    
    2018/11/26 16:16:14 WARN  | 929ns            | No parameters set!                      
    2018/11/26 16:16:14 ERROR | 194.941µs        | Action failed! Shutdown is recommended! 
    2018/11/26 16:16:14 FATAL | 134.146µs        | Action failed! Shutting down...
    