package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	fc "github.com/aliyun/fc-go-sdk"
)

const (
	ACCOUNT_ID   = ""
	serviceName  = ""
	functionName = ""
)

func main() {
	client, _ := fc.NewClient(
		os.Getenv("ENDPOINT"), "2016-08-15", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"),
		fc.WithTransport(&http.Transport{MaxIdleConnsPerHost: 100}),
	)
	client.Config.AccountID = ACCOUNT_ID

	instances, err := client.ListInstances(fc.NewListInstancesInput(serviceName, functionName))
	if err != nil {
		fmt.Printf("List instance failed: %s\n", err)
		return
	}
	instanceID := ""
	for _, ins := range instances.Instances {
		instanceID = ins.InstanceID
		fmt.Println(ins.InstanceID, ins.VersionID)
	}

	// Exec command without -it, stdin=false, tty=false
	if instanceID != "" {
		command := []string{"pwd"}
		_, err := client.InstanceExec(
			fc.NewInstanceExecInput(
				serviceName, functionName, instanceID, command,
			).WithStdin(false).
				WithStdout(true).
				WithStderr(true).
				WithTTY(false).
				WithIdleTimeout(120).
				OnStdout(func(data []byte) {
					fmt.Printf("STDOUT: %s\n", data)
				}).OnStderr(func(data []byte) {
				fmt.Printf("STDERR: %s\n", data)
			}),
		)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		time.Sleep(time.Second * 1)
	}

	// Exec command with stdout, stderr in different stream
	if instanceID != "" {
		//command := []string{"pwd"}
		command := []string{"bash", "-c", "echo stdout-text && echo stderr-text 1>&2;"}
		_, err := client.InstanceExec(
			fc.NewInstanceExecInput(
				serviceName, functionName, instanceID, command,
			).WithStdin(false).
				WithStdout(true).
				WithStderr(true).
				WithTTY(false).
				WithIdleTimeout(120).
				OnStdout(func(data []byte) {
					fmt.Printf("STDOUT: %s\n", data)
				}).OnStderr(func(data []byte) {
				fmt.Printf("STDERR: %s\n", data)
			}),
		)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		time.Sleep(time.Second * 1)
	}


	// Exec command with -it : stdin=true, tty=true
	if instanceID != "" {
		command := []string{"/bin/bash"}
		exec, err := client.InstanceExec(
			fc.NewInstanceExecInput(
				serviceName, functionName, instanceID, command,
			).WithStdin(true).
				WithStdout(true).
				WithStderr(false).
				WithTTY(true).
				WithIdleTimeout(120).
				OnStdout(
					func(data []byte) { fmt.Printf("STDOUT: %s\n", data) }).
				OnStderr(
					func(data []byte) { fmt.Printf("STDERR: %s\n", data) }))
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		go func() {
			fmt.Println("error:", <-exec.ErrorChannel())
		}()
		if err := exec.WriteStdin([]byte("ls\r")); err != nil {
			fmt.Println("error", err)
		}
		time.Sleep(time.Second * 1)
		if err := exec.WriteStdin([]byte("ls --color\r")); err != nil {
			fmt.Println("error", err)
		}
		time.Sleep(time.Second * 1)

		if err := exec.WriteStdin([]byte("exit\r")); err != nil {
			fmt.Println("error", err)
		}
		time.Sleep(time.Second * 1)
	}
}