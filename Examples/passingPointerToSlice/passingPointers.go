// test
package main

import (
	"fmt"
)


func alter(sl []byte) {
	sl[0] = 65
}

func alter2(sl []byte) {
	sl = append(sl, []byte(" this is an addition")...)
	fmt.Println(sl)
}
func alter3(sl *[]byte) {
	*sl = append(*sl, []byte(" this is an addition")...)
	fmt.Println(*sl)
}
func main() {
	sl := []byte("this is a test slice")
	fmt.Println(sl)
	alter3(&sl)
	fmt.Println(sl)
}

