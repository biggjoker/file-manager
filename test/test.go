package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main()  {

	cmdInput := bytes.NewBuffer(nil)
	cmdResult := bytes.NewBuffer(nil)

	cmd  := exec.Command("powershell","ls\n")
	cmd.Stdin = cmdInput
	cmd.Stdout = cmdResult
	if err := cmd.Start();err != nil{
		fmt.Println(err)
	}
	if err := cmd.Wait();err != nil {
		fmt.Println(err)
	}
	cmdInput.WriteString("echo 11 \n")
	rt := cmdResult.String()
	fmt.Println(rt)
	time.Sleep(time.Second)
	cmdInput.WriteString("echo 123123 \n")
	rt = cmdResult.String()
	fmt.Println(rt)
}
