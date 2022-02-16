package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"bufio"
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


// Start functions here

func SingleChar() {

		var name string
		// var name string
		// var lastName string
		// fmt.Println("\nEnter Character First and Last Name (Capitalized): \n" + "Example: Harry Potter\n")
		fmt.Printf("\nEnter Character First and Last Name (Capitalized): \n" + "Example: Harry Potter\n")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			name = scanner.Text()
		}

		// response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/gryffindor")
		response, err := http.Get("http://hp-api.herokuapp.com/api/characters")
	
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

	
		for charNum, charData := range responseObject {
		if (name) == charData.Name{
			fmt.Println("\n--------------------------------------------")
			fmt.Println("Record Number:    ", charNum)
			fmt.Println("Character Name:   ", charData.Name)
			fmt.Println("House:            ", charData.House)
			fmt.Println("Species:          ", charData.Species)
			fmt.Println("Gender:           ", charData.Gender)
			fmt.Println("Date of Birth:    ", charData.DateOfBirth)
			fmt.Println("Year of Birth:    ", charData.YearOfBirth)
			fmt.Println("Wizard:           ", charData.Wizard)
			fmt.Println("Ancestry:         ", charData.Ancestry)
			fmt.Println("Eye Color:        ", charData.EyeColour)
			fmt.Println("Hair Color:       ", charData.HairColour)
			fmt.Println("Alternate Names:  ", charData.AlternateNames)
			fmt.Println("Patronus:         ", charData.Patronus)
			fmt.Println("Hogwarts Student: ", charData.HogwartsStudent)
			fmt.Println("Hogwarts Staff:   ", charData.HogwartsStaff)
			fmt.Println("Actor:            ", charData.Actor)
			fmt.Println("Alternate Actors: ", charData.AlternateActors)
			fmt.Println("Alive:            ", charData.Alive)
			fmt.Println("Image:            ", charData.Image)
			fmt.Println("Wand Details")
			fmt.Println(" Wand Wood Type:  ", charData.Wand.Wood)
			fmt.Println(" Wand Core:       ", charData.Wand.Core)
			fmt.Println(" Wand Length:     ", charData.Wand.Length)
			fmt.Println("\n")
			   }
		}
}

func House() {

	var house string
	// var name string
	// var lastName string
	fmt.Println("\nEnter the House Name you would like to see (Capitalized): \n" + "Gryffindor\n" + "Hufflepuff\n" + "Ravenclaw\n" + "Slytherin\n")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		house = scanner.Text()
	}

	// response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/gryffindor")
	response, err := http.Get("http://hp-api.herokuapp.com/api/characters")

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


	for charNum, charData := range responseObject {
	if (house) == charData.House{
		fmt.Println("\n--------------------------------------------")
			fmt.Println("Record Number:    ", charNum)
			fmt.Println("Character Name:   ", charData.Name)
			fmt.Println("House:            ", charData.House)
			fmt.Println("Species:          ", charData.Species)
			fmt.Println("Gender:           ", charData.Gender)
			fmt.Println("Date of Birth:    ", charData.DateOfBirth)
			fmt.Println("Year of Birth:    ", charData.YearOfBirth)
			fmt.Println("Wizard:           ", charData.Wizard)
			fmt.Println("Ancestry:         ", charData.Ancestry)
			fmt.Println("Eye Color:        ", charData.EyeColour)
			fmt.Println("Hair Color:       ", charData.HairColour)
			fmt.Println("Alternate Names:  ", charData.AlternateNames)
			fmt.Println("Patronus:		   ", charData.Patronus)
			fmt.Println("Hogwarts Student: ", charData.HogwartsStudent)
			fmt.Println("Hogwarts Staff:   ", charData.HogwartsStaff)
			fmt.Println("Actor:            ", charData.Actor)
			fmt.Println("Alternate Actors: ", charData.AlternateActors)
			fmt.Println("Alive:            ", charData.Alive)
			fmt.Println("Image:            ", charData.Image)
			fmt.Println("Wand Details")
			fmt.Println(" Wand Wood Type:  ", charData.Wand.Wood)
			fmt.Println(" Wand Core:       ", charData.Wand.Core)
			fmt.Println(" Wand Length:     ", charData.Wand.Length)
			fmt.Println("\n")
		   }
	}
}

func Staff() {

	// response, err := http.Get("http://hp-api.herokuapp.com/api/characters/house/gryffindor")
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


	for charNum, charData := range responseObject {
		fmt.Println("\n--------------------------------------------")
			fmt.Println("Record Number:    ", charNum)
			fmt.Println("Character Name:   ", charData.Name)
			fmt.Println("House:            ", charData.House)
			fmt.Println("Species:          ", charData.Species)
			fmt.Println("Gender:           ", charData.Gender)
			fmt.Println("Date of Birth:    ", charData.DateOfBirth)
			fmt.Println("Year of Birth:    ", charData.YearOfBirth)
			fmt.Println("Wizard:           ", charData.Wizard)
			fmt.Println("Ancestry:         ", charData.Ancestry)
			fmt.Println("Eye Color:        ", charData.EyeColour)
			fmt.Println("Hair Color:       ", charData.HairColour)
			fmt.Println("Alternate Names:  ", charData.AlternateNames)
			fmt.Println("Patronus:		   ", charData.Patronus)
			fmt.Println("Hogwarts Student: ", charData.HogwartsStudent)
			fmt.Println("Hogwarts Staff:   ", charData.HogwartsStaff)
			fmt.Println("Actor:            ", charData.Actor)
			fmt.Println("Alternate Actors: ", charData.AlternateActors)
			fmt.Println("Alive:            ", charData.Alive)
			fmt.Println("Image:            ", charData.Image)
			fmt.Println("Wand Details")
			fmt.Println(" Wand Wood Type:  ", charData.Wand.Wood)
			fmt.Println(" Wand Core:       ", charData.Wand.Core)
			fmt.Println(" Wand Length:     ", charData.Wand.Length)
			fmt.Println("\n")
	}
}

func main() {

	var input string
	fmt.Print("\nEnter 1 to get specific character details\n" + "Enter 2 for All Character details per house\n" + "Enter 3 for all Staff member details\n")
	fmt.Scanf("%s", &input)

	if input == "1" {
		SingleChar()		
	} else if input == "2" {
		House()		
	} else if input == "3" {
		Staff()		
	}
}