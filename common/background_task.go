package common

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

type BackgroundTask interface {
	InitEmailSchedulers()
	getFirstEmailQueue() Email
	AddEmailQueue(queueOfEmail Email)
	SendNextEmail(email Email) error
}
type backgroundTask struct {
	cron  *gocron.Scheduler
	queue []Email
}

type Email struct {
	ToEmail  []string
	Role     string
	RegToken string
}

func NewBackgroundTask() *backgroundTask {
	cron := gocron.NewScheduler(time.Now().Location())
	return &backgroundTask{cron, []Email{}}
}

func (bg *backgroundTask) InitEmailSchedulers() {
	bg.cron.Every(10).Second().Do(bg.CheckEmailQueue)
	bg.cron.StartAsync()
	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	// s.StartAsync()
	// starts the scheduler and blocks current execution path
	// s.StartBlocking()

}
func (bg *backgroundTask) CheckEmailQueue() {
	fmt.Printf("=========INSIDE CheckEmailQueue %s \n", time.Now())
	if len(bg.queue) != 0 {
		email := bg.getFirstEmailQueue()
		if bg.SendNextEmail(email) == nil {
			bg.RemoveFirstEmailQueue()
		}
	}
}
func (bg *backgroundTask) getFirstEmailQueue() Email {
	// fmt.Println("=========INSIDE getFirstEmailQueue")
	var queueElement Email
	if len(bg.queue) != 0 {
		queueElement = bg.queue[0]
		// bg.queue = bg.queue[1:]
	}
	return queueElement
}
func (bg *backgroundTask) RemoveFirstEmailQueue() Email {
	// fmt.Println("=========INSIDE getFirstEmailQueue")
	var queueElement Email
	if len(bg.queue) != 0 {
		queueElement = bg.queue[0]
		bg.queue = bg.queue[1:]
	}
	return queueElement
}
func (bg *backgroundTask) AddEmailQueue(queueOfEmail Email) {
	// fmt.Println("=========INSIDE AddEmailQueue")
	bg.queue = append(bg.queue, queueOfEmail)
}

func (bg *backgroundTask) SendNextEmail(email Email) error {
	// fmt.Println("=========INSIDE SendNextEmail")
	//Send Confirmation Email
	// regToken := randstr.Hex(16) // generate 128-bit hex string
	msg := []byte("To: " + email.ToEmail[0] + "\r\n" +
		"Subject: Registration Confirmation Email from News App!\r\n" +
		"\r\n" +
		"This is the email body.\r\n" +
		"http://localhost:8080/api/register/confirmation?email=" + email.ToEmail[0] + "&token=" + email.RegToken + "&role=" + email.Role)

	return SendEmail(email.ToEmail, msg)
}

func (bg *backgroundTask) JustFunc() { fmt.Println("alhamdulillah") }
