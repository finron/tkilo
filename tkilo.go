package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	readFromStdin()
}
func aMain() {
	fmt.Println("==start==")
	buf := make([]byte, 1)
	for b, err := os.Stdin.Read(buf); err == nil && b == 1 && buf[0] != 'q'; b, err = os.Stdin.Read(buf) {
		fmt.Printf("Resp:=%s \n", buf) // Use Ctrl - D to reach the end of file // Enter is anther right input char = " ""
	}
	os.Exit(0) // echo $?
}
func readFromStdin() {
	fmt.Println("==start==")
	buf := make([]byte, utf8.UTFMax)
	var prevBuf []byte
	var cnt int
	var idx int
	for b, err := os.Stdin.Read(buf); err == nil && b != 0 && buf[0] != 'q'; b, err = os.Stdin.Read(buf) {
		if idx == 0 {
			prevBuf = buf[:]
		} else {

			prevBuf = append(prevBuf, buf[:]...)

		}
		fmt.Printf("Resp s :=%s \n", prevBuf)
		buf = make([]byte, utf8.UTFMax)
		fmt.Println("idx =", idx)
		idx++
		// use  晚上吃什么  //test error
		for {
			rs, acnt := utf8.DecodeRune(prevBuf)
			// _, acnt := utf8.DecodeRune(prevBuf)
			// if rs == utf8.RuneError || !strconv.IsPrint(rs) {
			if rs == utf8.RuneError || rs == 0 {
				break
			}
			cnt = acnt
			fmt.Println("cnt =", cnt)
			// curVal := prevBuf[:cnt]
			prevBuf = prevBuf[cnt:][:]
			// fmt.Printf("Resp:=%s \n", buf) // Use Ctrl - D to reach the end of file // Enter is anther right input char = " ""
			fmt.Printf("Resp rune source :=%v \n", rs) //
			fmt.Printf("Resp rune :=%v \n\n", string(rs))
			// fmt.Printf("Resp rune :=%s \n", string(curVal))
		}

	}
	// test with "我是中国人,万里长城万里长" //fixed

}
