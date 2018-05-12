# 패키지

> 다른언어의 모듈이나 라이브러리와 유사한 개념, 패키지 이름과 디렉터리의 이름이 같아야함 

- main : 패키지 이름이 main인 경우, 실행 가능한 프로그램으로 인식
- 라이브러리 : main 이외에는 라이브러리로 만들짐
- 접근제어 : 소문자로 시작(private), 대문자로 시작(public)
- 별칭 

```go
package main
import (
	"fmt"
	mylib "go/ast" // mylib으로 사용가능
)

func main(){
	mylib.ArrayType{}
}
```

- init 함수 : 패키지가 로드될 때 가장먼저 호출되는 함수로, 패키지의 초기화 로직이 필요할때 선택적으로 사용하면됌
