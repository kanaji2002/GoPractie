package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n"
  fmt.Printf("a[brand]\t%v\n", a["brand"])
  fmt.Printf("a[model]\t%v\n", a["model"])
  fmt.Printf("a[year]\t%v\n", a["year"])
  fmt.Printf("b[Oslo]\t%v\n", b["Oslo"])
}


