package write

import "io"

// 写一个带有如下函数签名的函数CountingWriter，
// 传入一个io.Writer接口类型，
// 返回一个把原来的Writer封装在里面的新的Writer类型
// 和一个表示新的写入字节数的int64类型指针。

type WrapperWriter struct {
	w io.Writer
	c *int64
}

func (ww WrapperWriter) Write(p []byte) (n int, err error) {
	n, err = ww.w.Write(p)
	if err != nil {
		return n, err
	}
	*ww.c += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	ww := WrapperWriter{w, &c}
	return ww, &c
}
