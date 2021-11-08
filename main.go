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