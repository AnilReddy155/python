package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
	"github.com/tebeka/selenium"
)

const (
	seleniumPath = `/Users/anilyarramreddy/Downloads/chromedriver`
	port         = 9515
)

type CheapestFlight struct {
	depa_time         string
	arrival_departure string
	journey_duration  string
	cost_info         string
}

func main() {

	from_loc := "chennai"
	to_loc := "bangalore"
	date := "10/10/2021"

	url := "https://www.expedia.co.in/Flights-Search?trip=oneway&leg1=from:" + from_loc + ",to:" + to_loc + ",departure:" + date + "TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass:economy&mode=search"

	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	check(err)
	defer service.Stop()
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//Call browser urlPrefix: test reference: defaulturlprefix =” http://127.0.0.1 :4444/wd/hub”
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	check(err)
	//Delay exiting chrome
	defer wd.Quit()
	println(url)
	err = wd.Get(url)
	check(err)
	time.Sleep(4 * time.Second)
	//Find Baidu input box id
	cost_info_list := make([]string, 10)
	depa_time_list := make([]string, 10)
	arrival_departure_list := make([]string, 10)
	journey_duration_list := make([]string, 10)
	costInfo, err := wd.FindElements(selenium.ByCSSSelector, "span[class='uitk-price-a11y is-visually-hidden'")
	println(costInfo)

	for i, webElement := range costInfo {
		data, _ := webElement.Text()
		cost_info_list[i] = data
	}
	depa_time, err := wd.FindElements(selenium.ByCSSSelector, "span[data-test-id='departure-time'")
	print(depa_time)

	for i, webElement := range depa_time {
		data, _ := webElement.Text()
		depa_time_list[i] = data
	}
	arrival_departure, err := wd.FindElements(selenium.ByCSSSelector, "div[data-test-id='arrival-departure'")

	for i, webElement := range arrival_departure {
		data, _ := webElement.Text()
		arrival_departure_list[i] = data
	}
	journey_duration, err := wd.FindElements(selenium.ByCSSSelector, "div[data-test-id='journey-duration'")

	for i, webElement := range journey_duration {
		data, _ := webElement.Text()
		journey_duration_list[i] = data
	}

	print("Here is the Cheapest flight ")

	cheapest_flight := CheapestFlight{
		depa_time:         depa_time_list[0],
		arrival_departure: arrival_departure_list[0],
		journey_duration:  journey_duration_list[0],
		cost_info:         cost_info_list[0],
	}
	fmt.Println(cheapest_flight)
	export_to_excel(cheapest_flight)
}

func export_to_excel(cf CheapestFlight) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("flight")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Arrival-Departure"
	cell = row.AddCell()
	cell.Value = "Journey-Duration"
	cell = row.AddCell()
	cell.Value = "Departure-Tine"
	cell = row.AddCell()
	cell.Value = "Price"

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = cf.arrival_departure
	cell = row.AddCell()
	cell.Value = cf.journey_duration
	cell = row.AddCell()
	cell.Value = cf.depa_time
	cell = row.AddCell()
	cell.Value = cf.cost_info
	err = file.Save("cheapest_flight.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
