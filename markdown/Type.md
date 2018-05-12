# 데이터 타입

- Boolean

```go
b1 := true
b2 := false 
```

- 숫자
    - int : 음과 양의 정수 (8,16,32,64)
    - uint : 양의 정수 (8,16,32,64)
    - byte : uint8의 별칭
    - rune : int32의 별칭
    - uintptr : uint와 같음
    
    
```go
func main() {
  a := 365 //10진수
  b := 0555 // 0을 앞에 붙이면 8진수
  c := 0x16D // ox를 앞에 붙이면 16진수
}
```

- 실수 
    - float32 : 소수점 7자리까지 가능
    - float64 : 소수점 15자리까지 가능
    
- 복소수
    - complex64 : 32비트의 실수부와 허수부
    - complex128 : 64비트의 실수부와 허수부 
    
- 연산 : 같은 타입끼리만 가능함
