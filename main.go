package main

import (
	"unicode/utf8"
	"fmt"
)

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
}

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %v:%v", a.username, a.password)
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

type sender struct {
	user
	rateLimit int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.number == 0 || mToSend.sender.name =="" || mToSend.recipient.number == 0 || mToSend.recipient.name == "" {
	return false
	}
	return true
}


func splitEmail(email string) (string, string) {
	username, domain := "", ""
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}


func bootup() {
	defer fmt.Println("TEXTIO BOOTUP DONE")
	ok := connectToDB()
	if !ok {
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All systems ready!")
}

// don't touch below this line

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}


func printReports(intro, body, outro string) {
	printCostReport(func(s string) int {
		return 2*len(s)
	},intro) 
	printCostReport(func(s string) int {
		return 3*len(s)
	},body)
	printCostReport(func(s string) int {
		return 4*len(s)
	},outro)
}

// don't touch below this line


func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}


func reformat(message string, formatter func(string) string) string {
	return "TEXTIO: " + formatter(formatter(formatter(message)))
}


func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}



func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

// don't touch below this line

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}


func getMonthlyPrice(tier string) int {
	switch tier {
		case "basic":
			return 10000
		case "premium":
			return 15000
		case "enterprise":
			return 50000
		default:
			return 0
	}
}
func concat(s1 string, s2 string) string {
	return s1 + s2
}





func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

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
	var startup string = "Textio SMS service booting up..."
	var message1 string = "Sending test message"
	var confirmation string = "Message sent!"

	// don't touch below this line

	fmt.Println(startup)
	fmt.Println(message1)
	fmt.Println(confirmation)
	senderName := "Syl"
	recipient := "Kaladin"
	message := "The Words, Kaladin. You have to speak the Words!"

	fmt.Printf("%s to %s: %s\n", senderName, recipient, message)
	var penniesPerText float64 = 2

	// don't edit below this line
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	fname := "Dalinar"
	lname := "Kholin"
	age4 := 45
	messageRate := 0.5
	isSubscribed := false
	message4 := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %v %v, Age: %v, Rate: %v, Is Subscribed: %v, Message: %v", fname, lname, age4, messageRate, isSubscribed, message4)

	// Don't touch below this line

	fmt.Println(userLog)
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))

	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
	test(true, true)
	test(false, true)
	test(true, false)
	test(false, false)
}

