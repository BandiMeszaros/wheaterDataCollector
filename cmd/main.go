package main

import (
	"log"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")

func main() {
	w, err := owm.NewCurrent("C", "EN", apiKey) // (internal - OpenWeatherMap reference for kelvin) with English output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Balatonföldvár, HU")
	//jsonData, _ := json.Marshal(w)
	//jsonString := string(jsonData)
	//fmt.Println(jsonString)
	//fmt.Println("Ran")
}
