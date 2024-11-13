Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Будет deadlock.

for-range с будет читать с канала до тех пок, пока канал не будет закрыт. Поэтому будет дедлок.
```