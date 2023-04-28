package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              //Don't get reflected in the Json response
	Tags     []string `json:"tags,omitempty"` //omit is to not add the attribute if it is null/nil
}

func main() {
	fmt.Println("Welcome to JSON")
	//EncodeJson()
	decodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React JS ", 300, "test.in", "test123", []string{"web", "js"}},
		{"Java ", 500, "test.in", "test345", []string{"web", "java", "spring"}},
		{"Angular JS ", 1000, "test.in", "test723", []string{"web", "js", "angular", "spring"}},
		{"Test ", 100, "test.in", "test823", nil},
	}

	//Package data as Json
	finalJson, err := json.Marshal(courses)
	checkError(err)

	fmt.Printf("%s\n", string(finalJson))

	finalJsonIndent, err := json.MarshalIndent(courses, "", "\t")
	checkError(err)

	fmt.Printf("%s\n", string(finalJsonIndent))

	finalJsonIndentPrefix, err := json.MarshalIndent(courses, "TEST", "\t")
	checkError(err)

	fmt.Printf("%s\n", string(finalJsonIndentPrefix))

}

func decodeJson() {
	jsonFromWeb := []byte(`
	[
        {
                "coursename": "React JS ",
                "price": 300,
                "website": "test.in",
                "tags": [
                        "web",
                        "js"
                ]
        },
        {
                "coursename": "Java ",
                "price": 500,
                "website": "test.in",
                "tags": [
                        "web",
                        "java",
                        "spring"
                ]
        },
        {
                "coursename": "Angular JS ",
                "price": 1000,
                "website": "test.in",
                "tags": [
                        "web",
                        "js",
                        "angular",
                        "spring"
                ]
        },
        {
                "coursename": "Test ",
                "price": 100,
                "website": "test.in"
        }
	]
	`)
	//Validate the json, as the result was LIST use slice []course
	var myCourse []course
	valid := json.Valid(jsonFromWeb)

	if valid {
		fmt.Println("JSON was Valid")
		//Use reference & to be sure that data is getting added in correct object
		err := json.Unmarshal(jsonFromWeb, &myCourse)
		checkError(err)

		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

	//It is like Map<String, Object>
	var myData []map[string]interface{}
	err := json.Unmarshal(jsonFromWeb, &myData)
	checkError(err)
	fmt.Printf("%#v\n", myData)

	for key, value := range myData {
		fmt.Printf("Key is %v and Value is %v and Type of the Data is %T\n", key, value, value)
	}

	//Get first element from the result
	dataOne, err := json.Marshal(myData[0])
	checkError(err)
	fmt.Println("*******************")

	fmt.Println(string(dataOne))

	var cCourse course
	myErr := json.Unmarshal(dataOne, &cCourse)
	checkError(myErr)
	fmt.Printf("%#v\n", cCourse)

	fmt.Println("**************")

	var mCourse map[string]interface{}
	myMapErr := json.Unmarshal(dataOne, &mCourse)
	checkError(myMapErr)
	for key, value := range mCourse {
		fmt.Printf("Key is %v and Value is %v and Type of the Data is %T\n", key, value, value)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
