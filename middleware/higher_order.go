package main

import "fmt"

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(c string) error {
		fmt.Println("do something before")
		r := c + " should be green"
		return next(r)
	}
}

func home(s string) error {
	fmt.Println("home", s)
	return nil
}

func main() {
	wrapped := Use(home)  //home กับ next(Decorator) เป็น func ที่เขียนเหมือนกัน เลยสามารถใช้ร่วมกันได้ ตามกฏของ first class
	w := wrapped("world") // โยน world ไปที่ func home ผนวกกับใช้งานกับ func Use โดยใช้ r เข้ามาตกเเต่งเพิ่มเติม
	fmt.Println("end result", w)
}
