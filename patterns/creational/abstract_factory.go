package creational

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	ORG_IXIGO      = "ixigo"
	ORG_MAKEMYTRIP = "makemytrip"
)

type iBookingAbstractFactory interface {
	bookFlight() iFlight
	bookTrain() iTrain
}

func getBookingFactory(orgName string) (baf iBookingAbstractFactory, err error) {
	switch orgName {
	case ORG_IXIGO:
		baf = &ixigo{}
	case ORG_MAKEMYTRIP:
		baf = &makemytrip{}
	default:
		err = fmt.Errorf("Booking org not defined!")
	}
	return
}

type iFlight interface {
	setFlight(pnr string)
	getFlight() (pnr string)
}
type iTrain interface {
	setTrain(pnr string)
	getTrain() (pnr string)
}

/*
Implement iFlight & iTrain interfaces by Flight & Train concrete struct types
*/
type Flight struct {
	pnr                string
	travelDate         time.Time
	sourceAirport      string
	destinationAirport string
	flightDetails      map[string]interface{}
}

type Train struct {
	pnr                string
	travelDate         time.Time
	sourceStation      string
	destinationStation string
	trainDetails       map[string]interface{}
}

func (f *Flight) setFlight(pnr string) {
	f.pnr = pnr
}
func (f *Flight) getFlight() (pnr string) {
	pnr = f.pnr
	return
}

func (t *Train) setTrain(pnr string) {
	t.pnr = pnr
}
func (t *Train) getTrain() (pnr string) {
	pnr = t.pnr
	return
}

/*
Concrete types for ixigo/makemytrip bookings - train & flights
*/
type ixigoFlight struct {
	Flight
}

type ixigoTrain struct {
	Train
}
type makemytripFlight struct {
	Flight
}

type makemytripTrain struct {
	Train
}

/*
Implement BookingAbstractFactory interface by ixigo & makemytrip concrete struct types
*/
type ixigo struct {
	//Flight *Flight
	//Train  *Train
}

type makemytrip struct {
	//Flight *Flight
	//Train  *Train
}

func (i *ixigo) bookFlight() iFlight {
	id, _ := uuid.NewUUID()
	return &ixigoFlight{Flight{pnr: "IXI-F" + id.String()}}

	//i.Flight = &Flight{}
	//return i.Flight
}

func (i *ixigo) bookTrain() iTrain {

	id, _ := uuid.NewRandom()
	return &ixigoTrain{Train{pnr: "IXI-T" + id.String()}}

	//i.Train = &Train{}
	//return i.Train
}

func (i *makemytrip) bookFlight() iFlight {
	id, _ := uuid.NewUUID()
	return &makemytripFlight{Flight{pnr: "MMT-F" + id.String()}}

	//i.Flight = &Flight{}
	//return i.Flight
}

func (i *makemytrip) bookTrain() iTrain {

	id, _ := uuid.NewRandom()
	return &makemytripTrain{Train{pnr: "MMT-T" + id.String()}}
	//i.Train = &Train{}
	//return i.Train
}

func ExecuteAbstractFactory() {
	ix, _ := getBookingFactory(ORG_IXIGO)
	//ix := &ixigo{}
	ixF := ix.bookFlight()
	PrintFlightDetails(ixF)
	ixT := ix.bookTrain()
	PrintTrainDetails(ixT)

	mmt, _ := getBookingFactory(ORG_MAKEMYTRIP)
	//mmt := &makemytrip{}
	mmtF := mmt.bookFlight()
	PrintFlightDetails(mmtF)
	mmtT := mmt.bookTrain()
	PrintTrainDetails(mmtT)
}

func PrintFlightDetails(f iFlight) {
	fmt.Println("Flight: ", f.getFlight())
}

func PrintTrainDetails(t iTrain) {
	fmt.Println("Train: ", t.getTrain())
}
