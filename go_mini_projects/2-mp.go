package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath = `/Users/anilyarramreddy/Downloads/chromedriver`
	port         = 9515
)

type cheapest_hotel struct {
	name  string
	price string
}

func main() {

	hotel_names_list := make([]string, 50)
	price_list := make([]string, 50)

	location := "Chennai, Tamil Nadu, India"
	from_date := "2021-10-07"
	to_date := "2021-10-08"
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	defer wd.Quit()
	if err != nil {
		panic(err)
	}
	if err := wd.Get("https://www.kayak.co.in/hotels/" + location + "-c13827/" + from_date + "/" + to_date + "/2adults?sort=rank_a"); err != nil {
		panic(err)
	}
	time.Sleep(6 * time.Second)

	dropdown, err := wd.FindElement(selenium.ByCSSSelector, "div[id='ResultsListSortDropdownLabel'")
	print(dropdown.Text())
	time.Sleep(10 * time.Second)
	dropdown.Click()
	time.Sleep(6 * time.Second)
	if err != nil {
		panic(err)
	}

	selectli, err := wd.FindElements(selenium.ByCSSSelector, "div[class='c5FFC-item'")
	//c5FFC-item
	print(selectli)
	for _, webElement := range selectli {
		data, _ := webElement.Text()
		print(data)
		if data == "Price (low to high)" {
			webElement.Click()
			break
		}
	}
	time.Sleep(5 * time.Second)
	hotel_names, err := wd.FindElements(selenium.ByCSSSelector, "div[class='FLpo-big-name'")
	if err != nil {
		panic(err)
	}

	for i, webElement := range hotel_names {
		data, _ := webElement.Text()
		hotel_names_list[i] = data
	}

	price, err := wd.FindElements(selenium.ByCSSSelector, "div[class='zV27-price'")

	for i, webElement := range price {
		data, _ := webElement.Text()
		price_list[i] = data
	}

	ch := cheapest_hotel{
		name:  hotel_names_list[0],
		price: price_list[0],
	}

	fmt.Println(ch.name, ch.price)
	defer service.Stop()

}
