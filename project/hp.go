package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)


// http://hp-api.herokuapp.com/api/characters - view all characters
// http://hp-api.herokuapp.com/api/characters/students - view all characters who are Hogwarts students during the book series
// http://hp-api.herokuapp.com/api/characters/staff - view all characters who are Hogwarts staff during the book series
// http://hp-api.herokuapp.com/api/characters/house/:house - view all characters in a certain house, e.g. /slytherin


type Details []struct {
        Name           string        `json:"name, omitempty"`
        AlternateNames []interface{} `json:"alternate_names, omitempty"`
        Species        string        `json:"species, omitempty"`
        Gender         string        `json:"gender, omitempty"`
        House          string        `json:"house, omitempty"`
        DateOfBirth    string        `json:"dateOfBirth, omitempty"`
        YearOfBirth    int           `json:"yearOfBirth, omitempty"`
        Wizard         bool          `json:"wizard, omitempty"`
        Ancestry       string        `json:"ancestry, omitempty"`
        EyeColour      string        `json:"eyeColour, omitempty"`
        HairColour     string        `json:"hairColour, omitempty"`
        Wand           struct {
                Wood   string `json:"wood, omitempty"`
                Core   string `json:"core, omitempty"`
                Length int    `json:"length, omitempty"`
        } `json:"wand"`
        Patronus        string        `json:"patronus, omitempty"`
        HogwartsStudent bool          `json:"hogwartsStudent, omitempty"`
        HogwartsStaff   bool          `json:"hogwartsStaff, omitempty"`
        Actor           string        `json:"actor, omitempty"`
        AlternateActors []interface{} `json:"alternate_actors, omitempty"`
        Alive           bool          `json:"alive, omitempty"`
        Image           string        `json:"image, omitempty"`
}


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
-- INSERT --                                                                                                                                                                                                                                                1,1           Top
[ttyd] 0:vim*                                                                                                                                                                                                                          "student@bchd: ~/mycod" 16:29 09-Feb-22
