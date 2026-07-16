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
	ans := returnfunction()
	fmt.Println(ans)

	sumOfProducts := calculateTotal(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("this is the sum of products", sumOfProducts)

	concatUserName := func(fname string, lname string) string {
		return fmt.Sprintf("%s %s", fname, lname)
	}
	fmt.Printf("This is user full name:- %v\n ", concatUserName("Niki", "naina"))

	w := Wishlist{
		Name:  "dell",
		Price: 8000,
		Stock: 8,
	}
	fmt.Println("this is the price of this much quantity of product", w.CalculateStock(3))

	// fmt.Println("after using this much remaining quantity of Wishlist:-",
	w.reduceStock(3)
	fmt.Println("final product details:-", w)

	app.Listen("localhost:9000")

}

// these are the types of the Cart and the wishlist
type Cart struct {
	Name  string
	Price float64
	Stock int
}

type Wishlist struct {
	Name  string
	Price int
	Stock int
}

// these are called receiver function

func (p Cart) Calculate(qty int) float64 {
	return p.Price * float64(qty)
}

func (w Wishlist) CalculateStock(qty int) int {
	return w.Price * qty
}

// for pass by value call it like this
// for pass by reference add pointed here in *Wishlist ()
func (w Wishlist) reduceStock(qty int) {
	if w.Stock > qty {
		w.Stock -= qty
	}
	// return w
}
func helperfunction() {
	fmt.Println("hello this is a helper function")
}

func returnfunction() string {
	return "we are testing this return function"
}

func calculateTotal(Products ...int) int {
	totalamount := 0
	for _, price := range Products {
		totalamount += price
	}
	return totalamount
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
