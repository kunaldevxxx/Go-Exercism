

# Go Packages

In Go, an application is organized in packages. A package is a collection of source files located in the same folder. All source files in a folder must have the same package name at the top of the file. By convention, packages are named to be the same as the folder they are located in.

## Importing Packages

Go provides an extensive standard library of packages which you can use in your program using the `import` keyword. Standard library packages are imported using their name.

```go
import "fmt"
```

An imported package is then addressed with the package name:

```go
world := "World"
fmt.Sprintf("Hello %s", world)
```

## Exported and Private Items

Go determines if an item can be called by code in other packages through how it is declared. To make a function, type, variable, constant, or struct field externally visible (known as exported), the name must start with a capital letter. This is analogous to how other programming languages like Java use access modifiers such as public and private.

### Example Functions

```go
// Hello is a public function (callable from other packages).
func Hello(name string) string {
    return "Hello " + name
}

// hello is a private function (not callable from other packages).
func hello(name string) string {
    return "Hello " + name
}
```
```

