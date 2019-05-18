# check

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/check)](https://goreportcard.com/report/github.com/andygeiss/check)

This package provides a simple error handling by logging the error message to stderr.

##### Table of Contents

- [Installation](README.md#installation)
    * [From Source](README.md#from-source)
- [Usage](README.md#usage)

## Installation

### From Source

    go get -u github.com/andygeiss/check

## Usage

The following code represents a traditional "happy path" with its error handling.

```go
file, err := os.Open("filename.json")
if err != nil {
	log.Fatal(err.Error())
}
defer file.Close()

in, err := ioutil.ReadAll(file)
if err != nil {
	log.Fatal(err.Error())
}

var data struct{}
err := json.Unmarshal(in, &data)
if err != nil {
	log.Fatal(err.Error())
}
```

The problem is that the error handling is bad for readability, too.
With <code>check</code> the above code looks like:

```go
file, err := os.Open("filename.json")
check.Fatal(err)
defer file.Close()

in, err := ioutil.ReadAll(file)
check.Fatal(err)

var data struct{}
err := json.Unmarshal(in, &data)
check.Fatal(err)
```