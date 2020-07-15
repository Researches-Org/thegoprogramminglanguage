package main

import "fmt"

// nil represents an empty list
type IntList struct {
    Value int
    Tail * IntList
}

func (list *IntList) Sum() int {
    if list == nil {
        return 0
    }
    return list.Value + list.Tail.Sum()
}

func main() {
    list := IntList{
        1,
        &IntList{
            2,
            nil,
        },
    }

    fmt.Println(list.Sum())    

    m := make(map[string][]string)

    fmt.Println(m)

    fmt.Println(m["1"])
    
    m["1"] = make([]string, 0)
 
    fmt.Println(m["1"])

    m["1"] = append(m["1"], "1")

    m = nil

    fmt.Println(m["1"])

    // panic assignment to entry in nil map
    // m["1"] = make([]string, 0)
}

