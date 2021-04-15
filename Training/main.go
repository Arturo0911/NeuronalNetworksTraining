package main

/**
* @author Arturo Negreiros
* @description Reading csv files to get the correct prediction
 */

import (
	"fmt"
	"io/ioutil"
	"time"
)

var PathFiles = "./.csv/.clouds_parameters"

/**
* @Struct: Each csv file contains the following features

*
**/
type OvercastClouds struct {
	TimeStart          time.Time
	TimeEnd            time.Time
	CloudDescription   string
	RelativityHumidity float32
	Clouds             int
	Precipitaion       float32
	Temperature        float32
	Icon               string
	Code               string
}

func MathProcess() {

}

func readingCSVFiles(cloudType string, year string) {

	newPathFile := "./.csv/.clouds_parameters" + "/." + cloudType + "/." + year + "/." + year + ".csv"
	readFile, err := ioutil.ReadFile(newPathFile)

	if err != nil {
		fmt.Println("An error ocurred!!")
	} else {
		fmt.Println(string(readFile))
	}
}

func main() {

	readingCSVFiles("Overcast_clouds", "2017")
	fmt.Println("My first neuronal network")
}
