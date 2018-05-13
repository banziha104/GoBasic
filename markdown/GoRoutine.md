# 고루틴 

> Go 프로그램 안에서 동시에 독립적으로 실행되는 흐름의 단위로, 스레드와 비슷한 개념, 다만 수 킬로바이트 정도의 아주 적은 리소스에서 동작함

```go
go f(x,y) // 고루틴으로 실행
```

```go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수시작")

	go long()
	go short()
	time.Sleep(5 * time.Second)
	fmt.Println("main 끝")
}

func short(){
	fmt.Println("short 시작")
	time.Sleep(2 * time.Second)
	fmt.Println("short 끝")
}

func long(){
	fmt.Println("long 시작")
	time.Sleep(3 * time.Second)
	fmt.Println("long 끝")
}
```

## 채널

> 고루틴끼리 정보를 교환하고 실행의 흐름을 동기화하기 위해 사용한다.

- go run -race mysrc.go // 교착상태를 확인
- 채널도 값에 의한 호출방식으로 전달함
- 채널은 기본적으로 양방향 통신이 가능한 상태, 단방향으로 사용할떈 별도로 지정
- 버퍼드 채널 : 일정한 크기의 버퍼를 가질 수 있음
- close : 채널을 닫음
- range : 채널이 닫힌상태인지 아닌지에 따라 전송함
- select : 하나의 고루틴이 여러채널과 통신할때 사용


```go
var ch chan string // string형 채널 생성
	done := make(chan bool) // bool형 채널 생성
	ch <- "msg" // 채널에 msg 전송
	m := <- ch // 채널로부터 메세지 수신
	chan<- string // 송신전용
	<-chan string // 수신전용
	make(chan int,2) // 버퍼드 채널
	v, ok := <-ch
	
	for i := range c{ // 버퍼가 열려있는 상태동안 실행
		
	}
	
	select {// 여러 고루틴과 소통할때 사용
    		case <-c :
    		     // 실행 가능한경우
    		case <- quit :
    			// quit가실행가능한경우
    			// 채널에 값이 전달되었을때
    	}
```

### 저수준 제어 

- sync.Mutex : 여러 고루틴이 공유하는 데이터를 보호할때 사용

```go
func (c *counter) increment(){
	c.mu.Lock() // i 값을 변경하는 부분을 뮤텍스 잠금
	c.i += 1
	c.mu.Unlock() // 뮤텍스해제
}
```

- sync.RWMutex : 읽기 동작과 쓰기 동작을 나누어 잠금처리 가능
- sync.Once : 특정 함수를 한 번만 실행해야할 경우 사용함

```go
func (c *counter) increment(){
	c.once.Do(func(){
		c.i += 1 
	})
}
```

- sync.WaitGroup : 모든 고루틴이 종료될 떄까지 대기후 사용
    - Add() : 고루틴 개수를 추가
    - Done() : 고루틴의 수행이 종료되는 것을 알려줌
    - Wait() : 모든 고루틴이 종료될 때까지 대기
    
- 원자성을 보장하는 연산 : 쪼갤 수 없는 연산을 말함 (atomic 클래스)
    - Add : 특정 포인터 변수에 값을 더함
    - CompareAndSwap : 주언진 값과 비교하여 같으면 새로운 값으로 대체
    - Load : 특정 포인터 변수의 값을 가져옴
    - 특정 포인터 변수에 값을 저장함
    - 특정 포인터 변수에 새로운 값을 저장하고 이전 값을 가져옮
   