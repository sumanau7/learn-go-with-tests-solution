package iteration

import "testing"

func Repeat(input string, times int) string {
    var repeated string
    for i:=0 ; i<times ; i++ {
        repeated += input
    }
    return repeated
}

func BenchmarkRepeat(b *testing.B){
    for i:= 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}
