# Real Mini Project 
from selenium import webdriver
from bs4 import BeautifulSoup
import time
import tkinter as tk
from tkinter import ttk
import pandas as pd
# import requests
# from_city = input("Flying from: ")
# to_city =  input("Flying to: ")
# day = input("What date (dd/mm/yyyy): ")

def popup(msg):
    popup = tk.Tk()
    popup.wm_title("Flights Data ") 
    NORM_FONT = ("Verdana", 20)
    label = ttk.Label(popup, text=msg, font=NORM_FONT)
    label.pack(side="top", fill="x", pady=10)
    B1 = ttk.Button(popup, text="Ok", command = popup.destroy)
    B1.pack()
    popup.mainloop()


def search_flight():
      
  try:
    # from_loc = e1.get()
    # to_loc = e2.get()
    # date = e3.get() 
    # print(from_loc)
    from_loc = "chennai" #input("Enter Source : ")
    to_loc = "bangalore" #input("Enter Destination : ")
    date = "30/09/2021" #input("")
    #url = "https://www.expedia.ie/Flights-Search?trip=oneway&leg1=from:"+from_loc+",to:"+to_loc+",departure:"+date+"TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass:economy&mode=search&origref=www.expedia.ie"

    url = "https://www.expedia.co.in/Flights-Search?trip=oneway&leg1=from:" + from_loc + ",to:" + to_loc + ",departure:" + date + "TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass:economy&mode=search"
    print(f"URL: {url}")
    print("The cheapest flights: \n")
    # r = requests.get(url)
    driver = webdriver.Safari()
    driver.get(url)
    time.sleep(5)
    depa_time = driver.find_elements_by_xpath("//span[contains(@data-test-id,'departure-time')]")
    departure_time_list = [a.text for a in depa_time]

    arrival_departure = driver.find_elements_by_xpath("//div[contains(@data-test-id,'arrival-departure')]")
    arrival_departure_list = [a.text for a in arrival_departure]

    journey_duration = driver.find_elements_by_xpath("//div[contains(@data-test-id,'journey-duration')]")
    journey_duration_list = [a.text for a in journey_duration]

    price_column = driver.find_elements_by_xpath("//div[contains(@data-test-id,'price-column')]")
    price_column_list = [a.text for a in price_column]
    print(departure_time_list)
    print(arrival_departure_list)
    print(journey_duration_list)
    print(price_column_list)
    
    # fetching data using BeautifulSoup
    #soup = BeautifulSoup(driver.page_source, 'html.parser')
    # Getting all the data from the website using html elements and tags.

    # departure_time = soup.find_all('span', attrs={'data-test-id': 'departure-time'}) 
    # arrival_departure = soup.find_all('div', attrs={'data-test-id': 'arrival-departure'}) 
    # journey_duration = soup.find_all('div', attrs={'data-test-id': 'journey-duration'})
    # price_column = soup.find_all('div', attrs={'data-test-id': 'price-column'})

    # Cleaning up the data, such as getting only text and removing whitespace. This all gets stored in list using list comprehension.
    # departure_time_list = [a.getText().strip() for a in departure_time]
    # arrival_departure_list = [a.getText().strip() for a in arrival_departure]
    # journey_duration_list = [b.getText().strip() for b in journey_duration]
    price_column_list = []
    for pr in price_column:
        p = str(pr.text)
        price = p[p.find('Rs')+1 : p.find('Rs',p.find('Rs', p.find('Rs')+1))]
        price_column_list.append('R'+price)
    
    if len(departure_time_list) == 0 :
        print("Flights are not available at this moment...... ")
        return
    driver.quit()
    #Adding all data to Dictionary
    flights = {"arrival - departure": arrival_departure_list,
              "departure_time" : departure_time_list,    
               "journey_duration": journey_duration_list,
                "price": price_column_list}    
    flights_data = pd.DataFrame(flights)
    
    print("Current Cheapest Flight is = ")

    # exporting the data into excel sheet
    flights_data.to_excel("output.xlsx", index=None)
  
    print(flights_data)
    
    msg="Flights data downloaded successful !! \n"
    popup(msg)
   
  except Exception as e: 
    print(e)
    popup("Error While Fetching Data....") 
    
    pass

root = tk.Tk()
text = tk.Text(root,bg='light blue')
text.grid(row=4,column=0,columnspan=2)  
root.title('Mini Project Find Chepest Flights')
tk.Label(root,text="Enter Source : ").grid(row=0)

e1 = tk.Entry(root)
e1.grid(row=0,column=1)

tk.Label(root,text="Enter Destination : ").grid(row=1)
e2 = tk.Entry(root)
e2.grid(row=1, column=1)
tk.Label(root,text="Enter Date (dd/mm/YYYY) : ").grid(row=2)
e3 = tk.Entry(root)
e3.grid(row=2,column=1)

tk.Button(root,
          text='Search', command=search_flight,anchor=tk.CENTER).grid(row=3,column=1,
                                                       sticky=tk.W,
                                                       pady=4) 
print(e1.get())                                                       
root.mainloop()     
# search_flight(from_city, to_city, day)
