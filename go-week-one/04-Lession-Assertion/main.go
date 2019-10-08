package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("start golang")
	jsonString := "{\"first_name\":\"Tin\",\"last_name\":\"Tran\",\"email_address\":\"tintran@gmail.com\",\"age\":80,\"course_name\":\"golang0110\",\"acedemy_name\":\"Nordic Coder\"}"

	var aMap map[string]interface{}
	bs := []byte(jsonString)
	json.Unmarshal(bs, &aMap)

	fmt.Println(aMap)

	var age interface{}
	age = aMap["age"]
	fmt.Println(age)

	var ageInt int
	if ageInt, er := age.(int); er == false {
		fmt.Println("can not assertion ", ageInt)
	}
	fmt.Println(ageInt)

	//chi parse qua dc kieu 64 bit tu chuoi json mang dang so

	var ageFloat float64
	ageFloat = age.(float64)
	fmt.Println(ageFloat)
}
