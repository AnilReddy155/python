# import tkinter as tk
# from tkinter import ttk
import pandas as pd

questionAnswers = {"hi":"Hello",
        "Hi":"Hello",
        "Hello":"Hi",
        "hello":"Hi,  How can i help you",
        "how are you?":"I'm fine and you?",
        "How are you?":"I'm fine and you?",
        "i'm fine too": "nice to hear that",
        "I need help": "Please explain your complaint",
        "Thank you":"Have a nice day",
        "Are you a robot?": "Yes I’m a robot but I’m a smart one!",
        "Are you human?": " No i am robot.",
        "Do you love me?":"I love you."
    }
d = {}
def load_data() :
    try:
        d = {}
        data = pd.read_excel("output.xlsx")
        print(type(data))
        df = pd.DataFrame(data, index=None)

        print(df)
        q = df.get("Question")
        a = df.get("Answer")
        
        size = len(q)
        
        for i in range(size):
            d[q[i]] = a[i]
        
        print(d)
    except Exception as e:
        print(e)
        d = questionAnswers      


def send():
     print("'''")
    # msg = e.get()
    # send = "You: "+ msg
    # text.insert(tk.END,"\n" + send)
    # if questionAnswers.get(msg):
    #     text.insert(tk.END, "\n" + "Bot: "+questionAnswers.get(msg))
    # else:
    #     text.insert(tk.END, "\n" + "Bot: Sorry I didnt get it.")

load_data()
# root = tk.Tk()
# text = tk.Text(root,bg='light blue')
# text.grid(row=0,column=0,columnspan=2)
# e = tk.Entry(root,width=60)
# send = tk.Button(root,text='Send',bg='blue',width=20,command=send).grid(row=1,column=1)
# e.grid(row=1,column=0)
# # root = Tk()
# root.title('Mini Project Sample Chat Bot')
# root.mainloop()