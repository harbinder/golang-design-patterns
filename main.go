package main

import (
	"examples/channels"
	"examples/data-structure/linklist"
	"examples/data-structure/sort"
	"examples/data-structure/stack"
	"examples/data-structure/tree"
	types "examples/data-types"
	"examples/data-types/channel"
	"examples/data-types/interfaces"
	"examples/data-types/strings"
	structs "examples/data-types/struct"
	"examples/misc"
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
	port := "3000"

	// Routes
	api := app.Group("golang")
	api.Get("/channel", chanExamples)
	api.Get("/channel/basics", chanBasicExamples)
	api.Get("/struct/basics", structExamples)
	api.Get("/struct/embedding", embeddingExample)
	api.Get("/interface", interfaceExamples)
	api.Get("/string", stringsExamples)
	api.Get("/array", arrayExamples)
	api.Get("/slice", sliceExamples)
	api.Get("/map", mapExamples)

	pattern := api.Group("pattern")
	pattern.Get("/structural/bridge", patternStructuralBridgeExamples)
	pattern.Get("/structural/decorator", patternStructuralDecorator)
	pattern.Get("/creational/factory", patternCreationalFactory)
	pattern.Get("/creational/singleton", patternCreationalSingleton)
	pattern.Get("/creational/abstract-factory", patternCreationalAbstractFactory)
	pattern.Get("/creational/object-pool", patternCreationalObjectPool)
	pattern.Get("/behavioural/template-method", patternBehaviouralTemplateMethod)
	pattern.Get("/behavioural/iterator", patternBehaviouralIterator)

	api.Get("/stack", stackExamples)
	api.Get("/linklist/single", linklistSingleExamples)
	api.Get("/linklist/double", linklistDoubleExamples)
	api.Get("/tree/bst/recursive", treeBstExamples)
	api.Get("/tree/bst/iterative", treeBstIterativeExamples)
	api.Get("/tree/bst/array", treeViaArrayExamples)
	api.Get("/sort", sortExamples)

	api.Get("/copy/deep-shallow", copyExamples)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}

// Handler

func stringsExamples(c *fiber.Ctx) error {
	strings.LongestSubstring()
	return c.SendString("Strings: Longest Substring Implementation")
}

func sortExamples(c *fiber.Ctx) error {
	sort.ExampleSort()
	return c.SendString("Sort : Using Interface Implementation")
}

func copyExamples(c *fiber.Ctx) error {
	misc.CopyDeepShallow()
	return c.SendString("Copy : Deep and Shallow")
}

func stackExamples(c *fiber.Ctx) error {
	stack.StackExample()
	return c.SendString("Stack : Slice Implementation")
}

func linklistSingleExamples(c *fiber.Ctx) error {
	linklist.LinklistExample()
	return c.SendString("Linklist : Singly Implementation")
}

func linklistDoubleExamples(c *fiber.Ctx) error {
	linklist.DoublyListExample()
	return c.SendString("Linklist : Doubly Implementation")
}

func treeBstExamples(c *fiber.Ctx) error {
	tree.TreeBstExample()
	return c.SendString("Tree : BST Recursive Implementation")
}

func treeBstIterativeExamples(c *fiber.Ctx) error {
	tree.TreeBstIterativeExample()
	return c.SendString("Tree : BST Iterative Implementation")
}

func treeViaArrayExamples(c *fiber.Ctx) error {
	tree.TreeViaArrayExample()
	return c.SendString("Tree : Array Repreresentation")
}

func chanBasicExamples(c *fiber.Ctx) error {
	//channel.Basic()
	//channel.GeneratorPattern()
	channel.FanOutFanInPattern()
	channel.SemaphoreExample()
	return c.SendString("Channel Basics")
}
func chanExamples(c *fiber.Ctx) error {
	gc := channels.GoChannel{}
	//gc.RoutineOne()
	gc.CheckOsSignal()
	return c.SendString("Channel Examples")
}

func interfaceExamples(c *fiber.Ctx) error {
	interfaces.ExampleInterfaceInStruct()
	return c.JSON(map[string]interface{}{})
}

func interfaceExamples1(c *fiber.Ctx) error {

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

func embeddingExample(c *fiber.Ctx) error {
	structs.EmbeddingExample()
	return c.SendString("Embedding: Structs")
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

func mapExamples(c *fiber.Ctx) error {
	types.MapCreation()
	//types.SliceAppend()
	types.MapIterate()
	//types.SliceMultiDimention()
	return c.JSON(map[string]interface{}{"success": true, "error": nil})

}
