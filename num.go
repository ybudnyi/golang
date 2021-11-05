package main

//import "fmt"

type number []int

func newNum() number {
    mySlice := number{}
    N := []int {10, -15, 20, 25, 30}
    for _, integer := range N {

            mySlice = append(mySlice, integer)
    }
    return mySlice
}


func (n number) print() {
    for _, i := range n {
            println(i)
    }
}

func MinMax(n number) (int, int) {
var max int = n[0]
var min int = n[0]
for _, value := range n {
    if max < value {
        max = value
    }
    if min > value {
        min = value
    }
}
    return min, max
}

func Averg(n number) (float64, float64) {
sum := 0   
l := len(n)
for k := range n {
    sum += n[k]
}
avg := (float64(sum) / float64(l))
return float64(avg), float64(sum)
}