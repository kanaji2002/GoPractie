package main
import ("fmt")

func main() {

  const a int =10
  fmt.Println(a)

  var i,j string = "hello","world"
  fmt.Print(i," ", j)
  fmt.Println("Hello World!")
  fmt.Printf("j has value %s\n", j)

  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}

// package main 
// import ("fmt")


// fun main (){
//   fmt.Println("hello world")
// }

