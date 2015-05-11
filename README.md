# colourize
An ANSI colour terminal package for Go

[![GoDoc](https://godoc.org/github.com/TreyBastian/colourize?status.svg)](https://godoc.org/github.com/TreyBastian/colourize) [![Build Status](https://travis-ci.org/TreyBastian/colourize.svg?branch=master)](https://travis-ci.org/TreyBastian/colourize) [![Coverage Status](https://coveralls.io/repos/TreyBastian/colourize/badge.svg?branch=master)](https://coveralls.io/r/TreyBastian/colourize?branch=master)

# Usage

    package main
    import (
        c "github.com/TreyBastian/colourize"
        "fmt"
    )
    
    func main() {
      fmt.Println(colourize("Hello World!", Green, Whitebg, Bold)
    }
    
Supports all ANSI colours and emphasis. Not compatible with Windows systems.
