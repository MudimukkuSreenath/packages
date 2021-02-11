package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Details interface{}
}

func main() {
	str := `{"name":"sreenath",
	"age": 30,
	"details": {"salary":100000}
	}`
	var data Person
	fmt.Println("===type of str===", reflect.TypeOf(str))
	fmt.Println("==str+==", str)
	fmt.Println("==[]byte(str)==", []byte(str))
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)

	}
	fmt.Printf("==data==%#v \n\n", data)
	fmt.Println("==name==", data.Name)
	details, ok := data.Details.(map[string]interface{})
	if ok {
		fmt.Println("==salary==", details["salary"])
	}
}
