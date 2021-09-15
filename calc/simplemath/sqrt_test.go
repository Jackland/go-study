package simplemath

import "testing"

/**
在编写单元测试时，需要引入 testing 包，你可以将其类比为 PHP 中的 PHPUnit 或 Java 中的 JUnit，
我们可以基于该包提供的方法来实现自动化测试，测试方法的格式如下所示：
*/

func TestSqrt(t *testing.T) {
	v := Sqrt(9)
	if v != 3 {
		t.Errorf("Sqrt(9) failed. Got %v, expected 3.", v)
	}
}
