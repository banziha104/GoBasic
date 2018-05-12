# 구조체 

> 실세계에 엔티티를 표현

```go
type 타입명 struct{
	필드명1 필드타입1
	필드명2 필드타입2
	필드명3,필드명4 필드타입3 // 두개가 같은 타입인경우
}
```

- 구조체생성(리터럴)

```go
package main

type Person struct {
	name string
	age int
}
func main() {
	p := Person{"이영준",27} // 구조체 리터럴생성
	p2 := &Person{"이영준2",28} // 구조체 포인터 생성, 초기값 할당에 유용
	p3 := new(Person) // new 함수로 구조체 포인터 생성, 초기값이 없는 경우에 유용
	p3.age = 29
	rects := []struct{w, h int}{{1,2},{},{-1,-2}} // 익명 구조체 생성
}
```

- 내부 필드 접근은 . 연산자를 이용
- 태그 : 태그는 필드에 추가된 문자열 , 주석같은 개념
- 구조체 임베딩  : 상속과 비슷한의미, 다른 구조체를 구조체안에 넘음으로써 가능
- 임베디드 필드 : 내부 필드와 이름이 같을때는 필드의 타입을 적어줌

```go
package main

type Person struct {
	name string
	age int
}

type People struct {
	Person
}
func main() {
	p := Person{"이영준",27} // 구조체 리터럴생성
	p2 := &Person{"이영준2",28} // 구조체 포인터 생성, 초기값 할당에 유용
	p3 := new(Person) // new 함수로 구조체 포인터 생성, 초기값이 없는 경우에 유용
	p3.age = 29
	rects := []struct{w, h int}{{1,2},{},{-1,-2}} // 익명 구조체 생성
	people := People{Person{"이영준",28}}
	people.Person.age 
}
```

- 메서드 오버라이딩

```go
package main

import "fmt"

type Item struct {
	name string
	price float64
	quantity int
}

func (t Item) Cost() float64{
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64{
	return t.Item.Cost() * (1.0 - t.discountRate /100)
}
func main() {
	shoes := Item{"Women's Walking Shoes",3000,2}
	eventShoes := DiscountItem{
		Item{"Sports Shoes", 50000,3},
		10.00,
	}
	fmt.Println(shoes.Cost())
	fmt.Println(eventShoes.Cost())
	fmt.Println(eventShoes.Item.Cost())
}
```

- 생성 함수 : 생성자가 없어 별도로 만들어야함

```go
func NewItem(name string, price float64, quantity int) * Item {
	if(price <= 0 || quantity <= 0 || len(name) ==0){
		return nil
	}
	return &Item{name,price, quantity}
}
```

- getter , setter 
    - getter는 GetName()이 아닌 Name()으로 컨벤션
    - setter는 SetName()으로 컨밴션 
    - 대소문자로 접근제어조절