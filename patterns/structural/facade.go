package structural

import (
	"fmt"
	"time"
)

/*
https://golangbyexample.com/facade-design-pattern-in-golang/
This design pattern is meant to hide the complexities of the underlying system and provide a simple interface to the client.
It provides a unified interface to underlying many interfaces in the system so that from the client perspective it is easier to use.
Basically it provides a higher level abstraction over a complicated system.

The term Facade itself means
the principal front of a building, that faces on to a street or open space
*/

type IService interface {
	//Create()
	//Read()
	//Update()
	Applicable()
}
type Service struct {
	IService      IService
	Name          string
	Desc          string
	Priority      int8
	VisibilityMap map[string]interface{}
}

// implement IService for ServiceAbstract struct
func (sa *Service) Applicable() {
	fmt.Println("Service: Applicable()")
}

func ExecuteFacade() {
	svc := &Service{}
	svc.Name = "Premium Customer Support"

	sf := &ServiceFacade{
		visibility: &Visibility{lob: "Flights", airlines: []string{"Indigo", "GoFirst"}},
	}
	fmt.Println(sf)
	//svc.VisibilityMap = *sf.visibility
}

/*
Facade of Service
*/
type ServiceFacade struct {
	visibility *Visibility
	refundable *Refundable
	timeframe  *Timeframe
}
type Visibility struct {
	lob      string
	airlines []string
}
type Refundable struct {
	customerCancel     bool
	airlineCancel      bool
	refundPercentage   int8
	refundDurationDays int8
}
type Timeframe struct {
	StartTime            time.Time
	EndTime              time.Time
	DaysBeforeTravelDate int8
	DaysAfterTravelDate  int8
}
