package main

import "fmt"
import "conversions/convhex"

/*
var asciiEnc = map[string]string{"61": "a", "62": "b", "63": "c",
	"64": "d", "65": "e", "66": "f", "67": "g", "68": "h", "69": "i",
	"6a": "j", "6b": "k", "6c": "l", "6d": "m", "6e": "n", "6f": "o",
	"70": "p", "71": "q", "72": "r", "73": "s", "74": "t", "75": "u",
	"76": "v", "77": "w", "78": "x", "79": "y", "7a": "z",
	"41": "A", "42": "B", "43": "C", "44": "D", "45": "E", "46": "F",
	"47": "G", "48": "H", "49": "I", "4a": "J", "4b": "K", "4c": "L",
	"4d": "M", "4e": "N", "4f": "O", "50": "P", "51": "Q", "52": "R",
	"53": "S", "54": "T", "55": "U", "56": "V", "57": "W", "58": "X",
	"59": "Y", "5a": "Z", "20": " "}

var asciiDec = map[string]string{"a": "61", "b": "62", "c": "63",
	"d": "64", "e": "65", "f": "66", "g": "67", "h": "68", "i": "69",
	"j": "6a", "k": "6b", "l": "6c", "m": "6d", "n": "6e", "o": "6f",
	"p": "70", "q": "71", "r": "72", "s": "73", "t": "74", "u": "75",
	"v": "76", "w": "77", "x": "78", "y": "79", "z": "7a",
	"A": "41", "B": "42", "C": "43", "D": "44", "E": "45", "F": "46",
	"G": "47", "H": "48", "I": "49", "J": "4a", "K": "4b", "L": "4c",
	"M": "4d", "N": "4e", "O": "4f", "P": "50", "Q": "51", "R": "52",
	"S": "53", "T": "54", "U": "55", "V": "56", "W": "57", "X": "58",
	"Y": "59", "Z": "5a", "0": "30", "1": "31", "2": "32", "3": "33",
	"4": "34", "5": "35", "6": "36", "7": "37", "8": "38", "9": "39",
	" ": "20"}

var charscore = map[string]int{"e": 13, "t": 9, "a": 8, "o": 8,
	"i": 7, "n": 7, "s": 6, "h": 6, "r": 6, "d": 4, "l": 3, "c": 2,
	"u": 1, "E": 13, "T": 9, "A": 8, "O": 8, "I": 7, "N": 7, "S": 6,
	"H": 6, "R": 6, "D": 4, "L": 3, "C": 2, "U": 1, " ": 10}
*/
func main() {
	//buf1 := "1c0111001f010100061a024b53535009181c"
	//buf2 := "686974207468652062756c6c277320657965"
	instr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	//buf2 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

	//var xortemp string
	//var xoresult string
	//var asciout string
	//var charsum int
	score := make(map[string]int)
	/*
		// taking each char in buf2 xor with all hex chars in instr
		for x := 0; x < len(buf2); x++ {
			// two hex char in each char.
			for i := 0; i < len(instr); i = i + 2 {
				xortemp = xortemp + convhex.Xor(string(instr[i:(i+2)]), string(asciiDec[(string(buf2[x]))]))
			}
			xoresult = xortemp
			xortemp = ""
			// fmt.Println("xor = ", xoresult)

			var tempchar string

			//convert hex string back to ascii for output
			for y := 0; y < len(xoresult); y = y + 2 {
				tempchar = asciiEnc[string(xoresult[y:(y+2)])]
				asciout = asciout + tempchar
				charsum = charsum + charscore[tempchar]
				//fmt.Println(xoresult[x:(x + 2)])
				//fmt.Println(asciout)
				// create a map of scores
				/*_, ok := score[tempchar]
				if ok {
					score[tempchar] = score[tempchar] + 1
				} else {
					score[tempchar] = 1
				}

				// find the score

			}  */
	score = convhex.Xorstring(instr)
	/*
			charsum = 0 // reset charsum for next string
			// Print key and output.
			fmt.Println("key = ", (string(buf2[x])))
			fmt.Println("hex out = ", asciout)
			fmt.Println("score = ", score[asciout])
			fmt.Println("")
			asciout = ""

		}
	*/

	// find high score
	highscore := 0
	possdecode := make(map[string]int)

	for key, value := range score {
		if value > highscore {
			highscore = value
			possdecode[key] = value

		}
	}

	// look for the high score string and Print
	fmt.Println("The string(s) with the High Score is:")
	for key, value := range possdecode {
		if value == highscore {
			fmt.Println(key, " = ", value)
		}
	}

}
