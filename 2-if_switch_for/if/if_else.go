package control

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}

// IfTest 当 if 结构内有 break、continue、goto 或者 return 语句时\
//Go 代码的常见写法是省略 else 部分
func IfTest(x bool) string {
	if x {
		return "输入了true"
	}
	return "输入了false"
}

func StrPlus(orig string) int {
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		os.Exit(1)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	fmt.Printf("The new interger is: %d\n", an)
	return an
}
