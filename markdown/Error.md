# 에러처리

> go는 try-catch 문이 없고 , 타입으로 정의 되어있음

- 에러 타입

```go
type error interface{
	Error() string // String을 반환하는 Error()가 있으면 에러로 사용될 수 있음
}
```

- 에러 생성

```go
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s 
}
func main() {
	errors.New() // dpfj todtjd 
}
```

- 사용자 정의 에러타입 

```go
type SqrtError struct {
	time time.Time
	value floalt64
	message string
}
```

### Panic

> 실행중에 에러가 발생하면 Go 런타임은 패틱을 발생시킴

```go
panic("A server Error") // 메세지 출력후 프로그램 종료
```

### Recover 

> 패니킹 작업으로부터 프로그램의 제어권을 다시 얻어 종료 절차를 중시키고 프로그램을 계속이어 나갈 수 있게한다.

- recover()는 반드시 defer 안에서 사용해야한다

```go
package main

import "fmt"

func main(){
	fmt.Println("result : ", divide(1,0))
}

func divide(a,b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	return a / b
}
```

### 에러처리

- 에러 확인 함수 사용

```go
func check(err error){
	if err != nil{
		panic(err)
	}
}
```

- 클로저로 에러처리

```go
package main

import (
	"log"
)

type fType func(int, int) int 

func errorHandler(fn fType) fType{
	return func(a ,b int) int {
		defer func() {
			if err,ok := recover().(error); ok {
				log.Print("run time panic %v", err)
			}
		}()
		return fn(a,b)
	}
} 
```

