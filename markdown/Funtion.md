# 함수

- 기본적으로 값에의한 호출
- 참조에 의한 호출은 포인터를 이용
- defer : 함수가 종료되기 전까지 특정구문의 실행을 지연시켰다가, 함수가 종료되기 직전에 지연시켰던 구문을 수행함

```go
package main

func main() {
	_, i := myFunc3() // 둘중한개만 필요할때 와일드 카드패턴사용
	
}

func myFunc(){ // void void

}

func myFunc2(s ...string){ // string 가변 인자

}

func myFunc3() (int,int){ // 반환값이 두개인경우 
	return 1,1
}
```

# 함수를 매개변수로 전달하기

```go
func callback(y int, f func(int,int)){
	f(y,2)
}
```


