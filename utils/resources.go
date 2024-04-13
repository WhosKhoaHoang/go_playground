package utils

type Test123 struct {
	IsKhoa bool `json:"isKhoa"`
	Name  string `json:"name"`
	Age  int `json:"age"`
}

type Target123 struct {
	IsHuman bool `json:"isHuman"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type Dude struct {
	X string
	Y string
}

// Users struct which contains
// an array of users
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

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}