# 객체지향

> go에는 사용자가 타입을 직접 정의할 수 있는 사용자 정의타입이 존재, 구조체와 인터페이스를 사용자 정의타입으로 사용

- type 타입명 타입명세

```go
type quantitiy int // 기본타입명
type costClauculator func(quantitiy, float32) float64	// 함수로 정의
type rect struct {	// 구조체를 정의
	width  float64
	height float64
}

type shaper interface { //인터페이스
	area() float64 
}
```

- 메서드 : 정의 타입 값에 호출할 수 있는 특별한 함수이다.
    - 기본적으로 값에 의한 호출
    - 참조에 의한호출은 포인터로 사용
    
```go

package main

type quantity int

func (q *quantity) increment() {*q++}

func (q *quantity) decrement() {*q--}

func main() {
	q := quantity(3)
	q.increment()
}


```

- 참조 타입의 경우
```go
package main

import "fmt"

type quantity []int

func (q quantity) increment() {q++}

func (q quantity) decrement() {q--}

func main() {
    q := quantity(3)
    q.increment()
    fmt.Println(q)
}

```

- 리시버 변수 생략

```go
package main

import "fmt"

type rect struct{
	width  float64
	height float64
}

func (rect) new() rect{
	return rect{}
}

func main() {
	r := rect{}.new()
	r.height = 10
	r.width = 30
	fmt.Println(r)
}
```

- 메서드의 함수 표현식 : 메서드도 변수에 할당할 수 있고, 함수의 매개변수로도 사용가능

```go
package main

import "fmt"

type rect struct{
	width  float64
	height float64
}

func (rect) new() rect{
	return rect{}
}

func (r rect) area() float64{
	return r.width * r.height
}

func (r *rect) resize(w,h float64){
	r.width += w 
	r.height += h
}

func main() {
	r := rect{3,4}
	r.resize(10,10)
	
	areaFn := rect.area  // 변수에 메서드를 할당
	resizeFn := (*rect).resize // 변수에 메서드를 할당
	
	fmt.Println(areaFn(r))
	resizeFn(&r,10,10)
}

```

