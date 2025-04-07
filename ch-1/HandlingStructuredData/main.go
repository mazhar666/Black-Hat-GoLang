package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {
	// --------------------- JSON MARSHAL/UNMARSHAL EXAMPLE ---------------------
	// Create a Foo struct instance with some data
	f := Foo{"Joe Junior", "Hello Shabado"}

	// Marshal the Foo struct into JSON
	b, _ := json.Marshal(f) // Converts struct to []byte JSON

	// Print the JSON as a string
	fmt.Println(string(b)) // Output: {"Bar":"Joe Junior","Baz":"Hello Shabado"}

	// Unmarshal the JSON back into the struct (just to demonstrate)
	json.Unmarshal(b, &f)

	// --------------------- XML MARSHAL EXAMPLE WITH TAGS ---------------------
	// Create a Boo struct instance with same values
	x := Boo{"Joe Junior", "Hello Shabado"}

	// Marshal the Boo struct into formatted XML
	output, err := xml.MarshalIndent(x, "", "  ") // Nicely indented
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the resulting XML string
	fmt.Println(string(output))
}

// Foo is a basic struct with two fields (used for JSON marshaling)
type Foo struct {
	Bar string
	Baz string
}

// Boo is a struct where XML tags define how it should be marshaled to XML
type Boo struct {
	Bar string `xml:"id,attr"`      // Bar will become an XML attribute called "id"
	Baz string `xml:"parent>child"` // Baz will be nested inside <parent><child>...</child></parent>
}
