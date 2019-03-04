// 统计字符个数
// author: baoqiang
// time: 2019/3/4 下午6:53
package ch04

import (
	"unicode/utf8"
	"os"
	"fmt"
	"unicode"
	"io"
	"bufio"
)

//	3
//	2
//	a
//	世
//	世界
//	3
//	世
//
//	rune	count
//	'3'	2
//	'\n'	7
//	'2'	1
//	'a'	1
//	'世'	3
//	'界'	1
//
//	len	count
//	1	11
//	2	0
//	3	4
//	4	0

func CountChar() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		// 不合法的字符
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF‐8 characters\n", invalid)
	}
}
