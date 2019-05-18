# assert

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/assert)](https://goreportcard.com/report/github.com/andygeiss/assert)

This package provides a simple assert extension by using <code>reflect</code> kungfu.

##### Table of Contents

- [Installation](README.md#installation)
    * [From Source](README.md#from-source)
- [Usage](README.md#usage)

## Installation

### From Source

    go get -u github.com/andygeiss/assert

## Usage

```go
func Test_Given_Nothing_When_Calling_Foo_Then_Return_42_And_No_Error(t *testing.T) {
    res, err := foo()
    assert.That(t, err, is.Equal(nil))
    assert.That(t, res, is.Equal(42))
}
```
