Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Будет выведено: error

Это происходит из-за того, что интерфейс error в Go работает с типами, 
которые реализуют метод Error(). В данном случае, функция test возвращает указатель на структуру CustomError, 
что приводит к тому, что интерфейс error будет хранить ссылку на этот тип. Поскольку указатель на CustomError 
не равен nil, то переменная err в функции main будет считаться не пустой, и будет выведено сообщение об ошибке.

```