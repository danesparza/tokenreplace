package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

func main() {
	//	Our command line options:
	tokenToReplace := flag.String("token", "{build}", "Token to replace")
	stringReplacement := flag.String("replacement", "Build number 1", "Replace the token with this string")
	filePath := flag.String("file", "filename.html", "File to use for token replacement")
	flag.Parse()

	//	Read a file into memory:
	buf, err := ioutil.ReadFile(*filePath)

	if err != nil {
		panic(err)
	}

	//	Convert from bytes to string:
	s := string(buf)

	//	Replace the token:
	s = strings.Replace(s, *tokenToReplace, *stringReplacement, -1)

	//	Print the result for now:
	//	fmt.Println(s)

	//	Convert to byte array and write the file
	ba := []byte(s)
	err = ioutil.WriteFile(*filePath, ba, 0755)
	if err != nil {
		panic(err)
	}
}
