package behavioural

import (
	"fmt"
)

/*
We want to create an iterator for collection of user struct.

The main idea behind this pattern is to expose the iteration logic of a Collection struct into a different
object (which implements the iterator interface). This iterator provides a generic method of iterating over a
collection independent of its type.
*/
type user struct {
	name string
	age  int
}

// concrete  implementation of collection iterface
type UserCollection struct {
	users []*user
}

// concrete  implementation of iterator interface
type UserIterator struct {
	index int
	users []*user
}

type iterator interface {
	hasNext() bool
	getNext() *user
}

type collection interface {
	createIterator() iterator
}

func (ui *UserIterator) hasNext() bool {
	if ui.index < len(ui.users) { // ui.users[ui.index] != nil {
		return true
	}
	return false
}
func (ui *UserIterator) getNext() (u *user) {
	if ui.hasNext() { //ui.users[ui.index] != nil {
		u = ui.users[ui.index]
		ui.index++
	}
	return
}

func (uc *UserCollection) createIterator() (i iterator) {
	i = &UserIterator{users: uc.users}
	return
}

func ExecuteIterator() {
	user1 := &user{"Harry", 41}
	user2 := &user{"Garry", 42}

	userCollection := &UserCollection{users: []*user{user1, user2}}
	iterateUsers := userCollection.createIterator()
	fmt.Println("Iterate Users")
	for iterateUsers.hasNext() {
		u := iterateUsers.getNext()
		fmt.Println(u)
	}
}
