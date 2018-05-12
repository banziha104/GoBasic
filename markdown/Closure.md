# 클로져

> 다른 언어의 즉시실행함수와 같음

```go
package main

func main() { 
	func(x,y int) int{
		return x + y
	}(3,4)
}
```