package colourize

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	fmt.Println("~~~* Testing Text Colour *~~~")
	fmt.Println(Colourize("Black Text", Black))
	fmt.Println(Colourize("Red Text", Red))
	fmt.Println(Colourize("Green Text", Green))
	fmt.Println(Colourize("Yellow Text", Yellow))
	fmt.Println(Colourize("Blue Text", Blue))
	fmt.Println(Colourize("Magenta Text", Magenta))
	fmt.Println(Colourize("Cyan Text", Cyan))
	fmt.Println(Colourize("White Text", White))
	fmt.Println(Colourize("Grey Text", Grey))
}

func TestBackground(t *testing.T) {
	fmt.Println("~~~* Testing Background Colour *~~~")
	fmt.Println(Colourize("Black Background", Blackbg))
	fmt.Println(Colourize("Red Background", Redbg))
	fmt.Println(Colourize("Green Background", Greenbg))
	fmt.Println(Colourize("Yellow Background", Yellowbg))
	fmt.Println(Colourize("Blue Background", Bluebg))
	fmt.Println(Colourize("Magenta Background", Magentabg))
	fmt.Println(Colourize("Cyan Background", Cyanbg))
	fmt.Println(Colourize("White Background", Whitebg))
}

func TestStyle(t *testing.T) {
	fmt.Println("~~~* Testing Style *~~~")
	fmt.Println(Colourize("Bold Test", Bold))
	fmt.Println(Colourize("Dim Test", Dim))
	fmt.Println(Colourize("Italic Test", Italic))
	fmt.Println(Colourize("Underline Test", Underline))
	fmt.Println(Colourize("Blinkslow Test", Blinkslow))
	fmt.Println(Colourize("Blinkfast Test", Blinkfast))
	fmt.Println(Colourize("Inverse Test", Inverse))
	fmt.Println(Colourize("Hidden Test", Hidden))
	fmt.Println(Colourize("Strikeout Test", Strikeout))
}

func TestFunctional(t *testing.T) {
	fmt.Println("~~~* Testing Functional Uses *~~~")
	fmt.Println(Colourize("Bold White and Green Background", Bold, White, Greenbg))
	fmt.Println(Colourize("Dim White with Cyan Background and ", White, Cyanbg, Dim), "Testing Automatic Style Reset")
	fmt.Println(Colourize("White background with Blue Text and", Whitebg, Blue), Colourize("Green Background and Yellow Text", Greenbg, Yellow))
	testString := "with a Printf test"
	fmt.Printf(Colourize("Blue Background and White Text %s", Bluebg, White), testString)
	fmt.Printf("\n")
	fmt.Println(Colourize(1234, Green), "Testing with Integer Input")
}
