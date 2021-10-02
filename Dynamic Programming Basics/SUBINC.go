package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)

func main() {
	defer writer.Flush()

	T := reader.NextInt64()
	for ; T > 0; T-- {
		N := reader.NextInt()
		A := make([]int, N)
		for i := 0; i < N; i++{
			A[i] = reader.NextInt()
		}
		var dp = make([]int, N)
		for i, j := 0, 1; j < N; j++ {
			if A[j] < A[j-1] {
				i = j
			}
			dp[j] = dp[j-1] + j - i + 1
		}

		writer.Printf("%d\n", 1 + dp[N-1])
	}
}

/*********************** I/O ***********************/

type BufferedWriter interface {
	Printf(format string, a ...interface{})
	Flush()
}

type writerImpl struct {
	*bufio.Writer
}

func NewBufferedWriter(writer io.Writer) BufferedWriter {
	return &writerImpl{Writer: bufio.NewWriter(writer)}
}

func (impl *writerImpl) Printf(f string, a ...interface{}) {
	fmt.Fprintf(impl.Writer, f, a...)
}

func (impl *writerImpl) Flush() {
	impl.Writer.Flush()
}

type WordScanner interface {
	NextInt() int
	NextInt64() int64
	NextString() string
}

type wordScannerImpl struct {
	*bufio.Scanner
}

func NewWordScanner(reader io.Reader) WordScanner {
	s := bufio.NewScanner(reader)
	s.Split(bufio.ScanWords)
	// adjust the following size as needed
	// 2 << 17 is good enough to scan strings of len 2*10^5
	// it could go as high as 2 << 26 before cf judge breaks
	size := 2 << 20
	buf := make([]byte, size)
	s.Buffer(buf, size)
	return &wordScannerImpl{Scanner: s}
}

func (impl *wordScannerImpl) NextInt() int {
	impl.Scan()
	bb := impl.Bytes()
	i := 0
	if bb[0] == '-' {
		for _, b := range bb[1:] {
			i *= 10
			i -= int(b - '0')
		}
		return i
	}
	for _, b := range bb {
		i *= 10
		i += int(b - '0')
	}
	return i
}

func (impl *wordScannerImpl) NextInt64() int64 {
	impl.Scan()
	bb := impl.Bytes()
	i := int64(0)
	if bb[0] == '-' {
		for _, b := range bb[1:] {
			i *= 10
			i -= int64(b - '0')
		}
		return i
	}
	for _, b := range bb {
		i *= 10
		i += int64(b - '0')
	}
	return i
}


func (impl *wordScannerImpl) NextString() string {
	impl.Scan()
	return impl.Text()
}