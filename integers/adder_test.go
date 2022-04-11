package integers

import "testing"
import "fmt"

/**
start godoc 
godoc -http=localhost:6060
**/
func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' bug got '%d'", expected, sum)
	}
}

func Add(n int, n2 int) int {
	return n + n2
}

/*
运行测试套件 go test -v
*/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
