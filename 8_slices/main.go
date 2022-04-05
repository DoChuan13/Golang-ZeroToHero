package main

import "fmt"

// Cũng giống như Array thì Slice là một tập hợp, được đánh chỉ mục(index) lưu trữ tập hợp các phần tử có cùng kiểu dữ liệu với nhau
// Kích thước cảu Slice là động, nó có thể tăng hoặc giảm và nó có thể được thay đổi trong thời gian chạy(Runtime)
// Giá trị mặc định của một Slice là nil

func main() {
	// Khai báo một Slice
	var firstSlice []int
	fmt.Printf("The default value of firstSlice: %#v\n", firstSlice)
	fmt.Println("Length of firstSlice:", len(firstSlice))

	// Khởi tạo và gán giá trị cho một Slice
	secondSlice := []int{1, 2, 3, 4}
	fmt.Printf("The default value of secondSlice: %#v\n", secondSlice)

	thirdSlice := make([]int, 4)
	fmt.Printf("The default value of thirdSlice: %#v\n", thirdSlice)

	// Duyệt từng phần tử xong Slice
	for index, value := range secondSlice {
		fmt.Printf("Index %d, value %d\n", index, value)
	}

	// So sánh 2 Slice
	fourthSlice := []int{1, 2, 3, 4}
	fifthSlice := []int{1, 2, 3, 4}
	// Không thể sử dụng toán tử == để so sanh 2 Slice
	isEqual := true
	if len(fourthSlice) != len(fifthSlice) {
		isEqual = false
		goto result
	}
	for index, _ := range fourthSlice {
		if fourthSlice[index] != fifthSlice[index] {
			isEqual = false

			break
		}
	}

result:

	if isEqual {
		fmt.Println("fourthSlice and fifthSlice are equal")
	} else {
		fmt.Println("fourthSlice and fifthSlice are not equal")
	}

	// Thêm phần tử vào Slice
	fifthSlice = append(fifthSlice, 10)
	fmt.Printf("The value of fifthSlice: %#v\n", fifthSlice)
	// Thêm nhiều phần tử vào Slice
	fifthSlice = append(fifthSlice, 10, 20, 30, 40)
	fmt.Printf("The value of fifthSlice: %#v\n", fifthSlice)
	// Thêm toàn bộ phần tử của Slice A vào Slice B
	fifthSlice = append(fifthSlice, fourthSlice...)
	fmt.Printf("The value of fifthSlice: %#v\n", fifthSlice)

	// Copy Slice
	sourceSlice := []string{"Trung", "Le", "Quang"}
	destSlice := make([]string, len(sourceSlice))
	length := copy(destSlice, sourceSlice)
	fmt.Printf("Source Slice %#v, Dest Slice %#v, Length of Slice %d\n", sourceSlice, destSlice, length)

	// Duyệt Slice với một khoảng index được cấu hình trước
	startIndex := 1
	finishIndex := 3
	temp := fifthSlice[startIndex:finishIndex]
	fmt.Printf("The value of temp: %v, Type of temp: %T\n", temp, temp)
	// Để trống finishIndex để lấy toàn bộ phần tử từ startIndex đến hết
	fmt.Println("The value of startIndex to end: ", fifthSlice[startIndex:])
	// Để trống startIndex để lấy toàn bộ phần tử từ đầu đến finishIndex
	fmt.Println("The value of start to finishIndex: ", fifthSlice[:finishIndex])
	// Để trống cả startIndex và finishIndex để lấy toàn bộ phần tử
	fmt.Println("The value of fifthSlice: ", fifthSlice[:])

	// Để thay thế toàn bộ phần tử từ một vị trí index bằng một giá trị X
	fmt.Println("The value of fifthSlice before:", fifthSlice)
	fifthSlice = append(fifthSlice[:5], 999)
	fmt.Println("The value of fifthSlice after:", fifthSlice)

	// Cấu trúc của một Slice
	// Khi khởi tạo một Slice, đồng thời Golang cũng khởi tạo một Array đằng sau Slice được gọi là Backing Array
	// Backing Array là nơi lưu trữ các phần tử
	// Golang implement một Slice giống như một Data Structure được gọi là Slice Header
	// Slice header bao gồm 3 thành phần sau
	// + Pointer: địa chỉ con trỏ của Backing Array
	// + Length: độ dài của Slice, len() trả về giá trị này
	// + Capacity: dung lượng của Slice, cap() trả về giá trị này
	// Một nil Slice sẽ không có Backing Array
	sixthSlice := []int{1, 2, 3, 4}
	seventhSlice, eighthSlice := sixthSlice[0:2], sixthSlice[1:3]

	// sixthSlice,seventhSlice, eighthSlice sử dụng cùng một Backing Array
	// Nên khi thay đổi giá trị pử index 1 = 100 thì cả 3 Slice đầu bị thay đổi theo
	// len(seventhSlice) = 2 do được tạo lấy từ 2 phần tử ở 3 index 0, 1
	// cap(seventhSlice) = 4 do việc sử dụng chung Backing Array nên capacity = 4
	fmt.Println("Length of seventhSlice:", len(seventhSlice), ", Capacity of seventhSlice:", cap(seventhSlice))
	eighthSlice = append(eighthSlice, 100)
	// TODO
	fmt.Println("Length of eighthSlice:", len(eighthSlice), ", Capacity of seventhSlice:", cap(eighthSlice))
	fmt.Println("Length of eighthSlice:", len(sixthSlice[2:]), ", Capacity of seventhSlice:", cap(sixthSlice[2:]))
	sixthSlice[1] = 100
	fmt.Println("The value of sixthSlice:", sixthSlice)
	fmt.Println("The value of seventhSlice:", seventhSlice)
	fmt.Println("The value of eighthSlice:", eighthSlice)

	// Để tránh việc sử dụng chung Backing Array khi tạo Slice từ một Slice khác sử dụng append
	ninthSlice := []int{100}
	ninthSlice = append(ninthSlice, sixthSlice[0:2]...)

	sixthSlice[0] = 999
	fmt.Println("The value of sixthSlice:", sixthSlice)
	fmt.Println("The value of ninthSlice:", ninthSlice)

}