# 기본 라이브러리 

### 문자열

- bytes : strings 패키지가 제공하는 기능과 유사한 작업가능
- strconv : 문자열을 다른 기본타입으로 변환하거나 다른타입을 문자열로 변환
- fmt : 문자열을 출력하거나 읽는 작업
- unicode :  유니코드
- text/template, html/template :  가변 데이터를 기반으로 텍스트 아웃풋을 만들수 있음

### 컬렉션

- container/heap : 힙
- container/list : 리스트
- container/ring : 링크드 리스크 
- database/sql : SQL 데이터 베이스를 사용하기 위한 기능을 제공함, 실제 데이터베이스와 연동해서 사용하려면 해당 데이터베이스 드라이버를 내려받아야함

### 기본 라이브러리

- os : 운영체제
- math : 수학 연산
- net : 네트워크
- net/http : HTTP 요청을 받아 처리할 수 있는 서버 기능을 제공
- net/url : url 정보를 파싱할 수 있고, 쿼리 문자열을 만들 수 있음
- net/rpc : RPC 서버를 만들 수 있다 
- net/smtp : 메일을 보내는 기능을 제공

### 리플렉션

> 프로그램의 메타 데이터를 활용하여 객체를 동적으로 제어할 수 있다

- reflect.TypeOf() : 타입 확인
- reflect.ValueOf() : 실제값 확인
- reflect.Value : 실제 값이 변경할 수 있는 값이라면 그 값을 동적으로 변경가능
- reflect.Value.CanSet() : 변경 가능한 값인지 확인
- 함수/ 메서드 동적호출

```go
package main

import (
	"strings"
	"fmt"
	"reflect"
)

func TitleCase(s string) string{
	return strings.Title(s)
}

func main() {
	caption := "go is an open source project"
	title := TitleCase(caption)
	fmt.Println(title)

	titleFuncValue := reflect.ValueOf(TitleCase)
	values := titleFuncValue.Call([]reflect.Value{reflect.ValueOf(caption)})
	title = values[0].String()
	fmt.Println(title)
}
```


