# 문자열

> Go언어의 소스코드는 UTF-8로 되어있음, 유동적으로 처리됌

- 유니코드 출력

```go
for i, r := range "가나다"{ // i는 int형, r은 rune형 -> int32비트 정수의 별칭
		fmt.Println(i,r)
	}
for _, r := range "가나다"{ // _ 와일카드 패턴
		fmt.Println(r)
	}
```

- 문자열 숫자 간 캐스팅

```go
    var i int
	var err error
	a:="1"
	i,err = strconv.Atoi(a)
	fmt.Println(i,err)
	
```

```go
    var i string
	a:=1
	i = strconv.Itoa(a)
	fmt.Println(i)
```

