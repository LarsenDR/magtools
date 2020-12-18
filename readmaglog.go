package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

//Magdata a json data structure
type Magdata struct {
	ts    string
	lTemp float64
	rTemp float64
	x     float64
	y     float64
	z     float64
	rx    float64
	ry    float64
	rz    float64
}

// ParseJSON JSON data files
func (md *Magdata) ParseJSON(data []byte) {
	var result map[string]interface{}
	//fmt.Printf("%#v\n", string(data))
	json.Unmarshal([]byte(data), &result)
	timesp := result["ts"].(map[string]interface{})

	fmt.Printf("%#v/n", string(data))
	//fmt.Printf("%#v/n", timesp)

	for key, value := range timesp {
		fmt.Println(key, value.(string))
	}

}

// ParseCSV csv data files
func (md *Magdata) ParseCSV(data []byte) {

}

func main() {
	var mgdt Magdata

	iFlag := flag.String("i", "name", "input file")
	tFlag := flag.String("t", "JSON", "parse as JSON (default) or CSV")

	flag.Parse()

	fmt.Printf("-i %v\n", *iFlag)
	fmt.Printf("-t %v\n\n", *tFlag)

	data, err := ioutil.ReadFile(*iFlag)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	mgdt.ParseJSON(data)

	fmt.Printf("%#v\n", mgdt)
}
