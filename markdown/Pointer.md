# 포인터 

- 슬라이스, 맵, 채널, 함수 , 메서드는 참조타입
- 메모리 주소에 직접 접근할 수 있게 하지만, 버그를 유발하기 쉬운 주소 값 연산은 허락하지 않음
- new 함수를 이용하면, 매개변수로 전달한 타입에 맞는 메모리 공간을 찾아 초기화하고 그 주소를 반환

```go
func main() {
 var p *int
 i := 1
 p = &i
 p := new(int) // new 함수를 이용하면, 매개변수로 전달한 타입에 맞는 메모리 공간을 찾아 초기화하고 그 주소를 반환
}
```

```go
package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5,6}
	multiply(numbers,5)
	fmt.Println(numbers)
}

func multiply(numbers []int, factor int){
	for i := range numbers {
		numbers[i] *= factor
	}
}
```