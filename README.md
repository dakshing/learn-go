### Go cmd Commands
```
go build -> only compile
go run -> compile and execute
go fmt -> format all code
go install -> compile and "install"
go test -> run tests
go get -> download source code 
```

### Two Types of Packages
- Executable 
    - package named 'main' is only executable. Other packages are not executable.
- Reusable
    - packages named other than 'main' can be used as dependencies.

### Links for Go Dev
- Go Standard Library packages: https://pkg.go.dev/std 
- Go Playground: https://go.dev/play/

### Variable Declaration
```go
var str string = "Sample" // Variable declaration (long-form)
str := "Sample" // Shorthand declaration (implicit type)
```

### Function Declaration 
```go
func newCard() string {     // func <receiver> <func_name>(<params>) <return_type> {...}
    return "New Card"
}
```

### Arrays and Slice
- Arrays: Fixed length and the elements are of same type
- Slice: Variable size and the elements are of same type

```go
// Array Declaration
arr := [3]string {"one", "two", "three"}

// Slice Declaration
slice := []string {"one", "two", "three"}
slice = append(slice, "four", "five")

// Range selector syntax
slice[1:3] // Returns 1st and 2nd elements
slice[:3]  // Returns first 3 elements

for i, element := range slice {
    fmt.Println(i, element)
}
```

**_NOTE:_** Go is not an object-oriented programming language like Java.

### Function return values
- Go supoorts mutiple return values
- A return value can be ignored using '_' underscore variable during assignment

### Useful String Utils from strings package
- strings.Contains(s, substr string) bool
- strings.Index(s, substr string) int
- strings.Replace(s, old, new string, n int) string
- strings.ToLower(s string) string
- strings.ToUpper(s string) string
- strings.Split(s, sep string) []string
- strings.Join(elems []string, sep string) string
- strings.Builder
```go
var sb strings.Builder
sb.WriteString("String ")
sb.WriteString("Building ")
fmt.Println("Builder result:", sb.String())
```

### Type Casting
- Convert one type to another. Below code convert string slice to byte slice.
`[]byte(strings.Join([]string{"one", "two", "three"}, ", "))`
- Type declarations can freely be casted to parent type.
```go
type deck []string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

### Swapping
Swapping in Go is as simple as:
`a, b = b, a`

### Go Tests
- Run tests using `go test`
- Test files end with '_test.go'

### Structs in Go
- A simple struct declaration:
```go
type person struct {
	firstName string
	lastName  string
}
```
- To print struct aling with the field names use formatting '%+v'.

### Pointers
- Receivers are copy of the original value, not reference.
- Pointers are used to pass values by reference.
- '&' is used to get the address of the variable.

**_NOTE:_** Slice stores the **ptr to head** only. So updates will be retained.

| Pass by Value Types | Pass by Reference Types |
| --- | --- |
| int | slices |
| float | maps | 
| string | channels | 
| bool | pointers | 
| structs | functions | 

### Interfaces
- A special type keyword - 'interface'
- Implicit implementation - there is no 'implements'. If all the methods in the interface are defined for a type, then the type becomes honorary member of the interface.
- Go does not have generic types
- Go does not allow overloading (complains functions with same name)
- Interfaces can be used within interfaces and structs
```go
type ReadCloser interface {
    Reader // another interface
    Closed // another interface
}
```

- Interface usage style in Go - [link](./interfaces/README.md)