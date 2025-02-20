package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"flag"
)


func main() {
	uri, currDay := cunstructUri()
	nextPrayer := flag.String("nextPrayer", "", "next prayer")
	flag.Parse()
	data, err := fetchAdhanData(&uri)
	if err != nil {
		fmt.Println(err)
		return
	}


	if *nextPrayer != "" {
		switch *nextPrayer {
		case "Fajr":
			println("Fajr: " + data.Data[currDay - 1].Timings.Fajr)
		case "Dhuhr":
			println("Dhuhr: " + data.Data[currDay - 1].Timings.Dhuhr)
		case "Asr":
			println("Asr: " + data.Data[currDay - 1].Timings.Asr)
		case "Maghrib":
			println("Maghrib: " + data.Data[currDay - 1].Timings.Maghrib)
		case "Isha":
			println("Isha: " + data.Data[currDay - 1].Timings.Isha)
		default:
			fmt.Println("Invalid prayer name")
		}
	} else {
		println("Fajr: " + data.Data[currDay - 1].Timings.Fajr)
		println("Dhuhr: " + data.Data[currDay - 1].Timings.Dhuhr)
		println("Asr: " + data.Data[currDay - 1].Timings.Asr)
		println("Maghrib: " + data.Data[currDay - 1].Timings.Maghrib)
		println("Isha: " + data.Data[currDay - 1].Timings.Isha)
	}

}


func fetchAdhanData(uri *string) (*Response, error) {	
	// fmt.Println("Fetching data from " + *uri)

	client := http.Client{}
	resp, err := client.Get(*uri)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	var data Response
	// print(resp)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &data, nil
}

func cunstructUri() (string, int) {
	currYear := time.Now().Year()
	currMonth := time.Now().Month()
	currDay := time.Now().Day()
	country := "Morocco"
	city := "Benguerir"

	uri := fmt.Sprintf("http://api.aladhan.com/v1/calendarByCity?city=%s&country=%s&method=21&month=%d&year=%d", city, country, currMonth, currYear)
	return uri, currDay
}