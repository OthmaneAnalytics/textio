package main

import (
	"fmt"
	"unicode/utf8"
)


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
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
	fmt.Println("The compiled textio server is starting")
	accountAgeFloat := 2.6
	fmt.Println("Your account has existed for", int64(accountAgeFloat), "years")
	var username2 string = "presidentSkroob"
	var password string = "12345"
	// don't edit below this line
	fmt.Println("Authorization: Basic", username2+":"+password)
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
	const name = "Saul Goodman"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %s, your open rate is %0.1f percent\n",name , openRate)

	// don't edit below this line

	fmt.Print(msg)

	const name3 = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name3))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name3))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name3)
}
