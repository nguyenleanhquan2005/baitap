package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var number int = 100
	a := []int{100, 120}
	slice = append(slice, a...)
	slice = append(slice, number)
	fmt.Println(slice)
	fmt.Println(so4("himalaya"))
	fmt.Println(so4("taj"))
	fmt.Println(so4("1,2,3,4"))
	var c int = 1000
	var d []int
	d = []int{1, 2, 3, 4, 5}
	fmt.Println(c, d)
	ketqua := appendSlice(d, 2)
	fmt.Println(ketqua)
	dapan := mergeSlice(a, d)
	fmt.Println(dapan)
	result := prependSlice(d, c)
	fmt.Println(result)
	la := so2(d)
	fmt.Println(la)
	resultat := checkInclude1(22, d)
	fmt.Println(resultat)

}
func appendSlice(intSlice []int, number int) []int {
	intSlice = append(intSlice, number)
	return intSlice
}

func prependSlice(intSlice []int, number int) []int {
	abSlice := []int{number}
	abSlice = append(abSlice, intSlice...)
	return abSlice
}

func mergeSlice(slice1 []int, slice2 []int) []int {
	slice1 = append(slice1, slice2...)
	return slice1
}

func so2(slice2 []int) []int {
	slice2 = append(slice2[:3], slice2[4:]...)
	return slice2
}

func so4(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
func checkInclude1(a int, slice []int) bool {
	result := false
	for _, v := range slice {
		if a == v {
			result = true
			break
		}
	}
	return result
}

func trim() {
	var str string = "a this is a string b"
	fmt.Println("The given string is:\n", str)
	fmt.Println()
	var left, right int
	for i, c := range str {
		if c != 'a' {
			left = i
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != 'b' {
			right = i
			break
		}
	}
	trimmed := str[left : right+1]
	fmt.Println("The string obtained after trimming the above string from both sides is:\n", trimmed)
}

//so1 them phan tu vao dau slice
//so2 loại bo phan tu 3 cua slice
//so3 viet func input slice,n output bo phan tu thu n
//so4 input func string , output chuoi dao nguoc. vi du: "Nguyen" output "neyugN"
//so5 input slice , output slice dao nguoc. vidu: 3,2,5 output 5,2,3
//so6 input slice int,n int kiem tra so n co thuoc slice. vidu: n=200, tra ve false. n=3, tra ve true
