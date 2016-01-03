//Package colourize implements simple ANSI colour codes to style terminal output text.
package colourize

import (
	"bytes"
	"fmt"
	"strconv"
)

//Colour Styles
const (
	//Text Colour
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
	Grey    = 90

	//Background Colour
	Blackbg   = 40
	Redbg     = 41
	Greenbg   = 42
	Yellowbg  = 43
	Bluebg    = 44
	Magentabg = 45
	Cyanbg    = 46
	Whitebg   = 47

	//Style
	Bold      = 1
	Dim       = 2
	Italic    = 3
	Underline = 4
	Blinkslow = 5
	Blinkfast = 6
	Inverse   = 7
	Hidden    = 8
	Strikeout = 9
)

/*Colourize is the function that provides colours for values passed to it.
package main

import(
	c "github.com/TreyBastian/colourize"
	"fmt"
)

func main() {
	fmt.Println(c.Colourize("Hello", c.Cyan))
}
*/
func Colourize(s interface{}, style ...int) string {
	b := new(bytes.Buffer)
	b.WriteString("\x1b[")
	l := len(style)
	for k, c := range style {
		b.WriteString(strconv.Itoa(c))
		if l > 1 && k < l-1 {
			b.WriteString(";")
		}
	}
	b.WriteString("m")

	return fmt.Sprintf("%s%v\x1b[0m", b.String(), s)
}
