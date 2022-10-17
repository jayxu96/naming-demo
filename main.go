package main

import (
	"fmt"

	naming "github.com/jayxu96/baby-naming/naming"
	naming2 "github.com/jayxu96/baby-naming/v2/naming"
)

func main() {
    var version int
    fmt.Println("***** Welcome to Baby Naming Generator *****")
    fmt.Println("Please select Generator version: 1,2,3")
    fmt.Scanln(&version)

    var lastName string
    var gender string
    var male bool
    if version == 1 {
        fmt.Println("***** This is a Baby Naming Generator V1 *****")
        fmt.Println("Please enter your Last Name: " )

        fmt.Scanln(&lastName)
        fmt.Println("The baby name is " + naming.CreateBabyName(lastName))
    } else if version == 2 {
//  ===================== v2 ========================

        fmt.Println("***** This is a Baby Naming Generator *****")
        fmt.Println("1. Please enter your Last Name: " )
        fmt.Scanln(&lastName)
        fmt.Println("2. Is your baby a boy or girl?")
        fmt.Scanln(&gender)
        fmt.Println("Your baby is a " + gender)
        if gender == "boy" {
           male = true
        } else if gender == "girl" {
          male = false
        } else {
          fmt.Println ("Sorry, this generator does not support generate names for other genders.")
          return
        }
        fmt.Println("Congrats! The baby name is " + naming2.CreateBabyName(male, lastName))
    } else if version == 3 {
        fmt.Println("Congrats! The baby name is " + naming2.CreateBabyName(male, lastName))
    } else {
        fmt.Println ("Sorry, no such version")
        return
    }

}