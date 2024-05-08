package main

import (
	"fmt"
	"log"
	"strings"
)

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

func HexToBinary(hexString string) string {
	binaryText := ""
	for i := 0; i < len(hexString); i++ {
		decimalNum := HexToDecimal(string(hexString[i]))
		binaryText += DecimalToBinary(decimalNum)
	}
	return binaryText
}

func BinaryToHex(binaryString string) string {
	output := ""

	for i := 0; i + 4 <= len(binaryString); i += 4 {
		decimalNum := BinaryToDecimal(binaryString[i:i+4])
		output += DecimalToHex(decimalNum)
	}

	return output
}

func XORBinary(binaryStr1 string, binaryStr2 string) string {
	if(len(binaryStr1) != len(binaryStr2)) {
		println("XORBinary strings were not equal length")
		return "-1"
	}
	
	output := ""

	for i := range len(binaryStr1) {
		digit1 := int(binaryStr1[i] - '0')
		digit2 := int(binaryStr2[i] - '0')

		// add 1 to output iff number of 1s is even else 0
		if(digit1 == 1 && digit2 != 1) {
			output += "1"
		} else if (digit2 == 1 && digit1 != 1) {
			output += "1"
		} else {
			output += "0"
		}
	}

	return output
}

func main() {

	//hexText := HexToDecimal(1c0111001f010100061a024b53535009181c)
	
    b1, b2 := "1", "1"

	for(b1 != "" && b2 != "") {
		println("Enter 2 hex strings")
		n, err := fmt.Scanln(&b1, &b2)
		if err != nil {
			log.Fatal(err,n)
		} else {	
			println("XORed result: ")
			binaryText := XORBinary(HexToBinary(b1),HexToBinary(b2))
			println(BinaryToHex(binaryText))
		}
	}

}
