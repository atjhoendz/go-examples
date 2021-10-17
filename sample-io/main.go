package main

import "fmt"

func main() {
  fmt.Println("Sample Input Output")

  var name string

  fmt.Print("Input your name: ")

  // notes: cannot scan interface{} type
  _, err := fmt.Scanf("%s", &name)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("Welcome %s!\n", name)
  
  fmt.Println("-----")

  fmt.Println("Calculate Area of Rectangle")
  var p, l, area float64

  fmt.Print("Long: ")
  _, err = fmt.Scanf("%f", &p)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Print("Wide: ")
  _, err = fmt.Scanf("%f", &l)

  if err != nil {
    fmt.Println(err)
  }

  area = p * l

  fmt.Printf("\nArea of a Rectangle with long %.2f and wide %.2f is %.2f\n", p, l, area)
}

