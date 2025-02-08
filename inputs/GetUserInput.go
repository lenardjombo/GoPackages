package inputs

import "fmt"


var Num1,Num2 int
var Operator string
func GetUserInput()bool{
    fmt.Print("Please enter an operator from the above list: ")
    fmt.Scan(&Operator)
    fmt.Print("Please enter the first number: ")
    fmt.Scan(&Num1)
    fmt.Print("Please enter the second number : ")
    fmt.Scan(&Num2)

    return ValidateUserInput(Num1,Num2)
    //Returns true if the inputs are valid
    //Returns false if the inputs are invalid
    //In this case, the inputs are invalid if either of the numbers is less than 0
    //The ValidateUserInput function is defined in the ValidateUserInput.go file
}
