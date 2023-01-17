package main

func main() {
	// A子协程 发送0-9数字
	// B子协程 计算输入数字的平方
	// 主协程输出最后的平方数

}

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	for i := range dest {
		println(i)
	}
}
