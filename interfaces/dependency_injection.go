// Dependency Injection

// Interfaces let you swap implementations, making testing easier. 
// Here, we define a `Database` interface with a `Save` method. 
// We then create two implementations: `MySQL` for real database operations and `MockDB` for testing. 
// The `ProcessData` function accepts any `Database` implementation, allowing us to inject different behaviors.
package main

import "fmt"

type Database interface {
	Save(data string)
}

type MySQL struct{}

func (MySQL) Save(data string) { fmt.Println("Saving to MySQL:", data) }

type MockDB struct{}

func (MockDB) Save(data string) { fmt.Println("Mock saving:", data) }

func ProcessData(db Database) {
	db.Save("Important Info")
}

func main() {
	realDB := MySQL{}
	mock := MockDB{}

	ProcessData(realDB)
	ProcessData(mock)
}
