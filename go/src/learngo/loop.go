package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func convertToBin(n int) string{
	result := ""
	for ;n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string)  {
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13))
	printFile("src/learngo/abc.txt")
	//s := `dasda
	//"DASDD"	dasdas
	//123
	//`
	//printFileContents(strings.NewReader(s))
}