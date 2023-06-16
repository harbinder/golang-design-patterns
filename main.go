package main

import (
	"examples/channels"
	types "examples/data-types"
	"examples/interfaces"
	"examples/patterns/behavioural"
	"examples/patterns/creational"
	"examples/patterns/structural"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/channel", chanExamples)
	app.Get("/struct", structExamples)
	app.Get("/interface", interfaceExamples)
	app.Get("/pattern/structural/bridge", patternStructuralBridgeExamples)
	app.Get("/pattern/structural/decorator", patternStructuralDecorator)
	app.Get("/pattern/creational/factory", patternCreationalFactory)
	app.Get("/pattern/creational/singleton", patternCreationalSingleton)
	app.Get("/pattern/creational/abstract-factory", patternCreationalAbstractFactory)
	app.Get("/pattern/creational/object-pool", patternCreationalObjectPool)
	app.Get("/pattern/behavioural/template-method", patternBehaviouralTemplateMethod)
	app.Get("/pattern/behavioural/iterator", patternBehaviouralIterator)
	app.Get("/array", arrayExamples)
	app.Get("/slice", sliceExamples)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func chanExamples(c *fiber.Ctx) error {
	gc := channels.GoChannel{}
	//gc.RoutineOne()
	gc.CheckOsSignal()
	return c.SendString("Channel Examples")
}

func interfaceExamples(c *fiber.Ctx) error {

	/*
		create an object of SportsPerson type
		BUT assigned to an interface type which is implemented by SportsPerson
		We can assign it in 2 ways
		1. Value Type
			ip = interfaces.SportsPerson{}
		2. Pointer Type
			ip = &interfaces.SportsPerson{}
		If assigned as Value DataType(ip=sp),
			ip will NOT BE able to access Pointer Receiver Methods of SportsPerson struct
		If assigned as Pointer DataType(ip=&sp),
			ip will BE able to access BOTH Pointer & Value Receiver Methods of SportsPerson struct

	*/
	var ip interfaces.IPerson
	// Memory Map -> ip : [nil,nil]
	sp := interfaces.SportsPerson{Name: "Sachin Tendulkar", Gender: "Male"}
	// Memory Map -> sp : [Name, Gender]
	//ip = sp
	/*
		Memory Map : ip =
		[
			Pointer -> sp,
			Pointer  -> SportsPerson.PrintPerson()
		]
	*/

	ip = &sp
	fmt.Printf("Underlying Type: %T\n", ip)
	fmt.Printf("Underlying Value: %v\n", ip)
	/*
		Memory Map : ip =
		[
			Pointer -> Pointer -> sp,
			Pointer  -> SportsPerson.PrintPerson()
		]
	*/

	// 2 Ways to Create an object of Cricketer type
	/*
		1. cr := new(interfaces.Cricketer) - cr is Pointer type struct variable
		   Memory Map : cr -> [Name: "", Gender: "", Team: ""]
		2. cr := interfaces.Cricketer{} - cr is Value type struct variable
	*/

	//cr := new(interfaces.Cricketer)
	cr := interfaces.Cricketer{}
	// Memory Map : cr = [Name: "", Gender: "", Team: ""]
	cr.Team = "Mumbai Indians"
	// Memory Map : cr = [Name: "", Gender: "", Team: "Mumbai Indians"]

	/*
		cr.PrintPerson()

		Cricketer struct can access PROMOTED Method -> PrintPerson()  of SportsPerson struct,
		without Cricketer struct implementing methods of IPerson interface,
		which would have to be defined, in case Cricketer struct implicitly implements Iperson interface by defining all methods of IPerson


	*/
	fmt.Println("cr.PrintCricketer() call using object of type -> Cricketer struct ")
	cr.PrintCricketer(ip)

	var crI interfaces.ICricketer
	/*
		Memory Map : crI =
		[
			Pointer -> nil,
			Pointer -> nil
		]
	*/
	crI = &cr
	fmt.Printf("Underlying Type: %T\n", crI)
	fmt.Printf("Underlying Value: %v\n", crI)
	// crI = cr
	//  - Compile Error :
	//  cr (interfaces.Cricketer) does not implement interfaces.ICricketer
	//  Because interfaces.Cricketer struct has implemented interface method via Pointer Receiver AND
	// a Value Type struct object CANNOT access Pointer Receiver Methods HOWEVER vice versa is accessible i.e
	// a Pointer Type struct object CAN  access Value Receiver Methods
	/*
		Memory Map : crI =
		[
			Pointer -> cr,
			Pointer -> cr.PrintCricketer()
		]
	*/
	fmt.Println("crI.PrintCricketer() call using variable of type -> ICricketer interface ")
	crI.PrintCricketer(ip)

	cr1 := new(interfaces.Cricketer)
	// cr1 is Pointer type struct (interfaces.Cricketer) object, AND is able to access BOTH
	// Value & Pointer Receiver Methods of struct (interfaces.Cricketer)

	cr1.IPerson = ip
	cr1.PrintCricketer(ip)
	cr1.SetCricketer("Harman Preet Kaur", "Female", "Chennai Super Kings")
	cr1.PrintCricketer(ip)

	cr1.PrintCricketer1()

	return c.JSON(map[string]interface{}{})
}

func structExamples(c *fiber.Ctx) error {
	//p := &interfaces.Person{Name: "Harry", Age: 41, Gender: "Male"}
	/*
		Instantiace struct with new function with default attribte values
	*/
	p := new(interfaces.Person)
	p.PrintDetails()

	e := &interfaces.Employee{P: interfaces.Person{Name: "Harry", Age: 41, Gender: "Male"}}

	/*
		Cannot access Person attributes as it is defined as private with attribute [p Person] instead of [P Person]
	*/
	ps := &interfaces.PersonalSecretory{} //{p: interfaces.Person{Name: "Harry", Age: 41, Gender: "Male"}}

	m := &interfaces.Manager{Person: interfaces.Person{Name: "Harry", Age: 41, Gender: "Male"}}
	return c.JSON(map[string]interface{}{"Person": p, "Employee": e, "Manager": m, "PersonalSecretory": ps})
}

func patternStructuralBridgeExamples(c *fiber.Ctx) error {
	structural.Execute()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func patternStructuralDecorator(c *fiber.Ctx) error {
	structural.ExecuteDecorator()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func patternCreationalFactory(c *fiber.Ctx) error {
	creational.ExecuteFactory()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}
func patternCreationalSingleton(c *fiber.Ctx) error {
	creational.ExecuteSingleton()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func patternCreationalAbstractFactory(c *fiber.Ctx) error {
	creational.ExecuteAbstractFactory()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func patternCreationalObjectPool(c *fiber.Ctx) error {
	creational.ExecuteObjectPool()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func patternBehaviouralTemplateMethod(c *fiber.Ctx) error {
	behavioural.ExecuteTemplateMethod()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func patternBehaviouralIterator(c *fiber.Ctx) error {
	behavioural.ExecuteIterator()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})
}

func arrayExamples(c *fiber.Ctx) error {
	types.ArrayDeclaration()
	types.ArrayIterate()
	types.ArrayMultiDimention()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})

}

func sliceExamples(c *fiber.Ctx) error {
	types.SliceCreation()
	types.SliceAppend()
	types.SliceIterate()
	//types.SliceMultiDimention()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})

}
