package main

import (
	"fmt"
	"io/ioutil"
)

/**
* @author Arturo Negreiros
* @description Reading csv files to get the correct prediction
 */

func readingCSVFiles() {
	readFile, err := ioutil.ReadFile("./.csv/.clouds_parameters/.Overcast_clouds/.2017/.2017.csv")

	if err != nil {
		fmt.Println("An error ocurred!!")
	} else {
		fmt.Println(string(readFile))
	}
}

func main() {

	readingCSVFiles()
	fmt.Println("My first neuronal network")
}
