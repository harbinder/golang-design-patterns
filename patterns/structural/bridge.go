package structural

import "fmt"

/*
https://golangbyexample.com/bridge-design-pattern-in-go/
Allows the separation of abstraction from its implementation
In this example, check how any type of notification (Email/Sms) BRIDGES with any type of Vendor (X,Y) on the runtime
*/

/*
We have 2 types of Vendors
1. Xvendor
2. Yvendor
*/
type vendor interface {
	send(map[string]interface{})
}
type Xvendor struct {
	mobile int8
}

func (xv *Xvendor) send(data map[string]interface{}) {
	fmt.Println("Sending notification via Xvendor")
	fmt.Println("Data : ", data)
}

type Yvendor struct {
	emailID string
}

func (yv *Yvendor) send(data map[string]interface{}) {
	fmt.Println("Sending notification via Yvendor")
	fmt.Println("Data : ", data)
}

/*
We have 2 types of notificatioins
1. Email
2. Sms
*/
type Notification interface {
	notify()
	setVendor(vendor)
}

type Email struct {
	vendor  vendor
	emailID string
}

func (e *Email) notify() {
	fmt.Println("\nTrigger Email Notification : " + e.emailID)
	data := map[string]interface{}{"mobile": e.emailID}
	e.vendor.send(data)
}

func (e *Email) setVendor(v vendor) {
	fmt.Printf("\nSet Email Notification Vendor : %T", v)
	e.vendor = v
}

type Sms struct {
	vendor vendor
	mobile string
}

func (s *Sms) notify() {
	fmt.Println("\nTrigger Sms Notification : " + s.mobile)
	data := map[string]interface{}{"mobile": s.mobile}
	s.vendor.send(data)
}

func (s *Sms) setVendor(v vendor) {
	fmt.Printf("\nSet Sms Notification Vendor : %T", v)
	s.vendor = v
}

func Execute() {
	// Initialise notification vendors
	xv := &Xvendor{}
	yv := &Yvendor{}

	// Send same SMS via both Vendors
	sms1 := &Sms{mobile: "9910825975"}
	// send SMS via vendor X
	sms1.setVendor(xv)
	sms1.notify()
	// send SMS via vendor Y
	sms1.setVendor(yv)
	sms1.notify()

	// Send same Email via both Vendors
	email1 := &Email{emailID: "harry@gmail.com"}
	// send Email via vendor X
	email1.setVendor(xv)
	email1.notify()
	// send Email via vendor Y
	email1.setVendor(yv)
	email1.notify()
}
