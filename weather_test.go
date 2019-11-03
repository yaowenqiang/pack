package pack
import (
	"testing"
)


func BenchmarkPrintWeather(b *testing.B) {
	/*
	for i := 0; i < b.N;i++ {
		PrintWeather(101010100)
	}
	*/
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PrintWeather(101010100)
		}
	})


}

