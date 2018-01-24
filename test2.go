package main

import (
	"bufio"
	"conversions/convhex"
	"fmt"
	"os"
)

/*type decStr struct {
	decodestr map[string]int
	key       string
}*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	inarray := make([]string, 60)
	var instr string
	// Open File to read strings
	file, err := os.Open("60charstr.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var success bool
	success = true

	for success != false {
		success = scanner.Scan()
		err = scanner.Err()
		if success == true {
			if err == nil {
				inarray = append(inarray, scanner.Text())
			} else {
				fmt.Println("ERR #1 reading file")
			}

		} else {
			if err == nil {
				fmt.Println("EOF reached")
			}
		}
	}

	result := make(map[convhex.DecStrKey]convhex.DecStr)
	oneresult := make(map[convhex.DecStrKey]convhex.DecStr)

	for i := 0; i < len(inarray); i++ {
		instr = inarray[i]
		oneresult = convhex.Xorstring(instr)
		for k, v := range oneresult {
			result[k] = v
		}
	}

	// find high score
	highscore := 0
	possdecode := make(map[convhex.DecStrKey]convhex.DecStr)
	var tempcountkey convhex.DecStr

	for key, value := range result {
		if value.Count > highscore {
			highscore = value.Count
			tempcountkey.Count = value.Count
			tempcountkey.Key = value.Key
			possdecode[key] = tempcountkey
		}
		tempcountkey.Count = 0
		tempcountkey.Key = ""
	}

	// look for the high score string and Print
	fmt.Println("The string(s) with the High Score is:")
	for key, value := range possdecode {
		if value.Count == highscore {
			fmt.Println(key.Decodedstr, " = ", value.Count, "Key = ", value.Key)
		}
	}

}
