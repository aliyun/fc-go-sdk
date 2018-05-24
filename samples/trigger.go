package main

import (
	"fmt"
	"os"

	"aliyun/serverless/lambda-go-sdk"
)

func main() {
	client, _ := fc.NewClient(os.Getenv("ENDPOINT"), "2016-08-15", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRETE"))

	serviceName := "wanghq_async_invoke"
	functionName := "hello_world"
	triggerName := "trigger3"

	// Create trigger
	fmt.Println("Creating trigger")
	createTriggerInput := fc.NewCreateTriggerInput(serviceName, functionName).WithTriggerName(triggerName).
		WithInvocationRole("acs:ram:cn-hangzhou:123:role1").WithTriggerType("oss").WithSourceARN("acs:oss:cn-hangzhou:123:fcbucket").
		WithTriggerConfig(
			fc.NewOSSTriggerConfig().WithEvents([]string{"oss:ObjectCreated:PostObject"}).WithFilterKeyPrefix("r").WithFilterKeySuffix("s"))

	createTriggerOutput, err := client.CreateTrigger(createTriggerInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("CreateTrigger response: %s \n", createTriggerOutput)
	}

	getTriggerOutput, err := client.GetTrigger(fc.NewGetTriggerInput(serviceName, functionName, triggerName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("GetTrigger response: %s \n", getTriggerOutput)
	}

	updateTriggerOutput, err := client.UpdateTrigger(fc.NewUpdateTriggerInput(serviceName, functionName, triggerName).WithInvocationRole("acs:ram:cn-hangzhou:123:role2"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateTrigger response: %s \n", updateTriggerOutput)
	}

	deleteTriggerOutput, err := client.DeleteTrigger(fc.NewDeleteTriggerInput(serviceName, functionName, triggerName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("DeleteTrigger response: %s \n", deleteTriggerOutput)
	}

	listTriggersOutput, err := client.ListTriggers(fc.NewListTriggersInput(serviceName, functionName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("ListTriggers response: %s \n", listTriggersOutput)
	}

}
