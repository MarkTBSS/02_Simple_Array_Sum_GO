/* var ar [6]int32
ar = [6]int32{1, 2, 3, 4, 10, 11}
//fmt.Println("ar[5] : ", ar[5])
// ใช้ Loop for วนตั้งแต่ index 0 จนหมด
for i, value := range ar {
	fmt.Printf("i ตำแหน่งที่ %d มีค่าเท่ากับ : %d\n", i, value)
} */

package main

import "fmt"

func main() {

	var slice []int32
	slice = []int32{1, 2, 3, 4, 10, 11}

	//slice = append(slice, 12) // เติมค่าเข้าไปหลังสุด
	//slice = slice[:len(slice)-1] // ตัดตัวหลังสุดออก 1 ตัว
	//slice = slice[1:] // ตัดตัวหน้าสุดออก 1 ตัว
	//copySlice := make([]int32, len(slice)) // สร้าง (make)
	//copy(copySlice, slice)                 // Copy

	copySlice := slice
	copySlice[0] = 100

	fmt.Println("แสดงค่าใน copySlice")
	for i, value := range copySlice {
		fmt.Printf("i ตำแหน่งที่ %d มีค่าเท่ากับ : %d\n", i, value)
	}
	fmt.Println("แสดงค่าใน slice เก่า")
	for i, value := range slice {
		fmt.Printf("i ตำแหน่งที่ %d มีค่าเท่ากับ : %d\n", i, value)
	}

}
