package main

import (
	"flag"
	"fmt"
	"regexp"
	"time"
	"github.com/golang/glog"
	"github.com/google/goexpect"
)

const (
	timeout = 10 * time.Millisecond
)

func main() {
	flag.Parse()
	e, _, err := expect.Spawn("telnet localhost 2605", -1)
	if err != nil {
		glog.Exit(err)
	}
	defer func() {
		fmt.Printf("close\n")
		e.Close()
	}()

	// str := "Trying 127.0.0.1...\n"
	// str += "Connected to localhost.\n"
	// str += "Escape character is '^]'.\n"
	// str += "\n"
	// str += "Hello, this is Quagga (version 0.99.24.1).\n"
	// str += "Copyright 1996-2005 Kunihiro Ishiguro, et al.\n"
	// str += "\n"
	// str += "\n"
	// str += "User Access Verification\n"
	// str += "\n"
	str := ".*Password: "
	passRE   := regexp.MustCompile(str)
	promptRE := regexp.MustCompile("router> ")

	fmt.Printf("start\n");
	res,_,_ := e.Expect(passRE, timeout)
	fmt.Printf("wait1 res=[%s]\n", res);
	fmt.Printf("--------------\n")
	e.Send("zebra\n")
	e.Expect(promptRE, timeout)
	fmt.Printf("wait2\n");
	e.Send("show memory\n")
	a,_,_ := e.Expect(promptRE, timeout)
	fmt.Println("++++++++++++++++++++++")
	fmt.Println(a)
	fmt.Println("++++++++++++++++++++++")

	// cmd  := flag.String("show memory", "", "command to run")
	// pass := flag.String("zebra", "", "password to use")

	// e.Expect(passRE, timeout)
	// e.Send(*pass + "\n")
	// e.Expect(promptRE, timeout)
	// result,_,_ := e.Expect(promptRE, timeout)
	// e.Send("show memory\n")
	// e.Send("exit\n")

}


