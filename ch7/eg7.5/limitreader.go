package reader

import (
	"io"
)

// io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：

type Reader struct {
	rd    io.Reader
	limit int
}

func (r Reader) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}

	temp := make([]byte, r.limit)
	n, e := r.rd.Read(temp)
	copy(p, temp)
	if len(p) > r.limit {
		e = io.EOF
	}
	return n, e
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return Reader{r, int(n)}
}
