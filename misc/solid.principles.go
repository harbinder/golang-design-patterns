package misc

/*
Reference: https://www.youtube.com/watch?v=XI7zep97c-Y

SOLID Principles: Rules to follow to write better code
Advantages of these are:
- Avoid Duplicate Code (resuable)
- Easy to maintain
- Easy to understand
- Flexible software
- Reduce Complexity

S: Single Responsibility
O: Open-Close
L: Liskov Substitution
I: Interface Segmented
D: Dependency Inversion

1. Single Responsibility
Each class should have single reason to change.
There should not be multiple things done in 1 class

2. Open-Closed
Each class should be OPEN to be EXTENED but CLOSED to be MODIFIED
Extending a class will not impact the existing code.

3. Liskov Substitution
Is a class B is sub-type of class A,
Then we should be able to REPLACE object of
Class A with B, without BREAKING the
Behaviour of the program

4. Interface Segmented
Interfaces should be such,
that each Class should NOT implement unnecessary functions they do not need,
which are declared in interface.
This means, we should further segregate the interfaces, if required.

5. Dependency Inversion
A Class should NOT be dependent on Concrete Classes directly,
it should depend on the Interface of those Concrete Classes
*/
