# colourize
[![GoDoc](https://godoc.org/github.com/TreyBastian/colourize?status.svg)](https://godoc.org/github.com/TreyBastian/colourize) [![Build Status](https://travis-ci.org/TreyBastian/colourize.svg?branch=master)](https://travis-ci.org/TreyBastian/colourize) [![Coverage Status](https://coveralls.io/repos/TreyBastian/colourize/badge.svg?branch=master)](https://coveralls.io/r/TreyBastian/colourize?branch=master) [![Go Report
Card](https://goreportcard.com/badge/github.com/TreyBastian/colourize)](https://goreportcard.com/report/github.com/TreyBastian/colourize)

An ANSI colour terminal package for Go.
Supports all ANSI colours and emphasis. Not compatible with Windows systems.

# Installation

     go get github.com/TreyBastian/colourize 

# Usage

    package main
    import (
        c "github.com/TreyBastian/colourize"
        "fmt"
    )
    
    func main() {
      fmt.Println(c.Colourize("Hello World!", c.Green, c.Whitebg, c.Bold))
    }


# Projects Using colourize
*  [nettee](https://github.com/Manaphy91/nettee)
