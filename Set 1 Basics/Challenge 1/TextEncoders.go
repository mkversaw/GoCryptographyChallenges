package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var base64Array []string
var base64IndexMap = make(map[string]int)

func InitBase64() {
	idxCounter := 0

	// start with alphabet and numbers
	for i := 'A'; i <= 'Z'; i++ {
		base64Array = append(base64Array, string(i))
		base64IndexMap[string(i)] = idxCounter
		idxCounter++
	}

	for i := 'a'; i <= 'z'; i++ {
		base64Array = append(base64Array, string(i))
		base64IndexMap[string(i)] = idxCounter
		idxCounter++
	}

	for i := '0'; i <= '9'; i++ {
		base64Array = append(base64Array, string(i))
		base64IndexMap[string(i)] = idxCounter
		idxCounter++
	}

	// add the two symbols for 62 and 63
	base64Array = append(base64Array, string('+'))
	base64IndexMap[string('+')] = idxCounter
	idxCounter++
	base64Array = append(base64Array, string('/'))
	base64IndexMap[string('/')] = idxCounter
	idxCounter++
}

func DecimalToHex(decimalNum int) string {
	if decimalNum == 0 {
		return "0"
	}

	hexDigits := "0123456789abcdef"
	hexString := ""

	for ; decimalNum > 0; decimalNum /= 16 {
		remainder := decimalNum % 16
		hexString = string(hexDigits[remainder]) + hexString
	}

	return hexString
}

func HexToDecimal(hexString string) int {
	hexString = strings.ToUpper(hexString)
	hexChars := "0123456789ABCDEF"
	decimal := 0

	if(len(hexString) > 10) {
		println("Length of string was too large")
		return -1
	}

	// each character of hexString, left to right
	for _, char := range hexString {
		digit := strings.Index(hexChars, string(char))

		if digit == -1 {
			// Invalid hex char
			println("Invalid hex char found")
			return -1
		}

		// multiply the previous result by 16 and add the current digit
		// this effectively multiplies each digit by 16^(the digits index from the right)
		// while also calculating the sum of these products
		decimal = (decimal * 16) + digit
	}

	return decimal
}

func DecimalToBinary(decimalNum int) string {
	//if decimalNum == 0 {
	//	return "0"
	//}

	output := ""

	for ; decimalNum > 0; decimalNum /= 2 {
		output = fmt.Sprint(decimalNum%2) + output
	}

	// pad output to be of form 0000
	for len(output) < 4 {
		output = "0" + output
	}

	return output
}

func BinaryToDecimal(binaryString string) int {
	decimal := 0

	// each character of binaryString, left to right
	for _, char := range binaryString {
		// convert the '0' or '1' char to its numeric value
		digit := int(char - '0')

		if digit != 0 && digit != 1 {
			// Invalid binary value
			return -1
		}

		// multiply the previous result by 2 and add the current digit
		// this effectively multiplies each digit by 2^(the digits index from the right)
		// while also calculating the sum of these products
		decimal = (decimal * 2) + digit
	}

	return decimal
}

func BinaryToBase64(binaryString string) string {
	paddingToAdd := 0

	if(len(binaryString) % 6 != 0) {
		paddingToAdd = 6 - (len(binaryString) % 6)
	}

	// pad with 0's so string can be split into sextets
	for i := 0; i < paddingToAdd; i++ {
		binaryString += "0"
	}

	newLen := len(binaryString)
	output := ""

	for i := 0; i + 6 <= newLen; i += 6 {
		b64Index := BinaryToDecimal(binaryString[i:i+6])
		output += base64Array[b64Index]
	}


	paddingToAdd2 := 0
	lenInBytes := ((len(binaryString) - paddingToAdd)) / 8

	if(lenInBytes % 3 != 0) {
		paddingToAdd2 = 3 - (lenInBytes % 3)
	}

	for i := 0; i < paddingToAdd2; i++ {
		output += "="
	}

	return output
}

func Base64ToBinary(b64String string) string {
	output := ""

	for i := 0; i < len(b64String); i++ {
		if(b64String[i] == '=') {
			// remove 2 bits for each padding character at end
			output = output[:len(output)-2]
		} else {
			b64Index := base64IndexMap[string(b64String[i])]
			binaryString := DecimalToBinary(b64Index)

			for len(binaryString) < 6 {
				binaryString = "0" + binaryString
			}

			output += binaryString
		}
	}
	return output
}

func ASCIIToHex(asciiString string) string {

	output := ""

	for i := 0; i < len(asciiString); i++ {
		output += DecimalToHex(int(asciiString[i]))
	}

	return output
}

func HexToAscii(hexString string) string {
	if(len(hexString) % 2 != 0) {
		print("hex string does not represent ascii text")
		return "-1"
	}

	output := ""

	for i := 0; i + 2 <= len(hexString); i += 2 {
		decimalNum := HexToDecimal(hexString[i:i+2])
		output += string(byte(decimalNum))
	}

	return output
}

func BinaryToAscii(binaryString string) string {

	output := ""

	for i := 0; i + 8 <= len(binaryString); i += 8 {
		decimalNum := BinaryToDecimal(binaryString[i:i+8])
		output += string(byte(decimalNum))
	}

	return output
}

func main() {

	InitBase64()

	fmt.Println("input ascii text:")

	var asciiText string

	reader := bufio.NewReader(os.Stdin)
	asciiText, err := reader.ReadString('\r')
	if err != nil {
		println("Invalid ascii text!")
		return
	}

	hexText := ASCIIToHex(asciiText)

	println("Converted to hex:")
	println(hexText)

	binaryText := ""
	for i := 0; i < len(hexText); i++ {
		decimalNum := HexToDecimal(string(hexText[i]))
		binaryText += DecimalToBinary(decimalNum)
	}

	println("Converted to binary:")
	println(binaryText)




	base64Text := BinaryToBase64(binaryText)
	println("Converted to Base64:")
	println(base64Text)

	binaryText = Base64ToBinary(base64Text)
	asciiText = BinaryToAscii(binaryText)

	println("Back to ascii:")
	println(asciiText)
}
