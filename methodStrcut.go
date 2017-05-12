package main

import (

	"fmt"

)
type Rect struct {
	x,y int
}


func  calrect(r Rect) int {
	var x =5
	r.x =x
	return r.x*x
}
func main()  {
	rr := Rect{0,0}
	fmt.Println(calrect(rr))

}
func init()  {
	
}