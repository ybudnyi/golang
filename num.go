package main


func (n number) print() {
    for _, i := range n {
            println(i)
    }
}

func MinMax(n number) (int64, int64) {
var max int64 = n[0]
var min int64 = n[0]
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
var sum int64 = 0   
l := len(n)
for k := range n {
    sum += n[k]
}
avg := (float64(sum) / float64(l))
return float64(avg), float64(sum)
}