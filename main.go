package main




func main() {
        println("Numbers in slice are:")
        newNum().print()
        m, x := MinMax(newNum())
        println("Min number is:", m)
        println("Max number is:", x)
        av, s := Averg(newNum())
        println("Average from numbers in slice is:", av, "\n","Sum of numbers is:", s)
}
type number []int64

func newNum() number {
    mySlice := number{}
    N := []int64 {10, -15, 20, 25, 30}
    for _, integer := range N {

            mySlice = append(mySlice, integer)
    }
    return mySlice
}
