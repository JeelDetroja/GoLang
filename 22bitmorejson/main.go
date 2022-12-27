package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"cousename"`
	Price     int
	Plateform string
	Password  string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to jason")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	Course := []course{
		{"golang", 344, "jehedb.com", "abc123", []string{"web", "js"}},
		{"C#", 122, "jehedb.com", "abc1234", []string{"web", "js"}},
		{"java", 233, "jehedb.com", "abc12345", nil},
	}

	finalJson, err := json.MarshalIndent(Course, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFormWeb := []byte(`
	{
		"cousename": "golang",
		"Price": 344,
		"Plateform": "jehedb.com",
		"tags": ["web","js"]
	}
	`)

	var Course course

	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("JSO was valid")
		json.Unmarshal(jsonDataFormWeb, &Course)
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFormWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key : %v and value : %v and type : %T\n", k, v, v)
	}
}
