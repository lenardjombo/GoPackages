package inputs

//validate user inputs : num1 and num2
func ValidateUserInput(num1,num2 int)bool{
    return  num1 > 0 && num2 > 0
    //Returns true if the inputs are valid
    //Returns false if the inputs are invalid
    //In this case, the inputs are invalid if either of the numbers is less than 0
}