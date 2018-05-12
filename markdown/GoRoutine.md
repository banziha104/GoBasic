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
