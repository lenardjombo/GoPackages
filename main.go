package main


import (
    "fmt"
    "CalcGo/mathutils"
    "CalcGo/inputs"
    )


func main(){
    fmt.Println("Welcome to the Go Calculator")
    fmt.Println("Operations \n1.Addition\n2.Subtraction\n3.Multiplication\n4.Division")

    if !inputs.GetUserInput(){
        fmt.Println("Invalid Inputs. Exiting program.")
        return
        //os.Exit(1)
        //Return will exit the program
    }

    // Perform the operation based on the user input
    switch inputs.Operator {
    case "1":
        fmt.Printf("Addition Result: %v\n", mathutils.Add(inputs.Num1, inputs.Num2))
    case "2":
        fmt.Printf("Subtraction Result: %v\n", mathutils.Subtract(inputs.Num1, inputs.Num2))
    case "3":
        fmt.Printf("Multiplication Result: %v\n", mathutils.Mul(inputs.Num1, inputs.Num2))
    case "4":
        result,err := mathutils.Div(inputs.Num1,inputs.Num2)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("Division Result: %v\n", result)
        }

    default:
        fmt.Printf("Please try again...Reload Failed")
    }
}









