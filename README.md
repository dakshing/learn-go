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

for i, element := range slice {
    fmt.Println(i, element)
}
```

**_NOTE:_** Go is not an object-oriented programming language like Java.