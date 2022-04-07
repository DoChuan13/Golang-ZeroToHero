#Clean architecture

### Sơ đồ tổng quát
![alt text](clean-arch.png)

### Một kiến trúc Clean Architecture bao gồm 4 layer:

- Models Layer: Là nơi chứa tất cả các models/entities sử dụng trong toàn bộ source code
```golang
    type Student struct {
	    id   int,
	    name string,
	    age  string,
    }   
```

