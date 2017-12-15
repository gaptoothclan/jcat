package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
	"reflect"
)

func main() {
	args := os.Args[1:]
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

	if len(args) == 2 {
		jsonFields := strings.Split(args[1], ".")
		subJson := getSubJson(m, jsonFields)
		byteSlice, _ := json.MarshalIndent(subJson, "", "    ")
		fmt.Println(string(byteSlice))
	} else {
		byteSlice, _ := json.MarshalIndent(m, "", "    ")
		fmt.Println(string(byteSlice))
	}

}

func getSubJson(json map[string]interface{}, key []string) map[string]interface{} {

	if len(key) == 1 && json[key[0]] != nil {

		if reflect.TypeOf(json[key[0]]) != reflect.TypeOf(json) {
			return json
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