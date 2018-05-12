# 인터페이스

> 객체의 동작을 표현하는 것, 덕 타이핑 방식을 이용함, 네이밍 컨벤셔는 e,er을 접두사로 붙임

```go
type 인터페이스명 interface{
	메서드1(매개변수) 반환타입
}
```

```go
package main

import "fmt"

type shaper interface {
	area() float64
}

func describe(s shaper){
	fmt.Println("area : ", s.area())
}

type rect struct {width, height float64}

func (r rect) area() float64{ 
	return r.width * r.height
}

func main(){
	r := rect{3,4}
	describe(r)  // r의 area의 반환값과 시그니쳐 매개변수가 맞음으로 실행됌(덕 타이핑)
}
```

- 익명 인터페이스 : 인터페이스도 타입으로 정의하지 않고 익명으로 사용할 수 있음

```go
func display(s interface {show()}){
	s.show()
}
```

- 빈 인터페이스 : 메서드를 정의하지 않은 인터페이스로, 어떤 값이든 전달 받을수 있음

```go
func main(){
	r := rect{3,4}
	display(r)
}

func display(s interface{}){
	fmt.Println(s)
}
```

- 제너릭 컬렉션

```go
type Items []Coster // 인터페이스 슬라이스로 Cost()메서드를 가진 타입은 모두 담을 수 있음
func (ts Items) Cost() (c float64) { // Items 이외에도 Cost메서드를 추가하면 Coster 인터페이스로 활용가능
	
}

```

- 인터페이스 임베딩 

```go
type Item interface{
	Coster
	fmt.Stringer
}

```

- 타입 변환 : 실제 타입과 관계없이 인터페이스에서 정의한 메서드를 가지기만 하면 어떤 값이라도 동적으로 할당받아서 사용할 수 있다.
    - 어설션 변환
    - switch문으로 변환
    
    
```go
v := interfaceValue.(Type) // 어설션으로 변환
```