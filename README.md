# Understanding Go Packages
Learn and understand Go Packages with this one Repository
##  Overview
CalcGo is a simple calculator project built in Go, demonstrating how to organize code using **Go packages**. Each function (addition, subtraction, multiplication, division, etc.) is placed in its **own independent file** within a structured package format.

##  Project Structure
```
CalcGo/ 
│── main.go          # Entry point of the program 
│── go.mod          # Module file for dependency management 
│── operations/ 
│   ├── Add.go      # Addition function 
│   ├── Sub.go # Subtraction function 
│   ├── Mul.go # Multiplication function 
│   ├── Div.go   # Division function 
```

## Using Go Packages

### 1️⃣ **Creating a Package**
Each function is defined inside an `operations` package. For example, `add.go` contains:

```go
package operations

func Add(num1, num2 float64) float64 {
    return num1 + num2
}
```

### 2️⃣ **Importing a Package**
To use the operations package in `main.go`, import it:

```go
package main

import (
    "fmt"
    "CalcGo/mathutils"
)

func main() {
    result := mathutils.Add(5, 3)
    fmt.Println("Sum:", result)
}
```

### 3️⃣ **Running the Project**
Run the Go program using:

```sh
go run main.go
```

##  Key Takeaways
- **Modular Code**: Each function is in a separate file within the `mathutils` package.
- **Reusability**: Functions can be reused across different parts of the project.
- **Scalability**: New operations can be added easily by creating new files in `operations/`.
- **Maintanability**: Independent source files would be easier to maintain 

---

 **Author**: Lenard Jombo  
 **GitHub Repo**:[GoPackages](https://github.com/lenardjombo/GoPackages)  

Happy Coding! 
