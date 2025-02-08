package mathutils 

import (
	"fmt"
)

func Div(num1,num2 float64) (float64,error){
	if num2 == 0{
	 return 0,fmt.Errorf("Error : Division by zero is not allowed")
	}
	return num1 / num2, nil
 }