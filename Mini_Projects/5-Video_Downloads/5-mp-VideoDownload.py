import os
from pytube import YouTube
import time
import tkinter as tk
from tkinter import ttk

NORM_FONT= ("Verdana", 20)

def popup(msg):
    popup = tk.Tk()
    popup.wm_title("Download Status")
    label = ttk.Label(popup, text=msg, font=NORM_FONT)
    label.pack(side="top", fill="x", pady=10)
    B1 = ttk.Button(popup, text="Ok", command = popup.destroy)
    B1.pack()
    popup.mainloop()

def entry_fields():
  try:
    start_time=time.time()
    download_location = "youtube_downloads/"
    YouTube(e1.get()).streams.first().download(download_location)
    end_time=time.time()
    str1="download successful !! \n"
    str2 = "Total time taken: {} seconds".format(round(end_time-start_time,3))
    msg=str1+str2
    popup(msg)
  except Exception as e :
      print("Error While Downloading the Video....")  


master=tk.Tk()
master.geometry("750x350")
master.wm_title("Download Video From Youtube")
tk.Label(master,text="Enter Youtube Video URL: ").grid(row=0)
e1 = tk.Entry(master)
e1.grid(row=0,column=1)

tk.Button(master,
          text='Download', command=entry_fields,anchor=tk.CENTER).grid(row=1,
                                                       sticky=tk.W,
                                                       pady=4)          
tk.mainloop()
