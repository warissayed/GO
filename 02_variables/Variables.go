package main

import "fmt"

//first letter is capital then it is a public variable
const LoginToken string = "this is the token Public"//public
const loginToken1 string = "this is the token Not Public"//not Public

func main() {
	var username string = "this is the name "
	fmt.Println("Variables")	
	fmt.Println(username)
	fmt.Printf("Variable is a type of : %T \n" , username)
	var number int = 10 
	fmt.Println(number)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	var smallVal uint8 = 255 
	fmt.Println(smallVal)

	var smallFloat float64 = 255.454545454545
	fmt.Println(smallFloat)

	//dealt values and some aliases
	var anotherVariable bool 
	fmt.Println(anotherVariable)			
	fmt.Printf("Variable is a type of : %T \n", anotherVariable)

	//implicit type of naming a variable
	var website = "this is the website"
	fmt.Printf( "The variable is of type : %T \n", website )

	// not using var
	numberOfUser := 300000.00
	fmt.Println(numberOfUser)	
	fmt.Println(LoginToken,loginToken1)
}
