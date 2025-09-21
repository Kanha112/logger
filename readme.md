# logger - toggleable logging in Go!

**logger** was developed to enable toggleable logging in your Go program via a simple bool flag. The package just wraps the standard [log](https://pkg.go.dev/log) package provided by Go.

## Usage

1 - add `"github.com/cqhudson/logger"` to your package imports list

2 - instantiate a new logger variable with `logger.NewLogger(bool)`, passing in `true` or `false` to enable or disable logging

```go
package main

import (
	"github.com/cqhudson/logger"
)	

func main() {
	
	// Instantiate a new logger, passing in true to enable logging
	logger := logger.NewLogger(true)


    // log.Print
	logger.Print("This is a test of Print")


    // log.Printf
	s := "Printf"
	logger.Printf("This is a test of %s", s)


    // log.Println
	logger.Println("This is a test of Println")


    // log.Panic
	logger.Panic("This is a test of Panic")

    // log.Panicf
	s = "Panicf"
	logger.Panicf("This is a test of %s", s)


    // log.Panicln
	logger.Panicln("This is a test of Panicln")


    // log.Fatal
	logger.Fatal("This is a test of Fatal")


    // log.Fatalf
	s = "Fatalf"
	logger.Fatalf("This is a test of %s", s)


    // log.Fatalln
	logger.Fatalln("This is a test of Fatalln")
}
```

## Why use logger over Go's builtin log package?

logger can be toggled using a boolean value. This allows for a cleaner codebase with less code duplication, with an easy way to toggle it on or off.

### Example:

Logging with Go's log package:

```go
shouldLog := true

if shouldLog == true {
  log.Print("About to do something")
}

...
do something
do something
do something
...

if shouldLog == true {
  log.Print("Finished doing something")
}
```

Logging with logger:

```go
shouldLog := true
logger := logger.NewLogger(shouldLog)

logger.Print("About to do something")

...
do something
do something
do something
...

logger.Print("Finished doing something")
```
