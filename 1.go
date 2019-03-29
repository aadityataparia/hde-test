package main
import (
  "fmt"
)

func ReadN(total int, n int) {
  if (n == 0) {
    fmt.Printf("%d\n", total)
  } else {
    var num int
    fmt.Scan(&num)

    if (num > 0) {
      total = total + num * num
    }

    ReadN(total, n - 1)
  }
}

func iterateInput(n int) {
  if ( n == 0 ) {
  } else {
    var numbers int
    fmt.Scan(&numbers)

    ReadN(0, numbers)
    n--
    iterateInput(n)
  }
}

func main() {
  var n int
  fmt.Scan(&n)

  iterateInput(n)
}
