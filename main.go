package main

import "fmt"

func main() {
	var smsSendingLimit int
	smsSendingLimit = 0
	var costPerSMS float64
	costPerSMS = 0
	var hasPermission bool
	hasPermission = false
	var username string
	username = ""
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	messageStart := "Happy birthday! you are now"
	age := 21
	messageEnd := "years old"
	fmt.Println(messageStart, age, messageEnd)

}

