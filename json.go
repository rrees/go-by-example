package main

import "fmt"
import "encoding/json"

func main() {
	booleanPayload, _ := json.Marshal(true)
	fmt.Println(string(booleanPayload))

	jsonString := []byte(`{"num" : 6.13, "strings" : ["a", "b"]}`)

	var data map[string]interface{}

	if error := json.Unmarshal(jsonString, &data); error != nil {
		panic(error)
	}

	num := data["num"].(float64)
	fmt.Println(num)

	strings := data["strings"].([]interface{})
	firstString := strings[0].(string)

	fmt.Println(firstString)

}