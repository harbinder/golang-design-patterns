package creational

import (
	"fmt"
)

/*
Refer: https://golangbyexample.com/golang-factory-design-pattern/

Factory pattern is used to create concrete objects
All the complexity of creating an object is hidden and a factory is used to create different types of objects

This pattern provides a way to hide the creation logic of the instances being created.
The client only interacts with a factory struct and tells the kind of instances that needs to be created.
The factory class interacts with the corresponding concrete structs and returns the correct instance back.
*/

type NotificationType string

const (
	SMS_NOTIFICATION      NotificationType = "SMS"
	WHATSAPP_NOTIFICATION NotificationType = "WHATSAPP"
)

type iNotification interface {
	createNotification() (content string)
	sendNotification(content string)
}

type Notification struct {
	mobile string
	name   NotificationType
}

func (n *Notification) createNotification() (content string) {
	fmt.Println("Create Notification: ", n.name)
	content = "Notification Content"
	return
}
func (n *Notification) sendNotification(content string) {
	fmt.Println("Send Notification: ", n.name)
	return
}

// embedd Notification struct in SMS struct
type NotificationSMS struct {
	Notification
}

func NewInstanceSMS() (s *NotificationSMS) {
	return &NotificationSMS{
		Notification{name: SMS_NOTIFICATION},
	}
}

// embedd Notification struct in WhatsApp struct
type NotificationWhatsApp struct {
	Notification
}

func NewInstanceWhatsapp() (s *NotificationWhatsApp) {
	return &NotificationWhatsApp{
		Notification{name: WHATSAPP_NOTIFICATION},
	}
}

func NotificationFactory(nt NotificationType) (n iNotification) {
	switch nt {
	case SMS_NOTIFICATION:
		n = NewInstanceSMS()
	case WHATSAPP_NOTIFICATION:
		n = NewInstanceWhatsapp()
	}
	return
}

func ExecuteFactory() {
	sms := NotificationFactory(SMS_NOTIFICATION)
	wa := NotificationFactory(WHATSAPP_NOTIFICATION)

	content := sms.createNotification()
	sms.sendNotification(content)

	content = wa.createNotification()
	wa.sendNotification(content)
}
