# 문자열

> Go언어의 소스코드는 UTF-8로 되어있음, 유동적으로 처리됌

- 유니코드 출력

```go
for i, r := range "가나다"{ // i는 int형, r은 rune형 -> int32비트 정수의 별칭
		fmt.Println(i,r)
	}
for _, r := range "가나다"{ // _ 와일카드 패턴
		fmt.Println(r)
	}
```

- 문자열 숫자 간 캐스팅

```go
    var i int
	var err error
	a:="1"
	i,err = strconv.Atoi(a)
	fmt.Println(i,err)
	
```

```go
    var i string
	a:=1
	i = strconv.Itoa(a)
	fmt.Println(i)
```

- 배열과 슬라이스
    - 배열은 미리 정해져있음
    - 슬라이스는 정해져 있지않음.
    ```go
    func main() {
     	fruits := [3]string{"사과","바나나","토마토"} // 배열
     	for _, fruit := range fruits {
     		fmt.Println(fruit,"는 맛 \n",)
     	}
   	    fruits2 := []string{"사과","바나나","토마토"} // 슬라이스
         	for _, fruit := range fruits {
         		fmt.Println(fruit,"는 맛 \n",)
         	}
     }
    ```
    
- 슬라이스 사용

```go
	fmt.Println(fruits[:2]) // 슬라이싱
	fruits = apeend(fruits,"포도","딸기") //슬라이싱 덧붙이기
	cap(fruits) // 용량
	len(fruits) //길이
	
	a := []int{1,2,3,4,5,6,7,8,9,10}
	a = append(a[3:9]) // 3부터 8번째인덱스까지 짜름
	a = a[:len(a)-1] // 뒤에하나를자름
	a[len(a)-1] = nil // 메모리 누수 방지
```

