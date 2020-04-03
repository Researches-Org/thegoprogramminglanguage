package main

import (
  "fmt"
)

const (
  MONTHS = 12
)

func main () {

  months := [...]string {
    1: "January",
    2: "February",
    3: "March",
    4: "April",
    5: "May",
    6: "June",
    7: "July",
    8: "August",
    9: "September",
   10: "October",
   11: "November",
   12: "December",
  }

  Q2 := months[4:7]
  summer := months[6:9]
  
  fmt.Println(Q2)
  fmt.Println(summer)

  fmt.Println(months)

  fmt.Println(common(Q2, summer))

}

func common(m1 []string, m2 []string) []string {
  c := []string{MONTHS: ""} 
  i := 0
  for _, s:= range m1 {
    for _, q := range m2 {
      if s == q {
        c[i] = s
        i++
      }
    }
  }
  return c[:i]
}

