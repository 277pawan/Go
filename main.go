package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// fmt.Println("hello")

	app := fiber.New()

	// // Routes

	// helperfunction()

	// // var myFamily [3]string

	// // myFamily[0] = "john"
	// // myFamily[1] = "jane"
	// // myFamily[2] = "bob"

	// myFamily := [3]string{"jon", "joe"}
	// myFamily[1] = "pawan"

	// // Slice
	// var myFriends []string
	// fmt.Printf("this is my family %v\n", myFamily)

	// myFriends = append(myFriends, "alice")
	// myFriends = append(myFriends, "bob")
	// fmt.Printf("theses are my friends from slice %v\n", myFriends)

	// // multi dimmensional array
	// myCourses := [2][2]int{
	// 	{1, 2},
	// 	{2, 3},
	// }
	// fmt.Println("this is our multi dimmensional array", myCourses)
	// // multi dimmensional dynamic array using slice.
	// mySliceCourse := [][]int{
	// 	{1, 2},
	// 	{2, 3},
	// 	{4, 63},
	// }
	// // newCourse := []int{5, 6}
	// mySliceCourse = append(mySliceCourse, []int{10, 12})
	// fmt.Println("this is our multi dimmensional array", mySliceCourse)

	// Using Make method code
	myWishlist := make(map[string]string)
	myWishlist["first"] = "new shoes"
	myWishlist["second"] = "900 B dollar"
	myWishlist["third"] = "new car"

	delete(myWishlist, "third")
	firstWish := myWishlist["first"]
	fmt.Println("this is my first wish:-", firstWish)
	fmt.Println("this is my wishlist:-", myWishlist)

	// Struct is basically  a object of js but we have to define it types before using it like a typescript.

	type Details struct {
		Description string
		images      string
	}
	type Product struct {
		Name    string  `json:"product_name"`
		Price   float64 `json:"price"`
		Details Details `json:"details"`
	}
	var product Product
	product = Product{
		Name:  "Nike Shoes",
		Price: 99.99,
	}
	fmt.Println("this is our first product :-", product)

	product = Product{
		Name:  "Dell Laptop",
		Price: 12000,
		Details: Details{
			Description: "this is a dell laptop image.",
			images:      "image.jpg",
		},
	}
	product.Name = "Super Dell"
	fmt.Println("this is the new product with nested details list:-", product)

	age := 20
	if age > 65 {
		fmt.Println("you are a senior citizen.")
	} else if age > 18 {
		fmt.Println("Adult citizen.")
	} else {
		fmt.Println("You are a child.")
	}

	seatClass := "firstclass"

	switch seatClass {
	case "firstclass":
		fmt.Println("You will get free drinks")
	case "secondclass":
		fmt.Println("you will get free water.")
	case "thirdclass":
		fmt.Println("You will be charged for everything.")
	}
	loopfunction()
	app.Listen("localhost:9000")
}

func helperfunction() {
	fmt.Println("hello this is a helper function")
}

func loopfunction() {
	var myFriends []string
	for i := 0; i < 10; i++ {
		myNewFriend := fmt.Sprintf("Friend %d", i)
		myFriends = append(myFriends, myNewFriend)
	}
	// fmt.Println()
	for _, value := range myFriends {
		fmt.Println(value)
	}

	left := 0
	for {
		fmt.Println(left)
		if left >= 10 {
			break
		}
		left++
	}
	fmt.Println(myFriends)
}
