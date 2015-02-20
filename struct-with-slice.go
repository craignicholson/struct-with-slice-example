// Copyright 2015 Craig Nicholson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
  "strconv"
)

// Customer with one or more address
type Customer struct {
	FirstName string
	LastName  string
	Address   []Location
}

// Location
type Location struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Country  string
	Lat      string
	Long     string
}

func main() {

  //Example 1 - We know the # of Slices
	customer := LengthKnown()
	PrintJSON(customer)

  //Example 2 - Dynamic Slice
  customer = DynamicSlice()
	PrintJSON(customer)

}

// Example showing how we can add multiple locations
// to one customer when you know the number of items
// to be added
func LengthKnown() Customer {

	//A slice must have the length supplied
	locations := make([]Location, 2)
	locations[0].Address1 = "101 Go Lang Blvd"
	locations[0].Address2 = "Unit 11"
	locations[0].City = "AnyCity"
	locations[0].State = "AnyState"
	locations[0].Country = "AnyCountry"

	locations[1].Address1 = "202 Go Lang Blvd"
	locations[1].Address2 = "Unit 22"
	locations[1].City = "2AnyCity"
	locations[1].State = "2AnyState"
	locations[1].Country = "2AnyCountry"

	// Add a customer with locations
	customer := Customer{}
	customer.FirstName = "Craig"
	customer.LastName = "Nicholson"
	customer.Address = locations

	return customer
}

// Example showing how we can add multiple locations
// to one customer when you do not know the # of
// locations to be added
func DynamicSlice() Customer {

	// A slice must have the length supplied
	locations := make([]Location, 0)

	for i := 0; i < 3; i++ {
		data := Location{}
		data.Address1 = strconv.Itoa(i)
		data.Address2 = strconv.Itoa(i)
		data.City = strconv.Itoa(i)
		data.State = strconv.Itoa(i)
		data.Country = strconv.Itoa(i)

    locations = append(locations,data)
	}

	// Add a customer with locations
	customer := Customer{}
	customer.FirstName = "Craig"
	customer.LastName = "Nicholson"
	customer.Address = locations

	return customer
}

// Print JSON out to the consoles
func PrintJSON(customer Customer){
  json, err := json.Marshal(customer)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Print the Results
  fmt.Println(string(json))
  fmt.Println("\n")
}
