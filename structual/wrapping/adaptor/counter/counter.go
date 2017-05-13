package main

import (
	"strconv"
	"io"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Counter struct {
  Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		 f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}
	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return f.Count(n - 1)
}

func main (){
	pr, pw := io.Pipe()
	defer pw.Close()
	defer pr.Close()

	c := Counter{
		Writer: pw,
	}

	file, _ := os.Create("file.txt")
	tee := io.TeeReader(pr, file)

	go func(){
		io.Copy(os.Stdout, tee)
	}()

	c.Count(3)
}