package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

const (
	allDigit          = 10
	allAlpha          = 26
	unicodeDigitBegin = 48
	unicodeUpperBegin = 65
	unicodeLowerBegin = 97
)

func main() {

	var tmp []byte
	var input, output string
	var co string
	var keyNum int
	var err error

	flag.Parse()
	co = flag.Args()[0]
	input = flag.Args()[1]
	output = flag.Args()[2]

	tmp, err = ioutil.ReadFile(input)
	checkError(err)

	if co == "en" || co == "encode" { // en, encode 符号化
		keyNum = makeKey(output)
		code(keyNum, tmp)
	} else if co == "de" || co == "decode" { // de, decode 暗号化
		keyNum = makeKey(input)
		code(-keyNum, tmp)
	} else {
		fmt.Println("err: argument")
	}

	err = ioutil.WriteFile(output, tmp, os.ModePerm)
	checkError(err)

}

func makeKey(fileName string) int {

	keyNum := 0
	nameChars := []byte(fileName)
	for i := range nameChars {
		if unicode.IsLetter(rune(nameChars[i])) || unicode.IsDigit(rune(nameChars[i])) {
			keyNum += int(nameChars[i])
		}
	}
	return keyNum % allAlpha

}

func code(keyNum int, data []byte) {

	for i := range data {
		if unicode.IsLower(rune(data[i])) {
			data[i] = (data[i]-unicodeLowerBegin+byte(keyNum))%allAlpha + unicodeLowerBegin
		} else if unicode.IsUpper(rune(data[i])) {
			data[i] = (data[i]-unicodeUpperBegin+byte(keyNum))%allAlpha + unicodeUpperBegin
		} else if unicode.IsDigit(rune(data[i])) {
			data[i] = (data[i]-unicodeDigitBegin+byte(keyNum))%allDigit + unicodeDigitBegin
		}
	}

}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}
