package main

import (
	"conversions/convhex"
	"fmt"
)

/*
func check(e error) {
	if e != nil {
		panic(e)
	}
}
*/

func main() {

	result2 := convhex.Xormultstr("60charstr.txt")

	possdecode := convhex.Findhighscore(result2, 8)

	// look for the high score string and Print
	fmt.Println("The string(s) with the High Score is:")
	for key, value := range possdecode {
		//if value.Count == highscore {
		fmt.Println(key.Decodedstr, " = ", value.Count, "Key = ", value.Key)
		//}
	}

}
