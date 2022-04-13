Go Notes

* In this tutorial, you wrote functions that you packaged into two modules: one with logic for sending greetings; the other as a consumer for the first.
* https://go.dev/doc/tutorial/create-module 

* Any Go function can return multiple values
* Common error handling in Go, is to return an error as a value so the caller can check for it
* Go slice - a slice is like an array, except that its size changes dynamically as you add and remove items
* In Go, you initialize a map with the following syntax: make(map[key-type]value-type)
* In Go, Test function names have the form TestName, where Name says something about the specific test. 
* test functions take a pointer to the testing package's testing.T type as a parameter. You use this parameter's methods for reporting and logging from your test.
* go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go)
* While the go run command is a useful shortcut for compiling and running a program when you're making frequent changes, it doesn't generate a binary executable