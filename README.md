# HelloIteratorDesignPattern

> [Source](https://golangbyexample.com/go-iterator-design-pattern/)

## What does `iterator` do

Like `for` , `range` it provides a way to traverse the collection but without directly manipulating the underlying data.
The traverse logic is encapsulated in the iterator.

## File structure

```
type collection interface {
	createIterator() iterator
}
```

Define the interface of collection with iterator.

```
type iterator interface {
	hasNext() bool
	getNext() *user
}
```

Define the iterator interface.

`user` , `userCollection` and `userIterator` are the concrete implementation.