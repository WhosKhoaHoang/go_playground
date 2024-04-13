package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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