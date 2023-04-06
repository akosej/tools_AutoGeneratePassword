package main

import (
	"flag"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	flag.Parse()
	num := flag.Arg(0)
	fmt.Println("generating " + num + " passwords")
	if num == "" {
		fmt.Println("Format: pwdgenerate amount")
		fmt.Println("pwdgenerate 100")
	} else {
		Run("rm", "./passwords.txt")
		Run("touch", "./passwords.txt")
		// Generate a password that is 64 characters long with 10 digits, 10 symbols,
		// allowing upper and lower case letters, disallowing repeat characters.
		cant, _ := strconv.Atoi(num)
		var init int
		var PASS []interface{}
		var exist int
		for {
			res, _ := password.Generate(8, 4, 0, false, false)
			for _, v := range PASS {
				if v == res {
					exist++
				}
			}
			if exist == 0 {
				PASS = append(PASS, res)
				init++
			}
			if init == cant {
				break
			}
		}
		for _, v := range PASS {
			str := fmt.Sprintf("%v", v)
			_ = AppendStrFile("./passwords.txt", str+"\n")
		}
		fmt.Println("Saved in the file passwords.txt")
	}
}

// ------------------------------------------------------------------------------------------------------------
// Execute a system command
// ------------------------------------------------------------------------------------------------------------
func Run(arg ...string) {
	head := arg[0]
	parts := arg[1:len(arg)]
	run := exec.Command(head, parts...)
	_ = run.Run()
}

// Add line at the end of the file
func AppendStrFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	//defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
