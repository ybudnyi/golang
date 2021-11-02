package main




func main() {
        numFunc := newNum()
        println("Numbers in slice are:")
        numFunc.print()
        m, x := MinMax(numFunc)
        println("Min number is:", m)
        println("Max number is:", x)
        av, s := Averg(numFunc)
        println("Average from numbers in slice is:", av, "\n","Sum of numbers is:", s)
}