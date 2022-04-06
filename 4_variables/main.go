package main

import "fmt"

// Biến trong Golang là một cái tên đại diện cho một địa chỉ trong bộ nhớ, nơi mà một giá trị cụ thể được lưu trữ
// Biến được tạo ra trong Runtime
// Biến khai báo ra bắt buộc phải được sử dụng
// Trong Golang để biểu thị một biến không sử dụng ta có thể sử dụng ký tự _

func main() {
	// Cú pháp khai báo biến đầy đủ trong Golang, sử dụng từ khóa var: var <tên_biến> <kiểu_dữ_liệu>
	var number int
	number = 10

	// Cú pháp khai báo ngắn gọn, Golang sẽ dựa vào giá trị mà bạn gán cho biến để quyết định xem kiểu dữ liệu của
	// biến đó là gì: <tên_biến> := <giá_trị_của_biến>
	secondNumber := 20

	// Khai báo một biến và không sử dụng nó
	var thirdNumber int
	// Khi compile chương trình sẽ báo lỗi do biến thirdNumber được khai báo nhưng không được sử dụng, để
	// quá trình compile không xảy ra lỗi ra có thể gán cho biến đó là một biến không sử dụng bằng cách sử dụng ký
	// tự _
	_ = thirdNumber

	// Trong Golang bạn hoàn toàn có thể khai báo nhiều biến cùng một lúc
	// Khai báo hai biến cùng kiểu dữ liệu
	age, id := 25, 01
	// Khai báo hai biến khác kiểu dữ liệu
	name, yearOfBirth := "Quang Trung", 1997

	// Khai báo nhiều biến sử dụng từ khóa var
	// khai báo 2 biến cùng kiểu dữ liệu
	var fourthNumber, fifthNumber int
	_, _ = fourthNumber, fifthNumber
	// khai báo hai biến khác kiểu dữ liệu
	var (
		sixthNumber int
		address     string
	)
	_, _ = sixthNumber, address

	// Sự khác biệt giữa := và =
	// khi khai báo một biến mới ta sẽ sử dụng cặp ký tự := để tạo một biến mới và gán giá trị cho nó
	fullName := "Le Quang Trung"
	// đối với một biến đã được khai báo, ta muốn gán lại giá trị cho biến đó, sử dụng ký tự =
	fullName = "Le Quang trung 1"

	// Thay đổi giá trị hai biến trong Golang(swapping variables)
	// nhờ việc hỗ trợ gán giá trị nhiều biến cùng một lúc nên trong Golang cú pháp để có thể hoán đổi giá trị 2 biến
	// cho nhau cực kỳ dễ dàng
	// khảo tạo 2 biến kiểu dữ liệu số
	seventhNumber, eighthNumber := 10, 20
	fmt.Println("Before swapping variables", seventhNumber, eighthNumber)
	seventhNumber, eighthNumber = eighthNumber, seventhNumber
	fmt.Println("After swapping variables", seventhNumber, eighthNumber)

	fmt.Println(number)
	fmt.Println(secondNumber)
	fmt.Println(age, id)
	fmt.Println(name, yearOfBirth)
	fmt.Println(fullName)
}
