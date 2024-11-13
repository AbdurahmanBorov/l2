package unpacker

import (
	"fmt"
	"testing"
)

// TestResult используется для хранения ожидаемых результатов и ошибок
type TestResult struct {
	result string
	err    error
}

func TestUnpack(t *testing.T) {
	// Тестовые строки
	testString := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}

	// Ожидаемые результаты
	expectedResults := []TestResult{
		{result: "aaaabccddddde", err: nil},
		{result: "abcd", err: nil},
		{result: "", err: fmt.Errorf("Некорректная строка")},
		{result: "", err: nil},
		{result: "qwe45", err: nil},
		{result: "qwe44444", err: nil},
		{result: `qwe\\\\\`, err: nil},
	}

	for i, v := range testString {
		result, err := Unpack(v)

		// Проверка, совпадает ли ошибка с ожидаемой
		if (err != nil) != (expectedResults[i].err != nil) {
			t.Errorf("Ошибка для %q: ожидалась ошибка %v, получено %v", v, expectedResults[i].err, err)
			continue
		}

		// Если ошибки совпадают, сравниваем результат
		if result != expectedResults[i].result {
			t.Errorf("Ошибка для %q: ожидалось %q, получено %q", v, expectedResults[i].result, result)
		}
	}
}
