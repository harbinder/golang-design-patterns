package structs

/*
Reference Link
https://golangbyexample.com/method-in-golang/

###### Methods #####

A method in golang is nothing but a function with a receiver.
A receiver is an instance of some specific type such as struct, but it can be an instance of any other custom type.
So basically when you attach a function to a type, then that function becomes a method for that type.
The method will have access to the properties of the receiver and can call the receiver’s other methods.

Why Method
Since method lets you define a function on a type, it lets you write object-oriented code in Golang.
There are also some other benefits such as two different methods can have the same name in the same package which is not possible with functions

Format of a Method
Below is the format for a method

func (receiver receiver_type) some_func_name(arguments) return_values
The method receiver and receiver type appear between the func keyword and the function name.
The return_values come at the last.

Also, let’s understand more differences between a function and a method. There are some important differences between them.
Below is the signature of a function

Function:

func some_func_name(arguments) return_values
We have already seen the signature of a method

Method:

func (receiver receiver_type) some_func_name(arguments) return_values
From the above signature, it is clear that the method has a receiver argument.
This is the only difference between function and method, but due to it they differ in terms of functionality they offer

A function can be used as first-order objects and can be passed around while methods cannot.
Methods can be used for chaining on the receiver while function cannot be used for the same.
There can exist different methods with the same name with a different receiver,
but there cannot exist two different functions with the same name in the same package.
*/
