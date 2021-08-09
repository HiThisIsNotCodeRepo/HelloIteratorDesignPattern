package main

import "fmt"

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}
	anUserCollection := &userCollection{users: []*user{user1, user2}}
	anIterator := anUserCollection.createIterator()
	for anIterator.hasNext() {
		tmpUser := anIterator.getNext()
		fmt.Printf("User is %v\n", tmpUser)
	}
}
