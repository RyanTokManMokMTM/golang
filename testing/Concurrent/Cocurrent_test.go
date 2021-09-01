//BenchMarketing Test
package main
import "testing"

func BenchmarkConcurrentAtomic(b *testing.B){
	//calculate time,need a timer here
	//reset the timer ,created in testing
	b.ResetTimer()
	println(b.N)
	//using a loop to try
	//looping time provided by Testing.B dynamically
	for i:= 0;i<b.N;i++{
		ConcurrentAtomic()
	}
}

func BenchmarkConcurrentMutex(b *testing.B){
	//calculate time,need a timer here
	//reset the timer ,created in testing
	b.ResetTimer()
	println(b.N)
	//using a loop to try
	//looping time provided by Testing.B dynamically
	for i:= 0;i<b.N;i++{
		ConcurrentMutex()
	}
}
