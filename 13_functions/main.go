package main

import "fmt"

// Golang khuyến cáo viết tên của function bằng các từ đơn giản theo cú pháp camelCase
// Trong cùng một package tên của các function phải là duy nhất
// Function trong Golang hỗ trợ tính năng trả về nhiều giá trị
// Golang không hỗ trợ nạp chồng function
// Cú pháp
// func (receiver) name(parameters) (returns) {
//	-> TODO logic here
// }

// Định nghĩa function tính tổng 2 số
func sum(a int, b int) int {
	return a + b
}

// Định nghĩa function trả về nhiều giá trị
func getInfo() (string, int) {
	return "Le Quang Trung", 26
}

func main() {
	fmt.Println("Sum of 5 and 7:", sum(5, 7))
	name, age := getInfo()
	fmt.Println("Get info function:", name, age)

	// Các tham số được truyền vào trong function sẽ được copy ra một bản khác, giá trị tính toán trong function là giá trị bản copy
	var year int = 2022
	var fullName string = "Le Quang Trung"
	fmt.Println("The address of year:", &year)
	fmt.Println("The address of fullName:", &fullName)
	// Do là bản copy nên việc thay đổi giá trị của biến trong function sẽ không làm ảnh hưởng đến các biến ban đầu
	showAddress(year, fullName)
	fmt.Println("The value of year:", year)
	fmt.Println("The value of fullName:", fullName)

	// TODO
	firstSlice := []string{"Trung"}
	fmt.Println("The value of firstSlice:", firstSlice)
	addItem(firstSlice)
	fmt.Println("The value of firstSlice:", firstSlice)

	// Như các function ở trên việc return giá trị trả vè thường khai báo biến trước đó rồi return về giá trị cảu biến đó
	// Golang cho phép ta định nghĩa tên biến ngay trong phần định nghĩa gá trị trả về gọi là Naked Returns
	name, age = demoNakedReturn()
	fmt.Println("Demo Naked Return:", name, age)

	// Các ví dụ ở trên đều trong trường hợp truyền số lượng tham số đầu vào được xác định trước
	// Trong trường hợp ta không biết số lượng tham số đâu vào, Golang hỗ trợ một loại function là Variadic Function
	demoVariadicFunction(1, 2, 3, 4, 5, 6)
	// Khi sử dụng Variadic function thực chất là ta truyền vào function một slice với một kiểu dữ liệu được xác định trước
	// Trường hợp không truyền tham số vào function thì tức là một slice nil
	demoVariadicFunction()

	addItemVariadicFunction(firstSlice...)
	fmt.Println(firstSlice)
}

func showAddress(year int, fullName string) {
	fmt.Println("The address of year in function:", &year)
	fmt.Println("The address of fullName in function:", &fullName)

	year++
	fullName = fullName + "123"
	fmt.Println("The value of year in function:", year)
	fmt.Println("The value of fullName in function:", fullName)
}

func addItem(slice []string) {
	//slice = append(slice, "Quang")
	slice[0] = "Quang1"
	fmt.Println("The value of firstSlice in function:", slice)
}

func demoNakedReturn() (name string, age int) {
	name = "Trung"
	age = 38

	return
}

func demoVariadicFunction(inputs ...int) {
	fmt.Printf("Demo variadic function: %T\n", inputs)
	fmt.Printf("Demo variadic function value: %#v\n", inputs)
}

func addItemVariadicFunction(slice ...string) {
	slice[0] = "Quang"
}
