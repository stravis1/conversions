package convhex

import "fmt"

func Hex2base64(hex string) string {

	fmt.Println("Hex string = ", hex)

	hexbinstr := Hex2bin(hex)
	// print output to see binary data
	/*
		fmt.Println("Binary string = ", hexbinstr)
		fmt.Println("Binary string length = ", len(hexbinstr))
	*/

	var base64output string
	base64output = Bin2base64(hexbinstr)
	return (base64output)

}

func Hex2bin(hex string) string {
	var hexDec = map[byte]string{'0': "0000", '1': "0001", '2': "0010",
		'3': "0011", '4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'a': "1010", 'b': "1011", 'c': "1100",
		'd': "1101", 'e': "1110", 'f': "1111"}

	// Create a binary string out of the hex data
	var hexbinstr string

	for i := 0; i < len(hex); i++ {
		hexbinstr = hexbinstr + hexDec[hex[i]]
	}
	return (hexbinstr)
}

func Bin2base64(hexbinstr string) string {
	var base64enc = map[string]string{"000000": "A", "000001": "B", "000010": "C",
		"000011": "D", "000100": "E", "000101": "F", "000110": "G", "000111": "H",
		"001000": "I", "001001": "J", "001010": "K", "001011": "L", "001100": "M",
		"001101": "N", "001110": "O", "001111": "P", "010000": "Q", "010001": "R",
		"010010": "S", "010011": "T", "010100": "U", "010101": "V", "010110": "W",
		"010111": "X", "011000": "Y", "011001": "Z", "011010": "a", "011011": "b",
		"011100": "c", "011101": "d", "011110": "e", "011111": "f", "100000": "g",
		"100001": "h", "100010": "i", "100011": "j", "100100": "k", "100101": "l",
		"100110": "m", "100111": "n", "101000": "o", "101001": "p", "101010": "q",
		"101011": "r", "101100": "s", "101101": "t", "101110": "u", "101111": "v",
		"110000": "w", "110001": "x", "110010": "y", "110011": "z", "110100": "0",
		"110101": "1", "110110": "2", "110111": "3", "111000": "4", "111001": "5",
		"111010": "6", "111011": "7", "111100": "8", "111101": "9", "111110": "+",
		"111111": "/"}

	var binBase64 string
	var base64output string

	// if the data string converted to binary is not a multiple of 12 bits.
	// pad the string with '0' to make it a multiple. This will result in the 'A'
	// character being used instead of '='
	if (len(hexbinstr) % 12) != 0 {
		remainder := len(hexbinstr) % 12
		if remainder == 4 {
			hexbinstr = hexbinstr + "00000000"
		} else if remainder == 8 {
			hexbinstr = hexbinstr + "0000"
		} else {
			fmt.Println("Error found on length of input string.")
		}
	}

	// cycle through the hexbinstring and encode 6 bits into base64
	//
	for i := 0; i < len(hexbinstr); i = i + 6 {
		for x := i; x <= (i + 5); x++ {
			binBase64 = binBase64 + string(hexbinstr[x])

		}
		// fmt.Println(binBase64, " ", base64enc[binBase64])

		base64output = base64output + base64enc[binBase64]
		binBase64 = "" //need to clear binBase64 for next loop.

	}

	//fmt.Println("base64output = ", base64output)
	return (base64output)
}

func Bin2hex(binstr string) string {
	var hexEnc = map[string]string{"0000": "0", "0001": "1", "0010": "2",
		"0011": "3", "0100": "4", "0101": "5", "0110": "6", "0111": "7",
		"1000": "8", "1001": "9", "1010": "a", "1011": "b", "1100": "c",
		"1101": "d", "1110": "e", "1111": "f"}

	// Create a binary string out of the hex data
	var hexstr string
	var tempstr string

	for i := 0; i < len(binstr); i = i + 4 {
		for x := i; x <= i+3; x++ {
			tempstr = tempstr + string(binstr[x])
		}
		hexstr = hexstr + hexEnc[tempstr]
		tempstr = ""

	}
	return (hexstr)
}

func Xor(buf1 string, buf2 string) string {

	var hexEnc = map[string]string{"0000": "0", "0001": "1", "0010": "2",
		"0011": "3", "0100": "4", "0101": "5", "0110": "6", "0111": "7",
		"1000": "8", "1001": "9", "1010": "a", "1011": "b", "1100": "c",
		"1101": "d", "1110": "e", "1111": "f"}

	// Create a binary string out of the hex data for each buffer
	var bufbin1str string
	var bufbin2str string

	bufbin1str = Hex2bin(buf1)
	bufbin2str = Hex2bin(buf2)

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
	// fmt.Println("output Xor = ", xorbin)

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

	// fmt.Println("xor output = ", xorout)
	return (xorout)

}
