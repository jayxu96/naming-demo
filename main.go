package main

import (
	"fmt"

	naming "github.com/jayxu96/baby-naming/naming"
)

func main() {
    var lastName string;
    fmt.Println("***** This is a Baby Naming Generator V1 *****")
    fmt.Println("Please enter your Last Name: " )

    fmt.Scanln(&lastName)
    fmt.Println("The baby name is " + naming.CreateBabyName(lastName))


}