# 기본

- 선언법

```go
var arr [5]int // 배열
var p *int // 정수형 포인터
var x int = 10 // 인트형 자료형을 선언
func(int,int) // 두정수를 받음
func(int) int // 정수를 받고 정수를 리턴함

func(int, func(int,int)) func(int) int // 함수를 인자로 받고 반환함

```

- 타입 추론

```go

var i = 10 // 타입추론으로 int형으로 들어감
var p = &i // 타입추론으로 i의 주소가 들어감

/*축약식*/

i := 10
p := &i

```

- 기본
    - if문은 소괄호로 둘러싸지 않음
    - for문 또한 소괄호로 둘러싸지않음
    

```go
// 팩토리얼 순환 예제
package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}
func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n-1)
}

```

```go
// 팩토리얼 반복예제
package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}
func fac(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n --
	}
	return result
}

```

```go
// for문의 다른 사용법

package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}
func fac(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}


```