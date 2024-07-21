package main

import (
	"ascii/art/font"
	"fmt"
    "strings"
)

func main() {
    str := "Maks Loh"

    str = strings.ToUpper(str)

    for i := range font.SimpleFont['A'] { 
        for _, char := range str {
            if char == ' ' {
                fmt.Print("     ")
            } else {
                charLines, exists := font.SimpleFont[char]
                if !exists {
                    charLines = []string{
                        "     ",
                        "     ",
                        "     ",
                        "     ",
                        "     ",
                    }
                }
                fmt.Print(charLines[i] + " ")
            }
        }
        fmt.Println()
    }
}
