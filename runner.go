package main
import (
	"fmt"
	"math/rand"
	"testing/quick"
	"time"
	"reflect"

)	

func main() {
	//val, ok :=  quick.Value(reflect.TypeOf(1.0),
	val, ok :=  quick.Value(reflect.TypeOf(MyStruct{}),
	rand.New(rand.NewSource(time.Now().Unix())))

	if ok {
		//fmt.Println(val.Float())
		//fmt.Println(val.Interface().(MyStruct))
		fmt.Println("%v", val)
	}
}

type MyStruct struct {
	MyInt int
	MyString string
	Myslice []float32
}
