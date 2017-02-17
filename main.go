package main

import "fmt"

const i = iota

func main() {

	// result, err := http.Get("http://www.pgatour.com/data/trigger/trigger.xml")
	// if err != nil {
	// fmt.Println(err)
	// }
	// page, _ := ioutil.ReadAll(result.Body)
	// result.Body.Close()
	// fmt.Printf("%s\n", page)
	var variable int = 10
	fmt.Print("Type in var: ")
	fmt.Scan(&variable)

	fmt.Println(variable)
}
