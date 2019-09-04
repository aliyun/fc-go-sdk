package main

import (
	"fmt"
	"os"

	"github.com/aliyun/fc-go-sdk"
)

func main() {
	client, _ := fc.NewClient(os.Getenv("ENDPOINT"), "2016-08-15", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRETE"))

	serviceName := "wanghq_async_invoke"
	functionName := "hello_world"
	triggerName := "trigger3"
	triggerName2 := "trigger4"
	qualifier := "LATEST"

	// Create trigger
	fmt.Println("Creating trigger")
	createTriggerInput := fc.NewCreateTriggerInput(serviceName, functionName).WithTriggerName(triggerName).
		WithDescription("create trigger").WithInvocationRole("acs:ram:cn-hangzhou:123:role1").WithTriggerType("oss").
		WithSourceARN("acs:oss:cn-hangzhou:123:fcbucket").WithTriggerConfig(
		fc.NewOSSTriggerConfig().WithEvents([]string{"oss:ObjectCreated:PostObject"}).WithFilterKeyPrefix("r").WithFilterKeySuffix("s"))

	createTriggerOutput, err := client.CreateTrigger(createTriggerInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("CreateTrigger response: %s \n", createTriggerOutput)
	}

	createTriggerInput2 := fc.NewCreateTriggerInput(serviceName, functionName).WithTriggerName(triggerName2).
		WithDescription("create trigger").WithInvocationRole("acs:ram:cn-hangzhou:123:role1").WithTriggerType("oss").
		WithSourceARN("acs:oss:cn-hangzhou:123:fcbucket").WithQualifier(qualifier).WithTriggerConfig(
		fc.NewOSSTriggerConfig().WithEvents([]string{"oss:ObjectCreated:PostObject"}).WithFilterKeyPrefix("r").WithFilterKeySuffix("s"))

	createTriggerOutput2, err := client.CreateTrigger(createTriggerInput2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("CreateTrigger response: %s \n", createTriggerOutput2)
	}

	getTriggerOutput, err := client.GetTrigger(fc.NewGetTriggerInput(serviceName, functionName, triggerName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("GetTrigger response: %s \n", getTriggerOutput)
	}

	getTriggerOutput2, err := client.GetTrigger(fc.NewGetTriggerInput(serviceName, functionName, triggerName2))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("GetTrigger response: %s \n", getTriggerOutput2)
	}

	updateTriggerOutput, err := client.UpdateTrigger(fc.NewUpdateTriggerInput(serviceName, functionName, triggerName).WithDescription("update trigger").WithInvocationRole("acs:ram:cn-hangzhou:123:role2"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateTrigger response: %s \n", updateTriggerOutput)
	}

	updateTriggerOutput2, err := client.UpdateTrigger(fc.NewUpdateTriggerInput(serviceName, functionName, triggerName2).WithDescription("update trigger").WithInvocationRole("acs:ram:cn-hangzhou:123:role2"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateTrigger response: %s \n", updateTriggerOutput2)
	}

	deleteTriggerOutput, err := client.DeleteTrigger(fc.NewDeleteTriggerInput(serviceName, functionName, triggerName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("DeleteTrigger response: %s \n", deleteTriggerOutput)
	}

	deleteTriggerOutput2, err := client.DeleteTrigger(fc.NewDeleteTriggerInput(serviceName, functionName, triggerName2))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("DeleteTrigger response: %s \n", deleteTriggerOutput2)
	}

	listTriggersOutput, err := client.ListTriggers(fc.NewListTriggersInput(serviceName, functionName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("ListTriggers response: %s \n", listTriggersOutput)
	}

}
