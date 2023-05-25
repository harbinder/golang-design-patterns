package creational

/*
refer link : https://golangbyexample.com/prototype-pattern-go/

It is a creational design pattern that lets you create copies of objects.
In this pattern, the responsibility of creating the clone objects is delegated to the actual object to clone.

The object to be cloned exposes a clone method which returns a clone copy of the object

When to Use
We use prototype pattern when the object to be cloned creation process is complex i.e the cloning may involve vases of
handling deep copies, hierarchical copies, etc. Moreover, there may be some private members too which cannot be directly
accessed.
A copy of the object is created instead of creating a new instance from scratch. This prevents costly operations involved
while creating a new object such as database operation.
When you want to create a copy of a new object, but it is only available to you as an interface. Hence you cannot directly
create copies of that object.
*/
