package main

import "fmt"

// 注：KiB与KB是不同单位。
// 其中，KiB 是 kilo binary byte的缩写，指的是千位二进制字节
// 而 KB 是 kilobyte 的缩写，指的是千字节
// 两者的区别在于KB与MB的以1000为倍数
const (
	KB = 1000
	MB = 1000 * 1000
	GB = 1000 * 1000 * 1000
	TB = 1000 * 1000 * 1000 * 1000
	PB = 1000 * 1000 * 1000 * 1000 * 1000
	EB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000
	ZB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000
	YB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000
)

// 1024的倍数可以通过移位来计算
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main() {
	var f float64 = 36
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	fmt.Printf("%T\n", (f-32)*5/9)
	fmt.Printf("%T\n", 5/9*(f-32))
	fmt.Printf("%T\n", 5.0/9.0*(f-32))
}
