package main 

import "fmt"

func main() {
  // empty string defalt
	var str string
	
	// with spec simbols 
	var hello string = "Привет\n\t"

	// without spec simbols
	var world string = `Мир\n\t`

	// UTF-8
	var helloWorld =  "Hello world"
	hi := '马克'

	// for bite 
	var rawBinary byte = '\x27'

	// rune (uint32) for UTF-8
	var someChinese rune = '马'

	helloWorld := "Привет мир" 

	andGoodMorning := helloWorld + " и доброе утро!"

	// strings immutable
	// cannot assign to helloworld[0]
	helloWorld[0] = 72

	// get lenght of string 
	byteLen := len(helloWorld)
	symbols := utf8.RuneCountInStrings(helloWorld)

	hello := helloWorld[:12] 
	H := helloWorld[0]

	byteString =[]byte(helloWorld)
	helloWorld = string(byteString)
}