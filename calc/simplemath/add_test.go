package simplemath

import "testing"

/**
在 Go 语言中，支持为功能模块编写单元测试代码，继续以上篇教程构建的计算器项目为例，
在 simplemath 包中，我们可以为每一个运算模块编写了对应的单元测试代码，单元测试文件以 _test 作为文件名后缀来标识，
比如我们通过 add_test.go 和 sqrt_test.go 文件分别为 add.go 和 sqrt.go 编写单元测试，对应的目录结构如下：


*/

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	//if r != 3 {
	//	t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	//}
	if r != 2 {
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}

}
