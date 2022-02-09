package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	// "strconv"
)


// http://hp-api.herokuapp.com/api/characters - view all characters
// http://hp-api.herokuapp.com/api/characters/students - view all characters who are Hogwarts students during the book series
// http://hp-api.herokuapp.com/api/characters/staff - view all characters who are Hogwarts staff during the book series
// http://hp-api.herokuapp.com/api/characters/house/:house - view all characters in a certain house, e.g. /slytherin

// Setup structure
type Details []struct {
	Name           string        `json:"name"`
	AlternateNames []interface{} `json:"alternate_names"`
	// AlternateNames []string 	 `json:"alternate_names"`
	Species        string        `json:"species"`
	Gender         string        `json:"gender"`
	House          string        `json:"house"`
	DateOfBirth    string        `json:"dateOfBirth"`
	YearOfBirth    int        	 `json:"yearOfBirth"`
	Wizard         bool          `json:"wizard"`
	Ancestry       string        `json:"ancestry"`
	EyeColour      string        `json:"eyeColour"`
	HairColour     string        `json:"hairColour"`
	Wand           struct {
		Wood   string `json:"wood"`
		Core   string `json:"core"`
		Length int    `json:"length"`
	} `json:"wand"`
	Patronus        string        `json:"patronus"`
	HogwartsStudent bool          `json:"hogwartsStudent"`
	HogwartsStaff   bool          `json:"hogwartsStaff"`
	Actor           string        `json:"actor"`
	AlternateActors []interface{} `json:"alternate_actors"`
	Alive           bool          `json:"alive"`
	Image           string        `json:"image"`
}

// func getdata(){
//     response, err := http.Get("http://hp-api.herokuapp.com/api/characters/")

//     if err != nil {
//         fmt.Print(err.Error())
//         os.Exit(1)
//     }

//     body, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//         log.Fatal(err)
//     }
//     // fmt.Println(string(body))

//     var responseObject Details
//     json.Unmarshal(body, &responseObject)

// }

// func writedata(){
//     for i, s := range responseObject {
//         fmt.Println("Position", i, "contains the string:", s)
//     }
// }


func main() {

    response, err := http.Get("http://hp-api.herokuapp.com/api/characters/")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    // fmt.Println(string(body))

    var responseObject Details
    json.Unmarshal(body, &responseObject)

	// for i, s := range responseObject {
    //     fmt.Println("Position", i, "contains the string:", s)
    // }

	for i := 0; i < len(responseObject); i++ {
        fmt.Println("Character Name: " + responseObject[i].Name)
        // fmt.Println("Alternate Name: " + responseObject[i].AlternateNames)
        fmt.Println("Species: " + responseObject[i].Species)
		fmt.Println("Gender: " + responseObject[i].Gender)
		fmt.Println("House: " + responseObject[i].House)
		fmt.Println("DateOfBirth: " + responseObject[i].DateOfBirth)
		// fmt.Println("YearOfBirth: " + responseObject[i].YearOfBirth)
		// fmt.Println("Wizard: " + responseObject[i].Wizard)
		fmt.Println("Ancestry: " + responseObject[i].Ancestry)
		fmt.Println("EyeColour: " + responseObject[i].EyeColour)
		fmt.Println("HairColour: " + responseObject[i].HairColour)
		fmt.Println("Wand Wood Type: " + responseObject[i].Wand.Wood)
		fmt.Println("Wand Core Type: " + responseObject[i].Wand.Core)
		fmt.Println("Wand Length: " + responseObject[i].Wand.Length)

	}

}
