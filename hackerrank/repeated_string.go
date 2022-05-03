// https://www.hackerrank.com/challenges/repeated-string/problem
package main

import (
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	/*
		A = s の中にある`a`の数
		X = n / len(s)
		Y = n mod len(s)
		として、A * X + Yにaが含まれている数を求める
		example0: A=1, X=1000000000000, Y=0 -> 1 * 1000000000000 + 0
		example: A=2, X=2, Y=0 -> 2 * 2 + 0
		example1: A=2, X=3, Y=1 -> 2 * 3 + 1
	*/
	if len(s) == 0 {
		return 0
	}
	a := strings.Count(s, "a")
	x := n / int64(len(s))
	y := n % int64(len(s))
	if y == 0 {
		return int64(a) * x
	}
	if int64(len(s)) < y {
		return int64(a) * x
	}
	substr := s[0:y]
	return int64(a)*x + int64(strings.Count(substr, "a"))
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
*/
