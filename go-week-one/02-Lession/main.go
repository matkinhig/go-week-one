package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	fmt.Println("______________________________________________________")

	converJSON()
	castingVALUE()
}

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	ClassName   string `json:"course_name"`
	AcademyName string `json:"acedemy_name"`
}

func converJSON() {
	jsonString := "{\"first_name\":\"Tin\",\"last_name\":\"Tran\",\"email_address\":\"tintran@gmail.com\",\"age\":80,\"course_name\":\"golang0110\",\"acedemy_name\":\"Nordic Coder\"}"

	bs := []byte(jsonString)
	var aStudent Student
	json.Unmarshal(bs, &aStudent)
	fmt.Printf("%+v", aStudent)

}

func castingVALUE() {
	aInt := 5
	var bFloat float64
	bFloat = float64(aInt)
	println(bFloat)
}
