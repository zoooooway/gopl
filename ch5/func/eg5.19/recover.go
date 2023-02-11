package recover

// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
// 与recover重名的话则会无法调用
func recover1() (i int) {
	defer func() {
		p := recover()
		// 需要先进行类型断言
		c, ok := p.(int)
		if ok {
			i = c
		}

	}()
	panic(1)
}
