package main
import("fmt")
func main(){
myslice := []int {1,2,3,3}

fmt.Println(myslice)
fmt.Println(myslice[0])


var a int = 10
var b int = 20

if a>b {

fmt.Println("a is greater than b")
}else {
fmt.Println("a is less than b")
}


var c int = 30
switch c {
case 10:
fmt.Println("c is 10")
case 20:
fmt.Println("c is 20")
case 30:
fmt.Println("c is 30")
default:
fmt.Println("c is not 10, 20 or 30")
}

a&=5
fmt.Println(a)


for i:=0; i<1; i++ {
fmt.Println("hello world")
}

}