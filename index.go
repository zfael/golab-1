package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	createFolders()
	createFiles()
	startOperation()
}

func createFolders() {
	os.Mkdir("multiplication", 0777)
	os.Mkdir("subtraction", 0777)
	os.Mkdir("addition", 0777)
	os.Mkdir("division", 0777)
}

func createFiles() {

	//addition folder
	f1, err := os.Create("addition/file.txt")
	check(err)

	defer f1.Close()

	d2 := []byte(`"11","12","13","14"`)
	_, err = f1.Write(d2)
	check(err)
	f1.Sync()

	//division folder
	f2, err := os.Create("division/file.txt")
	check(err)

	defer f2.Close()

	d2 = []byte(`80,70,60`)
	_, err = f2.Write(d2)
	check(err)

	f2.Sync()

	//multiplication folder
	f3, err := os.Create("multiplication/file.txt")
	check(err)

	defer f3.Close()

	d2 = []byte(`[10,20,30,40]`)
	_, err = f3.Write(d2)
	check(err)

	f3.Sync()

	//multiplication folder
	f4, err := os.Create("subtraction/file.txt")
	check(err)

	defer f4.Close()

	d2 = []byte(`1.1,0.5,0.001`)
	_, err = f4.Write(d2)
	check(err)

	f4.Sync()
}

func startOperation() {
	multResult := multiplicationOperation()
	subResult := subtractionOperation()
	addResult := additionOperation()
	divResult := divisionOperation()

	fmt.Println("Multiplication", multResult)
	fmt.Println("Subtraction", subResult)
	fmt.Println("Addition", addResult)
	fmt.Println("Division", divResult)
}

func multiplicationOperation() int {

	result := 1
	datb, err := ioutil.ReadFile("multiplication/file.txt")
	check(err)

	r := strings.NewReplacer("[", "", "]", "")
	dats := r.Replace(string(datb))

	arr := strings.Split(dats, ",")

	for i := 0; i < len(arr); i++ {
		value, err := strconv.Atoi(arr[i])
		check(err)
		result = result * value
	}

	return result
}

func subtractionOperation() float64 {

	var result float64
	datb, err := ioutil.ReadFile("subtraction/file.txt")
	check(err)

	dats := string(datb)
	arr := strings.Split(dats, ",")

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result, err = strconv.ParseFloat(arr[i], 64)
			check(err)
		} else {
			value, err := strconv.ParseFloat(arr[i], 64)
			check(err)
			result = result - value
		}
	}

	return result
}

func additionOperation() int {

	var result int
	datb, err := ioutil.ReadFile("addition/file.txt")
	check(err)

	r := strings.NewReplacer(`"`, "")
	dats := r.Replace(string(datb))

	arr := strings.Split(dats, ",")

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result, err = strconv.Atoi(arr[i])
			check(err)
		} else {
			value, err := strconv.Atoi(arr[i])
			check(err)
			result = result + value
		}
	}

	return result
}

func divisionOperation() float64 {

	var result float64
	datb, err := ioutil.ReadFile("division/file.txt")
	check(err)

	dats := string(datb)
	arr := strings.Split(dats, ",")

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result, err = strconv.ParseFloat(arr[i], 64)
			check(err)
		} else {
			value, err := strconv.ParseFloat(arr[i], 64)
			check(err)
			result = result / value
		}
	}

	return result
}
