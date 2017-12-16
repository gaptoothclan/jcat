package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
	"reflect"
	"flag"
)

func main() {
	var keyFlag = flag.Bool("k", false, "Get json keys")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please provide a json file");
		return
	}
	filename := args[0]

	if _, err := os.Stat(filename); err != nil {
		fmt.Println("This is not a file");
		return
	}

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file");
		return
	}

	var c interface{}
	json.Unmarshal(contents, &c)

	m := c.(map[string]interface{})

	byteSlice := getJsonByteSlice(args, m)

	if *keyFlag == true {
		var c1 interface{}
		json.Unmarshal(byteSlice, &c1)
		m1 := c1.(map[string]interface{})
		for key, _ := range m1 {
			fmt.Printf("%s\n", key)
		}
	} else {
		fmt.Println(string(byteSlice))
	}
	
}

func getJsonByteSlice(args []string, jsonData map[string]interface{}) []byte {
	if len(args) > 1 {
		jsonFields := strings.Split(args[1], ".")
		subJson := getSubJson(jsonData, jsonFields)
		byteSlice, _ := json.MarshalIndent(subJson, "", "    ")
		return byteSlice
	} else {
		byteSlice, _ := json.MarshalIndent(jsonData, "", "    ")
		return byteSlice
	}
}

func getSubJson(json map[string]interface{}, key []string) map[string]interface{} {

	if len(key) == 1 && json[key[0]] != nil {

		if reflect.TypeOf(json[key[0]]) != reflect.TypeOf(json) {
			var newJson = make(map[string]interface{})
			newJson[key[0]] = json[key[0]]
			return newJson
		}

		return json[key[0]].(map[string]interface{})
	} else if len(key) == 1 {
		return json
	}

	if json[key[0]] == nil {
		return json
	}

	return getSubJson(json[key[0]].(map[string]interface{}), key[1:])
}
