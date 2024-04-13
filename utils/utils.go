package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func TestReadFile(fileName string) Test123 {
	// infile, err := os.Open(fileName)
	infile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("DANG! SOMETHING WENT WRONG WITH ReadFile...")
	}

	test123 := Test123{}
	err = json.Unmarshal(infile, &test123)
	if err != nil {
		fmt.Println("SOMETHING WENT WRONG WITH UNMARSHAL!")
	}

	return test123
}

func TestReadTarget(fileName string) []Target123 {

	infile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("SOMETHING WENT WRONG WITH READFILE!")
	}

	var target123 []Target123
	json.Unmarshal(infile, &target123)

	fmt.Println("INSIDE OF TestReadTarget:")
	fmt.Println(infile)
	fmt.Println(len(infile))

	var result []Target123
	return result
}

func TestReadUser(fileName string) {
	byteValue, _ := ioutil.ReadFile(fileName)

	var users Users
	json.Unmarshal(byteValue, &users)

	fmt.Println("In TestReadUser...")
	fmt.Println(byteValue)
	fmt.Println(len(byteValue))
	fmt.Println(users.Users)
	fmt.Println(len(users.Users))

	// for i := 0 ; i < len(users.Users); i++ {
  //   fmt.Println("User Type: " + users.Users[i].Type)
  //   fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
  //   fmt.Println("User Name: " + users.Users[i].Name)
  //   fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }

	// TODO: play around with how you can use goroutines to process users.Users
	for _, v := range users.Users {
		fmt.Println("User Type: " + v.Type)
    fmt.Println("User Age: " + strconv.Itoa(v.Age))
    fmt.Println("User Name: " + v.Name)
    fmt.Println("Facebook Url: " + v.Social.Facebook)
	}
}