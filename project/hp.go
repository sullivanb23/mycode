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


func main() {
	var input string
	// var character string
	fmt.Print("\nEnter one of the values below to see character details: \n" + "gryffindor\n" + "hufflepuff\n" + "ravenclaw\n" + "slytherin\n" + "staff\n")
	fmt.Scanf("%s", &input)
	// fmt.Scanf("%s", &character)

	if input == "gryffindor" || input == "Gryffindor" {
    response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/gryffindor")

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

	// for i := 0; i < len(responseObject); i++ {
	// 	fmt.Println("\n--------------------------------------------")
    //     fmt.Println("Character Name: " + responseObject[i].Name)
	// 	fmt.Println(" House: " + responseObject[i].House)
    //     // fmt.Println("Alternate Name: " + responseObject[i].AlternateNames)
    //     fmt.Println(" Species: " + responseObject[i].Species)
	// 	fmt.Println(" Gender: " + responseObject[i].Gender)
	// 	fmt.Println(" DateOfBirth: " + responseObject[i].DateOfBirth)
	// 	// fmt.Println("YearOfBirth: " + responseObject[i].YearOfBirth)
	// 	// fmt.Println("Wizard: " + responseObject[i].Wizard)
	// 	fmt.Println(" Ancestry: " + responseObject[i].Ancestry)
	// 	fmt.Println(" EyeColour: " + responseObject[i].EyeColour)
	// 	fmt.Println(" HairColour: " + responseObject[i].HairColour)
	// 	fmt.Println(" Wand: ")
	// 	fmt.Println("  Wand Wood Type: " + responseObject[i].Wand.Wood)
		// fmt.Println("  Wand Core Type: " + responseObject[i].Wand.Core)
	// 	// fmt.Println("Wand Length: " + responseObject[i].Wand.Length)

	// }

	for charNum, charData := range responseObject {
		fmt.Println("\n--------------------------------------------")
		fmt.Println("Record Number:   ", charNum)
		fmt.Println("Character Name:  ", charData.Name)
		fmt.Println("House:           ", charData.House)
		fmt.Println("Species:         ", charData.Species)
		fmt.Println("Gender:          ", charData.Gender)
		fmt.Println("Date of Birth:   ", charData.DateOfBirth)
		fmt.Println("Year of Birth:   ", charData.YearOfBirth)
		fmt.Println("Wizard:          ", charData.Wizard)
		fmt.Println("Ancestry:        ", charData.Ancestry)
		fmt.Println("Eye Color:       ", charData.EyeColour)
		fmt.Println("Hair Color:      ", charData.HairColour)
		fmt.Println("Wand Details")
		fmt.Println(" Wand Wood Type: ", charData.Wand.Wood)
		fmt.Println(" Wand Core:      ", charData.Wand.Core)
		fmt.Println(" Wand Length:    ", charData.Wand.Length)
		fmt.Println("Alternate Names: ", charData.AlternateNames)
		fmt.Println("Alternate Actors:", charData.AlternateActors)
		fmt.Println("\n")
		   }
	}

	if input == "hufflepuff" || input == "Hufflepuff" {
		response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/hufflepuff")
	
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
			fmt.Println("\n--------------------------------------------")
			fmt.Println("Character Name: " + responseObject[i].Name)
			fmt.Println(" House: " + responseObject[i].House)
			// fmt.Println("Alternate Name: " + responseObject[i].AlternateNames)
			fmt.Println(" Species: " + responseObject[i].Species)
			fmt.Println(" Gender: " + responseObject[i].Gender)
			fmt.Println(" DateOfBirth: " + responseObject[i].DateOfBirth)
			// fmt.Println("YearOfBirth: " + responseObject[i].YearOfBirth)
			// fmt.Println("Wizard: " + responseObject[i].Wizard)
			fmt.Println(" Ancestry: " + responseObject[i].Ancestry)
			fmt.Println(" EyeColour: " + responseObject[i].EyeColour)
			fmt.Println(" HairColour: " + responseObject[i].HairColour)
			fmt.Println(" Wand: ")
			fmt.Println("  Wand Wood Type: " + responseObject[i].Wand.Wood)
			fmt.Println("  Wand Core Type: " + responseObject[i].Wand.Core)
			// fmt.Println("Wand Length: " + responseObject[i].Wand.Length)
		}
		}

	if input == "ravenclaw" || input == "Ravenclaw" {
		response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/ravenclaw")
	
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
			fmt.Println("\n--------------------------------------------")
			fmt.Println("Character Name: " + responseObject[i].Name)
			fmt.Println(" House: " + responseObject[i].House)
			// fmt.Println("Alternate Name: " + responseObject[i].AlternateNames)
			fmt.Println(" Species: " + responseObject[i].Species)
			fmt.Println(" Gender: " + responseObject[i].Gender)
			fmt.Println(" DateOfBirth: " + responseObject[i].DateOfBirth)
			// fmt.Println("YearOfBirth: " + responseObject[i].YearOfBirth)
			// fmt.Println("Wizard: " + responseObject[i].Wizard)
			fmt.Println(" Ancestry: " + responseObject[i].Ancestry)
			fmt.Println(" EyeColour: " + responseObject[i].EyeColour)
			fmt.Println(" HairColour: " + responseObject[i].HairColour)
			fmt.Println(" Wand: ")
			fmt.Println("  Wand Wood Type: " + responseObject[i].Wand.Wood)
			fmt.Println("  Wand Core Type: " + responseObject[i].Wand.Core)
			// fmt.Println("Wand Length: " + responseObject[i].Wand.Length)
	
		}
		}

	if input == "slytherin" || input == "Slytherin" {
		response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/slytherin")
		
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
			fmt.Println("\n--------------------------------------------")
			fmt.Println("Character Name: " + responseObject[i].Name)
			fmt.Println(" House: " + responseObject[i].House)
			// fmt.Println("Alternate Name: " + responseObject[i].AlternateNames)
			fmt.Println(" Species: " + responseObject[i].Species)
			fmt.Println(" Gender: " + responseObject[i].Gender)
			fmt.Println(" DateOfBirth: " + responseObject[i].DateOfBirth)
			// fmt.Println("YearOfBirth: " + responseObject[i].YearOfBirth)
			// fmt.Println("Wizard: " + responseObject[i].Wizard)
			fmt.Println(" Ancestry: " + responseObject[i].Ancestry)
			fmt.Println(" EyeColour: " + responseObject[i].EyeColour)
			fmt.Println(" HairColour: " + responseObject[i].HairColour)
			fmt.Println(" Wand: ")
			fmt.Println("  Wand Wood Type: " + responseObject[i].Wand.Wood)
			fmt.Println("  Wand Core Type: " + responseObject[i].Wand.Core)
			// fmt.Println("Wand Length: " + responseObject[i].Wand.Length)
	
		}
		}

	if input == "staff" || input == "Staff" {
		response, err := http.Get("http://hp-api.herokuapp.com/api/characters/staff")
		
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
			fmt.Println("\n--------------------------------------------")
			fmt.Println("Character Name: " + responseObject[i].Name)
			fmt.Println(" House: " + responseObject[i].House)
			// fmt.Println("Alternate Name: " + responseObject[i].AlternateNames)
			fmt.Println(" Species: " + responseObject[i].Species)
			fmt.Println(" Gender: " + responseObject[i].Gender)
			fmt.Println(" DateOfBirth: " + responseObject[i].DateOfBirth)
			// fmt.Println("YearOfBirth: " + responseObject[i].YearOfBirth)
			// fmt.Println("Wizard: " + responseObject[i].Wizard)
			fmt.Println(" Ancestry: " + responseObject[i].Ancestry)
			fmt.Println(" EyeColour: " + responseObject[i].EyeColour)
			fmt.Println(" HairColour: " + responseObject[i].HairColour)
			fmt.Println(" Wand: ")
			fmt.Println("  Wand Wood Type: " + responseObject[i].Wand.Wood)
			fmt.Println("  Wand Core Type: " + responseObject[i].Wand.Core)
			// fmt.Println("Wand Length: " + responseObject[i].Wand.Length)
	
		}
		}
}
