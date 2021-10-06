package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	qa := map[string]string{
		"hi":               "Hello",
		"Hi":               "Hello",
		"Hello":            "Hi",
		"hello":            "Hi,  How can i help you",
		"how are you?":     "I'm fine and you?",
		"How are you?":     "I'm fine and you?",
		"i'm fine too":     "nice to hear that",
		"I need help":      "Please explain your complaint",
		"Thank you":        "Have a nice day",
		"Are you a robot?": "Yes I’m a robot but I’m a smart one!",
		"Are you human?":   " No i am robot.",
		"Do you love me?":  "I love you.",
	}

	// Taking input from user

	v, found := qa["How are you?"]
	fmt.Println(v, found)

	fmt.Println(qa["How are you?"])

	for {
		fmt.Print("You : ")
		// var user_msg string
		// fmt.Scanf("%s", &user_msg)
		// fmt.Println(user_msg)
		in := bufio.NewReader(os.Stdin)

		user_msg, err := in.ReadString('\n')
		user_msg = strings.TrimSpace(user_msg)
		//fmt.Println(user_msg)
		if err != nil {
			return
		}
		if user_msg == "exit" {
			break
		}
		if v, found := qa[user_msg]; found {

			fmt.Println("Bot:", v)
		} else {
			fmt.Println("i didn't get you! ")
		}

	}

}
