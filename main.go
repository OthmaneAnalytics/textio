package main

import (
/*
	"unicode/utf8"
	"time"
	"errors"
	"strings"
	"math/rand"
	"sync"
*/
	"fmt"
)

type emailStatus = int

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)

func (a *analytics) handleEmailBounce(em email) error {
	err1 := em.recipient.updateStatus(em.status) 
	if err1 != nil {
		return fmt.Errorf("error updating user status: %w", err1)
	}
	err2 := a.track(em.status)
	if err2 != nil {
		return fmt.Errorf("error tracking user bounce: %w", err2)
	}
	return nil
}
/*
type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if balance < newItem.GetCost(){
		var zerroT []T
		return zerroT, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, balance - newItem.GetCost(), nil

}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}




func getLast[T any](s []T) T {
	if len(s) == 0 {
		var myZero T
		return myZero
	}
	return s[len(s) - 1]
}


type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}
*/
/*
type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}
*/

/*
func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
			case <- snapshotTicker:
				takeSnapshot(logChan)
			case <- saveAfter:
				saveSnapshot(logChan)
				return
			default:
				waitForData(logChan)
				time.Sleep(500 * time.Millisecond)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func logMessages(chEmails, chSms chan string) {
	for {
		select {
			case e, ok := <- chEmails:
				if !ok {
					return
				} else {
					logEmail(e)
				}
			case s, ok := <- chSms:
				if !ok {
					return
				} else {
					logSms(s)
				}
		}
	}
}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}
func test(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}


func concurrentFib(n int) []int {
	fib := []int{}
	ch := make(chan int)
	go fibonacci(n, ch)
	for item := range ch {
		fib = append(fib, item)
	}
	return fib
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}


func countReports(numSentCh chan int) int {
	n := 0
	for {
		if v , ok:= <- numSentCh; !ok {
			break
		} else {
			n += v
		}

	}
	return n
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}


func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for i:= 0 ; i< len(emails); i++{
		ch <- emails[i]
	}
	return ch
}

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i:= 0; i< numDBs; i++ {
		<-dbChan
	}
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}


type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(data *Analytics, msg Message){
	data.MessagesTotal++
	if msg.Success {
		data.MessagesSucceeded++	
	} else {
		data.MessagesFailed++
	}
}

func getNameCounts(names []string) map[rune]map[string]int {
	data := map[rune]map[string]int{}
	for i := 0 ; i < len(names); i++ {
		runes := []rune(names[i])
		if _, ok1 := data[runes[0]]; ok1 {
			if _, ok2 := data[runes[0]][names[i]]; ok2 {
				data[runes[0]][names[i]]++
			} else {
			data[runes[0]][names[i]] = 1
			}
		} else {
			nmap := map[string]int{names[i]: 1}
			data[runes[0]] = nmap
		}
	}
	return data
}

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for i := 0 ; i < len(messagedUsers) ; i++ {
		if _, ok := validUsers[messagedUsers[i]]; ok {
			validUsers[messagedUsers[i]]++
		}
	}
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	elem, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	} else if !elem.scheduledForDeletion {
		return false, nil
	} else {
		delete(users, name)
		return true, nil
	}
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalide sizes")
	}
	userMap := make(map[string]user)
	for i:= 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name : names[i], 
			phoneNumber : phoneNumbers[i],
		}
	}
	return userMap, nil
}



func indexOfFirstBadWord(msg []string, badWords []string) int {
	index := -1
	for i, element := range msg {
		for _, bad := range badWords {
			if element == bad && index == -1 {
				index = i
				break
			}
		}
	}
	return index
}

func createMatrix(rows, cols int) [][]int {
	mat := make([][]int,rows)
	for i := 0; i < rows; i++{
		mat[i] = make([]int,cols)
		for j := 0 ; j < cols; j++{
			mat[i][j] = i*j
		}
	}
	return mat
}


type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	if len(costs) == 0 {
		return nil
	} else {
		cpd := make([]float64,1)
		for i := 0; i< len(costs); i++{
			if costs[i].day < len(cpd) {
				cpd[costs[i].day] += costs[i].value
			} else {
				rest := make( []float64 , costs[i].day + 1 - len( cpd ) )
				rest[ len(rest)-1 ] = costs[i].value
				cpd = append( cpd , rest... )
			}
		}
		return cpd
	}
}

func sum(nums ...int) int {
	sum := 0
	for i := 0; i< len(nums); i++{
		sum += nums[i]
	}
	return sum
}

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64,len(messages))
	for i := 0; i < len(messages); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}  
	return costs
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan{
		case "free":
		return messages[:2], nil
		case "pro":
			return messages[:], nil
		default:
			return nil, errors.New("unsupported plan")
	}
}


func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	return [3]string{primary, secondary, tertiary}, [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

}

func printPrimes(max int) {
	for n := 2; n <=max; n++{ 
		if n == 2 {
			fmt.Println(n)
		} else if n % 2 == 0 {
			continue
		} else {
			prime := true
			for i:= 2; i*i <= n; i++ {
				if n % i == 0 {
					prime = false
					break
				}
			}
			if prime {
				fmt.Println(n)
			}
		}
	}
}


func fizzbuzz() {
	for i := 1; i<=100 ; i++{
		if i % 15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 5 == 0 {
			fmt.Println("buss")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		} 
	}
}


func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0{
		maxMessagesToSend--
	}
	return maxMessagesToSend
}


func maxMessages(thresh int) int {	
	s := 0
	for i := 0 ; ; i++{
		s += 100 + i
		if s > thresh{
			return i	
		}
	}
}


func bulkSend(numMessages int) float64 {
	s := 0.0
	for i := 0 ; i< numMessages; i++{
		s += 1.0 + float64(i)*0.01
	}
	return s
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

type divideError struct {
	dividend float64
}

func (e divideError) Error() string{
	return fmt.Sprintf("can not divide %v by zero", e.dividend)
}

func divide2(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent",cost ,recipient)
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	i1, err1 := sendSMS(msgToCustomer)
	if err1 != nil {
		return 0, err1
	}
	i2, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0, err2
	}
	return i1 + i2, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}


// don't touch below this line

type expense interface {
	cost() float64
}


type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}


func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}





type formatter interface {
	format() string
}


type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getSalary() int{
	return c.hourlyPay * c.hoursPerYear
}

func (c contractor) getName() string {
	return c.name
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}


func sendMessage(msg message) (string, int) {
	content := msg.getMessage()
	l := len(content)
	return content, 3*l
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}


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


type sender struct {
	user
	rateLimit int
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
	fizzbuzz()


	rand.Seed(0)

	test(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)

}
*/
