# logger - toggleable logging in Go!

**logger** was developed to enable toggleable logging in your Go program via a simple bool flag. The package just wraps the standard [log](https://pkg.go.dev/log) package provided by Go.

## Usage

1 - add `"github.com/cqhudson/logger"` to your package imports list

2 - instantiate a new logger variable with `logger.NewLogger(enablePrint, enablePanic, enableFatal)`, passing in `true` or `false` to enable or disable different logging types.

```go
package main

import (
    "github.com/cqhudson/logger"
)   

func main() {
    
    // Instantiate a new logger using NewLogger(enablePrint, enablePanic, enableFatal)
    l := logger.NewLogger(true, false, true)

    // Set a prefix
    const prefix string = "[main function]"
    l.SetPrefix(prefix)

    // Set some flags
    l.SetFlags(logger.Lshortfile | logger.Ltime | logger.Lmicroseconds | logger.Lmsgprefix)

    // log.Print
    l.Print("This is a test of Print")


    // log.Printf
    s := "Printf"
    l.Printf("This is a test of %s", s)


    // log.Println
    l.Println("This is a test of Println")


    // log.Panic
    l.Panic("This is a test of Panic")


    // log.Panicf
    s = "Panicf"
    l.Panicf("This is a test of %s", s)


    // log.Panicln
    l.Panicln("This is a test of Panicln")


    // log.Fatal
    l.Fatal("This is a test of Fatal")


    // log.Fatalf
    s = "Fatalf"
    l.Fatalf("This is a test of %s", s)


    // log.Fatalln
    l.Fatalln("This is a test of Fatalln")
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
logPrints := true
logPanics := false
logFatals := false
l := logger.NewLogger(logPrints, logPanics, logFatals)

l.Print("About to do something")

...
do something
do something
do something
...

l.Print("Finished doing something")
```
