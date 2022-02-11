/* Alta3 Research | RZFeeser
   API - RESTful API lookup, and decoding JSON on 200 response  */
   package main

   import (
		   "encoding/json"
		   "fmt"
		   "log"
		   "net/http"
		//    "time"
   )
   
   
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
		   // define our URL (API) as a string
		   url := "http://hp-api.herokuapp.com/api/characters"
   
		   // Build the request
		   req, err := http.NewRequest("GET", url, nil)
		   if err != nil {
				   log.Fatal("NewRequest: ", err)
				   return
		   }
   
		   // For control over HTTP client headers,
		   // redirect policy, and other settings,
		   // create a Client
		   // A Client is an HTTP client - you can load it up with
		   // parameters if you wanted (timeouts, tls, retries, etc.)
		   // You only need ONE of these (supported by goroutines)
		   client := &http.Client{}
   
		   // Send the request via a client
		   // Do sends an HTTP request and
		   // returns an HTTP response
		   resp, err := client.Do(req)
		   if err != nil {
				   log.Fatal("Do: ", err)
				   return
		   }
   
		   // Callers should close resp.Body
		   // when done reading from it
		   // Defer the closing of the body
		   defer resp.Body.Close()
   
		   // Fill the record with the data from the JSON
		   var record Details
   
		   // Use json.Decode for reading streams of JSON data
		   if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
				   log.Println(err)
		   }
   
		   // the _ is an int
		   // launchData is the data
		   for charNum, charData := range record {
		//    fmt.Println("Capsule Record     =\n", launchData)
		   fmt.Println("Record Number      =", charNum)
		   fmt.Println("Character Name     =", charData.Name)
		   fmt.Println("Wand Wood   	   =", charData.Wand.Wood)
		   fmt.Println("Wand Core   	   =", charData.Wand.Core)
		   fmt.Println("Wand Length   	   =", charData.Wand.Length)
		   fmt.Println("End of Record\n\n")
		   }
   }
   