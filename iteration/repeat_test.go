package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	repeated := RepeatTimes("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

/** 基准测试命令 go test -bench=.
 Windows Powershell 环境下使用 go test -bench="."
**/
func BenchmarkRepeated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

/*
doc
*/
func ExampleRepeated() {
	RepeatTimes("a", 10)
	fmt.Println("aaaaaaaaaa")
	// Output: aaaaaaaaaa
}

const repeatTimes = 5

func Repeat(msg string) string {
	p := &msg
	sum := ""
	for i := 0; i < repeatTimes; i++ {
		sum += *p
	}
	// fmt.Println("sum", sum, "msg", msg)
	return sum
}

func RepeatTimes(msg string, times int) string {
	p := &msg
	sum := ""
	for i := 0; i < times; i++ {
		sum += *p
	}
	// fmt.Println("sum", sum, "msg", msg)
	return sum
}
