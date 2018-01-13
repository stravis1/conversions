package main

//import "os"
import "fmt"

func main() {
	// using this inline assignment for now to test the conversion
	// will replace with a parameter to read a file
	buf1 := "1c0111001f010100061a024b53535009181c"
	buf2 := "686974207468652062756c6c277320657965"

	/*
		var hexenc = "0123456789abcdef"  // used for encoding to hex
	*/
	var hexDec = map[byte]string{'0': "0000", '1': "0001", '2': "0010",
		'3': "0011", '4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'a': "1010", 'b': "1011", 'c': "1100",
		'd': "1101", 'e': "1110", 'f': "1111"}

	var hexEnc = map[string]string{"0000": "0", "0001": "1", "0010": "2",
		"0011": "3", "0100": "4", "0101": "5", "0110": "6", "0111": "7",
		"1000": "8", "1001": "9", "1010": "a", "1011": "b", "1100": "c",
		"1101": "d", "1110": "e", "1111": "f"}

	// Create a binary string out of the hex data for each buffer
	var bufbin1str string
	var bufbin2str string

	lbuf := len(buf1)

	// buf1
	for i := 0; i < lbuf; i++ {
		bufbin1str = bufbin1str + hexDec[buf1[i]]
	}

	// buf2
	for i := 0; i < lbuf; i++ {
		bufbin2str = bufbin2str + hexDec[buf2[i]]
	}

	lbufstr := len(bufbin1str)

	var xorbin string
	// Xor bitwise compare
	for i := 0; i < lbufstr; i++ {
		if bufbin1str[i] != bufbin2str[i] {
			xorbin = xorbin + "1"
		} else {
			xorbin = xorbin + "0"
		}

	}
	fmt.Println("output Xor = ", xorbin)

	// Encode the binary string for hex output
	var xorout string
	var xortemp string

	for i := 0; i < len(xorbin); i = i + 4 {
		for x := i; x <= (i + 3); x++ {
			xortemp = xortemp + string(xorbin[x])
		}
		xorout = xorout + hexEnc[xortemp]
		xortemp = ""
	}

	fmt.Println("xor output = ", xorout)

}
