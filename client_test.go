package fc

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

//noinspection SpellCheckingInspection
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

var endPoint = os.Getenv("ENDPOINT")
var accessKeyId = os.Getenv("ACCESS_KEY_ID")
var accessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
var codeBucketName = os.Getenv("CODE_BUCKET")
var region = os.Getenv("REGION")
var accountID = os.Getenv("ACCOUNT_ID")
var invocationRole = os.Getenv("INVOCATION_ROLE")
var logProject = os.Getenv("LOG_PROJECT")
var logStore = os.Getenv("LOG_STORE")

type FcClientTestSuite struct {
	suite.Suite
}

func TestFcClient(t *testing.T) {
	suite.Run(t, new(FcClientTestSuite))
}

func (s *FcClientTestSuite) TestService() {
	assert := s.Require()

	serviceName := fmt.Sprintf("go-service-%s", RandStringBytes(8))
	client, err := NewClient(endPoint, "2016-08-15", accessKeyId, accessKeySecret)
	assert.Nil(err)

	// clear
	defer func() {
		listServices, err := client.ListServices(
			NewListServicesInput().
				WithLimit(100).
				WithPrefix("go-service-"),
		)
		assert.Nil(err)
		for _, serviceMetadata := range listServices.Services {
			s.clearService(client, *serviceMetadata.ServiceName)
		}
	}()

	// CreateService
	createServiceOutput, err := client.CreateService(NewCreateServiceInput().
		WithServiceName(serviceName).
		WithDescription("this is a service test for go sdk"))

	assert.Nil(err)
	assert.Equal(*createServiceOutput.ServiceName, serviceName)
	assert.Equal(*createServiceOutput.Description, "this is a service test for go sdk")
	assert.NotNil(*createServiceOutput.CreatedTime)
	assert.NotNil(*createServiceOutput.LastModifiedTime)
	assert.NotNil(*createServiceOutput.LogConfig)
	assert.NotNil(*createServiceOutput.Role)
	assert.NotNil(*createServiceOutput.ServiceID)

	// GetService
	getServiceOutput, err := client.GetService(NewGetServiceInput(serviceName))
	assert.Nil(err)

	assert.Equal(*getServiceOutput.ServiceName, serviceName)
	assert.Equal(*getServiceOutput.Description, "this is a service test for go sdk")

	// UpdateService
	updateServiceInput := NewUpdateServiceInput(serviceName).
		WithDescription("new description")
	updateServiceOutput, err := client.UpdateService(updateServiceInput)
	assert.Nil(err)
	assert.Equal(*updateServiceOutput.Description, "new description")

	// UpdateService with IfMatch
	updateServiceInput2 := NewUpdateServiceInput(serviceName).
		WithDescription("new description2").
		WithIfMatch(updateServiceOutput.Header.Get("ETag"))
	updateServiceOutput2, err := client.UpdateService(updateServiceInput2)
	assert.Nil(err)
	assert.Equal(*updateServiceOutput2.Description, "new description2")

	// UpdateService with wrong IfMatch
	updateServiceInput3 := NewUpdateServiceInput(serviceName).
		WithDescription("new description2").
		WithIfMatch("1234")
	_, errNoMatch := client.UpdateService(updateServiceInput3)
	assert.NotNil(errNoMatch)

	// ListServices
	listServicesOutput, err := client.ListServices(
		NewListServicesInput().
			WithLimit(100).
			WithPrefix("go-service-"),
	)
	assert.Nil(err)
	assert.Equal(len(listServicesOutput.Services), 1)
	assert.Equal(*listServicesOutput.Services[0].ServiceName, serviceName)

	for a := 0; a < 10; a++ {
		listServiceName := fmt.Sprintf("go-service-%s", RandStringBytes(8))
		_, errListService := client.CreateService(
			NewCreateServiceInput().
				WithServiceName(listServiceName).
				WithDescription("this is a service test for go sdk"),
		)
		assert.Nil(errListService)
		listServicesOutput, err := client.ListServices(
			NewListServicesInput().
				WithLimit(100).
				WithPrefix("go-service-"),
		)
		assert.Nil(err)
		assert.Equal(len(listServicesOutput.Services), a+2)
	}

	// DeleteService
	_, errDelService := client.DeleteService(NewDeleteServiceInput(serviceName))
	assert.Nil(errDelService)
}

func (s *FcClientTestSuite) TestFunction() {
	assert := s.Require()
	serviceName := fmt.Sprintf("go-service-%s", RandStringBytes(8))
	client, err := NewClient(endPoint, "2016-08-15", accessKeyId, accessKeySecret)

	assert.Nil(err)

	defer s.clearService(client, serviceName)

	// CreateService
	_, err2 := client.CreateService(
		NewCreateServiceInput().
			WithServiceName(serviceName).
			WithDescription("this is a function test for go sdk"),
	)
	assert.Nil(err2)

	// CreateFunction
	functionName := fmt.Sprintf("go-function-%s", RandStringBytes(8))
	createFunctionInput1 := NewCreateFunctionInput(serviceName).
		WithFunctionName(functionName).
		WithDescription("go sdk test function").
		WithHandler("hello_world.handler").
		WithRuntime("nodejs6").
		WithCode(
			NewCode().
				WithOSSBucketName(codeBucketName).
				WithOSSObjectName("hello_world_nodejs"),
		).
		WithTimeout(5)
	createFunctionOutput, err := client.CreateFunction(createFunctionInput1)
	assert.Nil(err)

	assert.Equal(*createFunctionOutput.FunctionName, functionName)
	assert.Equal(*createFunctionOutput.Description, "go sdk test function")
	assert.Equal(*createFunctionOutput.Runtime, "nodejs6")
	assert.Equal(*createFunctionOutput.Handler, "hello_world.handler")
	assert.NotNil(*createFunctionOutput.CreatedTime)
	assert.NotNil(*createFunctionOutput.LastModifiedTime)
	assert.NotNil(*createFunctionOutput.CodeChecksum)
	assert.NotNil(*createFunctionOutput.CodeSize)
	assert.NotNil(*createFunctionOutput.FunctionID)
	assert.NotNil(*createFunctionOutput.MemorySize)
	assert.NotNil(*createFunctionOutput.Timeout)

	// GetFunction
	getFunctionOutput, err := client.GetFunction(NewGetFunctionInput(serviceName, functionName))
	assert.Nil(err)
	assert.Equal(*getFunctionOutput.FunctionName, functionName)
	assert.Equal(*getFunctionOutput.Description, "go sdk test function")
	assert.Equal(*getFunctionOutput.Runtime, "nodejs6")
	assert.Equal(*getFunctionOutput.Handler, "hello_world.handler")
	assert.Equal(*getFunctionOutput.CreatedTime, *createFunctionOutput.CreatedTime)
	assert.Equal(*getFunctionOutput.LastModifiedTime, *createFunctionOutput.LastModifiedTime)
	assert.Equal(*getFunctionOutput.CodeChecksum, *createFunctionOutput.CodeChecksum)
	assert.Equal(*createFunctionOutput.CodeSize, *createFunctionOutput.CodeSize)
	assert.Equal(*createFunctionOutput.FunctionID, *createFunctionOutput.FunctionID)
	assert.Equal(*createFunctionOutput.MemorySize, *createFunctionOutput.MemorySize)
	assert.Equal(*createFunctionOutput.Timeout, *createFunctionOutput.Timeout)

	functionName2 := fmt.Sprintf("go-function-%s", RandStringBytes(8))
	_, errReCreate := client.CreateFunction(createFunctionInput1.WithFunctionName(functionName2))
	assert.Nil(errReCreate)

	// ListFunctions
	listFunctionsOutput, err := client.ListFunctions(
		NewListFunctionsInput(serviceName).
			WithPrefix("go-function-"),
	)
	assert.Nil(err)
	assert.Equal(len(listFunctionsOutput.Functions), 2)
	assert.True(
		*listFunctionsOutput.Functions[0].FunctionName == functionName ||
			*listFunctionsOutput.Functions[1].FunctionName == functionName,
	)
	assert.True(
		*listFunctionsOutput.Functions[0].FunctionName == functionName2 ||
			*listFunctionsOutput.Functions[1].FunctionName == functionName2,
	)

	// UpdateFunction
	updateFunctionOutput, err := client.
		UpdateFunction(NewUpdateFunctionInput(serviceName, functionName).
			WithDescription("newdesc"))
	assert.Equal(*updateFunctionOutput.Description, "newdesc")

	// InvokeFunction
	invokeInput := NewInvokeFunctionInput(serviceName, functionName).WithLogType("Tail")
	invokeOutput, err := client.InvokeFunction(invokeInput)
	assert.Nil(err)
	logResult, err := invokeOutput.GetLogResult()
	assert.NotNil(logResult)
	assert.NotNil(invokeOutput.GetRequestID())
	assert.Equal(string(invokeOutput.Payload), "hello world")

	invokeInput = NewInvokeFunctionInput(serviceName, functionName).WithLogType("None")
	invokeOutput, err = client.InvokeFunction(invokeInput)
	assert.NotNil(invokeOutput.GetRequestID())
	assert.Equal(string(invokeOutput.Payload), "hello world")

	// TestFunction use local zip file
	functionName = fmt.Sprintf("go-function-%s", RandStringBytes(8))
	createFunctionInput := NewCreateFunctionInput(serviceName).
		WithFunctionName(functionName).
		WithDescription("go sdk test function").
		WithHandler("main.my_handler").
		WithRuntime("python2.7").
		WithCode(NewCode().WithFiles("./testCode/hello_world.zip")).
		WithTimeout(5)
	_, errCreateLocalFile := client.CreateFunction(createFunctionInput)
	assert.Nil(errCreateLocalFile)
	invokeOutput, err = client.InvokeFunction(invokeInput)
	assert.Nil(err)
	assert.NotNil(invokeOutput.GetRequestID())
	assert.Equal(string(invokeOutput.Payload), "hello world")
}

func (s *FcClientTestSuite) TestCustomDomain() {
	assertThat := s.Require()

	domainName := fmt.Sprintf("custom-domain-%s", RandStringBytes(8))
	client, err := NewClient(endPoint, "2016-08-15", accessKeyId, accessKeySecret)
	assertThat.Nil(err)

	// clear
	defer func() {
		listDomains, err := client.ListCustomDomain(
			NewListCustomDomainInput().
				WithLimit(100).
				WithPrefix("custom-domain-"),
		)
		assertThat.Nil(err)
		for _, customDomainMetadata := range listDomains.CustomDomains {
			s.clearCustomDomain(client, *customDomainMetadata.DomainName)
		}
	}()

	pathConfig := NewPathConfig().
		WithPath("/").
		WithServiceName("serviceName").
		WithFunctionName("functionName").
		WithQualifier("")
	routeConfig := &RouteConfig{[]PathConfig{*pathConfig}}

	// CreateCustomDomain
	createCustomDomainOutput, err := client.CreateCustomDomain(
		NewCreateCustomDomainInput().
			WithDomainName(domainName).
			WithProtocol("HTTP").
			WithRouteConfig(*routeConfig),
	)
	assertThat.Nil(err)
	assertThat.Equal(*createCustomDomainOutput.DomainName, domainName)
	assertThat.Equal(*createCustomDomainOutput.Protocol, "HTTP")
	assertEqualRouteConfig(assertThat, *createCustomDomainOutput.RouteConfig, *routeConfig)
	assertThat.NotNil(*createCustomDomainOutput.AccountId)
	assertThat.NotNil(*createCustomDomainOutput.ApiVersion)
	assertThat.NotNil(*createCustomDomainOutput.CreatedTime)
	assertThat.NotNil(*createCustomDomainOutput.LastModifiedTime)

	// GetCustomDomain
	getCustomDomainOutput, err := client.GetCustomDomain(NewGetCustomDomainInput(domainName))
	assertThat.Nil(err)
	assertThat.Equal(*getCustomDomainOutput.DomainName, domainName)

	// UpdateCustomDomain update Protocol
	updateCustomDomainInput := NewUpdateCustomDomainInput(domainName).WithProtocol("HTTPS")
	updateCustomDomainOutput, err := client.UpdateCustomDomain(updateCustomDomainInput)
	assertThat.Nil(err)
	assertThat.Equal(updateCustomDomainOutput.Protocol, "HTTPS")

	pathConfig1 := NewPathConfig().
		WithPath("/login").
		WithServiceName("serviceName1").
		WithFunctionName("functionName1").
		WithQualifier("")
	routeConfig1 := &RouteConfig{[]PathConfig{*pathConfig1}}

	// UpdateCustomDomain update RouteConfig
	updateCustomDomainInput1 := NewUpdateCustomDomainInput(domainName).WithRouteConfig(*routeConfig1)
	updateCustomDomainOutput1, err := client.UpdateCustomDomain(updateCustomDomainInput1)
	assertThat.Nil(err)
	assertEqualRouteConfig(assertThat, *updateCustomDomainOutput1.RouteConfig, *routeConfig1)

	// ListCustomDomain
	listCustomDomainOutput, err := client.ListCustomDomain(
		NewListCustomDomainInput().
			WithLimit(100).
			WithPrefix("custom-domain-"),
	)
	assertThat.Nil(err)
	assertThat.Equal(len(listCustomDomainOutput.CustomDomains), 1)
	assertThat.Equal(listCustomDomainOutput.CustomDomains[0].DomainName, domainName)

	for a := 0; a < 10; a++ {
		listDomainName := fmt.Sprintf("custom-domain-%s", RandStringBytes(8))
		_, errListDomainName := client.CreateCustomDomain(
			NewCreateCustomDomainInput().
				WithDomainName(listDomainName).
				WithProtocol("HTTP").
				WithRouteConfig(*routeConfig),
		)
		assertThat.Nil(errListDomainName)

		listDomainsOutput, err := client.ListCustomDomain(
			NewListCustomDomainInput().
				WithLimit(100).
				WithPrefix("custom-domain-"),
		)
		assertThat.Nil(err)
		assertThat.Equal(len(listDomainsOutput.CustomDomains), a+2)
	}

	// DeleteCustomDomain
	_, errDelCustomDomain := client.DeleteCustomDomain(NewDeleteCustomDomainInput(domainName))
	assertThat.Nil(errDelCustomDomain)
}

func assertEqualRouteConfig(assertThat *require.Assertions, actual RouteConfig, expect RouteConfig) {
	actualLength := len(actual.Routes)
	assertThat.Equal(actualLength, len(expect.Routes))

	for i := 0; i < actualLength; i++ {
		assertThat.Equal(
			actual.Routes[i].ServiceName,
			expect.Routes[i].ServiceName,
		)
		assertThat.Equal(
			actual.Routes[i].FunctionName,
			expect.Routes[i].FunctionName,
		)
		assertThat.Equal(
			actual.Routes[i].Path,
			expect.Routes[i].Path,
		)
		assertThat.Equal(
			actual.Routes[i].Qualifier,
			expect.Routes[i].Qualifier,
		)
	}
}

func (s *FcClientTestSuite) clearCustomDomain(c *Client, domainName string) {
	assert := s.Require()

	_, errDelCustomDomain := c.DeleteCustomDomain(
		NewDeleteCustomDomainInput(domainName),
	)
	assert.Nil(errDelCustomDomain)
}

func (s *FcClientTestSuite) TestTrigger() {
	assert := s.Require()
	serviceName := fmt.Sprintf("go-service-%s", RandStringBytes(8))
	functionName := fmt.Sprintf("go-function-%s", RandStringBytes(8))
	client, err := NewClient(endPoint, "2016-08-15", accessKeyId, accessKeySecret)

	assert.Nil(err)

	defer s.clearService(client, serviceName)

	// CreateService
	_, err2 := client.CreateService(NewCreateServiceInput().
		WithServiceName(serviceName).
		WithDescription("this is a function test for go sdk"))
	assert.Nil(err2)

	// CreateFunction
	createFunctionInput1 := NewCreateFunctionInput(serviceName).
		WithFunctionName(functionName).
		WithDescription("go sdk test function").
		WithHandler("main.my_handler").WithRuntime("python2.7").
		WithCode(NewCode().
			WithOSSBucketName(codeBucketName).
			WithOSSObjectName("hello_world.zip")).
		WithTimeout(5)
	_, errCreate := client.CreateFunction(createFunctionInput1)
	assert.Nil(errCreate)

	functionName2 := fmt.Sprintf("go-function-%s", RandStringBytes(8))
	_, errReCreate := client.CreateFunction(
		createFunctionInput1.
			WithFunctionName(functionName2).
			WithHandler("main.wsgi_echo_handler"),
	)
	assert.Nil(errReCreate)
	s.testOssTrigger(client, serviceName, functionName)
	s.testLogTrigger(client, serviceName, functionName)
	s.testHttpTrigger(client, serviceName, functionName2)
}

func (s *FcClientTestSuite) testOssTrigger(client *Client, serviceName, functionName string) {
	assert := s.Require()
	sourceArn := fmt.Sprintf("acs:oss:%s:%s:%s", region, accountID, codeBucketName)
	prefix := "pre"
	suffix := "suf"
	triggerName := "test-oss-trigger"

	createTriggerInput := NewCreateTriggerInput(serviceName, functionName).
		WithTriggerName(triggerName).
		WithInvocationRole(invocationRole).
		WithTriggerType(TriggerTypeOss).
		WithSourceARN(sourceArn).
		WithTriggerConfig(
			NewOSSTriggerConfig().
				WithEvents([]string{"oss:ObjectCreated:PostObject"}).
				WithFilterKeyPrefix(prefix).
				WithFilterKeySuffix(suffix),
		)

	createTriggerOutput, err := client.CreateTrigger(createTriggerInput)
	assert.Nil(err)
	s.checkTriggerResponse(
		&createTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeOss,
		sourceArn,
		invocationRole,
	)

	getTriggerOutput, err := client.GetTrigger(NewGetTriggerInput(serviceName, functionName, triggerName))
	assert.Nil(err)
	s.checkTriggerResponse(
		&getTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeOss,
		sourceArn,
		invocationRole,
	)

	tempUpdateTriggerInput := NewUpdateTriggerInput(serviceName, functionName, triggerName)
	tempOSSTriggerConfig := NewOSSTriggerConfig().WithEvents([]string{"oss:ObjectCreated:*"})
	tempUpdateTriggerInput = tempUpdateTriggerInput.WithTriggerConfig(tempOSSTriggerConfig)
	updateTriggerOutput, err := client.UpdateTrigger(tempUpdateTriggerInput)
	assert.Nil(err)
	s.checkTriggerResponse(
		&updateTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeOss,
		sourceArn,
		invocationRole,
	)
	assert.Equal([]string{"oss:ObjectCreated:*"}, updateTriggerOutput.TriggerConfig.(*OSSTriggerConfig).Events)

	listTriggersOutput, err := client.ListTriggers(NewListTriggersInput(serviceName, functionName))
	assert.Nil(err)
	assert.Equal(len(listTriggersOutput.Triggers), 1)
	_, errReCreate := client.CreateTrigger(
		createTriggerInput.WithTriggerName(triggerName + "-new").
			WithTriggerConfig(
				NewOSSTriggerConfig().
					WithEvents([]string{"oss:ObjectCreated:PostObject"}).
					WithFilterKeyPrefix(prefix + "-new").
					WithFilterKeySuffix(suffix + "-new"),
			),
	)
	assert.Nil(errReCreate)
	listTriggersOutput2, err := client.ListTriggers(NewListTriggersInput(serviceName, functionName))
	assert.Nil(err)
	assert.Equal(len(listTriggersOutput2.Triggers), 2)

	_, errDelTrigger := client.DeleteTrigger(NewDeleteTriggerInput(serviceName, functionName, triggerName))
	assert.Nil(errDelTrigger)

	_, errDelTrigger2 := client.DeleteTrigger(NewDeleteTriggerInput(serviceName, functionName, triggerName+"-new"))
	assert.Nil(errDelTrigger2)
}

func (s *FcClientTestSuite) testLogTrigger(client *Client, serviceName, functionName string) {
	assert := s.Require()
	sourceArn := fmt.Sprintf("acs:log:%s:%s:project/%s", region, accountID, logProject)
	triggerName := "test-log-trigger"

	logTriggerConfig := NewLogTriggerConfig().
		WithSourceConfig(
			NewSourceConfig().WithLogstore(logStore + "_source"),
		).
		WithJobConfig(
			NewJobConfig().
				WithMaxRetryTime(10).
				WithTriggerInterval(60),
		).
		WithFunctionParameter(map[string]interface{}{}).
		WithLogConfig(
			NewJobLogConfig().
				WithProject(logProject).
				WithLogstore(logStore),
		).
		WithEnable(false)

	createTriggerInput := NewCreateTriggerInput(serviceName, functionName).
		WithTriggerName(triggerName).
		WithInvocationRole(invocationRole).
		WithTriggerType(TriggerTypeLog).
		WithSourceARN(sourceArn).
		WithTriggerConfig(logTriggerConfig)

	createTriggerOutput, err := client.CreateTrigger(createTriggerInput)
	assert.Nil(err)
	s.checkTriggerResponse(
		&createTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeLog,
		sourceArn,
		invocationRole,
	)

	getTriggerOutput, err := client.GetTrigger(NewGetTriggerInput(serviceName, functionName, triggerName))
	assert.Nil(err)
	s.checkTriggerResponse(
		&getTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeLog,
		sourceArn,
		invocationRole,
	)

	updateTriggerOutput, err := client.UpdateTrigger(
		NewUpdateTriggerInput(serviceName, functionName, triggerName).
			WithTriggerConfig(logTriggerConfig.WithEnable(true)),
	)
	assert.Nil(err)
	s.checkTriggerResponse(
		&updateTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeLog,
		sourceArn,
		invocationRole,
	)
	assert.Equal(true, *updateTriggerOutput.TriggerConfig.(*LogTriggerConfig).Enable)

	listTriggersOutput, err := client.ListTriggers(NewListTriggersInput(serviceName, functionName))
	assert.Nil(err)
	assert.Equal(len(listTriggersOutput.Triggers), 1)

	_, errDelTrigger := client.DeleteTrigger(NewDeleteTriggerInput(serviceName, functionName, triggerName))
	assert.Nil(errDelTrigger)
}

func (s *FcClientTestSuite) testHttpTrigger(client *Client, serviceName, functionName string) {
	assert := s.Require()
	sourceArn := "dummy_arn"
	invocationRole := ""
	triggerName := "test-http-trigger"

	createTriggerInput := NewCreateTriggerInput(serviceName, functionName).
		WithTriggerName(triggerName).
		WithInvocationRole(invocationRole).
		WithTriggerType(TriggerTypeHttp).
		WithSourceARN(sourceArn).
		WithTriggerConfig(
			NewHTTPTriggerConfig().
				WithAuthType("function").
				WithMethods("GET", "POST"),
		)

	createTriggerOutput, err := client.CreateTrigger(createTriggerInput)
	assert.Nil(err)
	s.checkTriggerResponse(
		&createTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeHttp,
		sourceArn,
		invocationRole,
	)

	getTriggerOutput, err := client.GetTrigger(NewGetTriggerInput(serviceName, functionName, triggerName))
	assert.Nil(err)
	s.checkTriggerResponse(
		&getTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeHttp,
		sourceArn,
		invocationRole,
	)

	updateTriggerOutput, err := client.UpdateTrigger(
		NewUpdateTriggerInput(serviceName, functionName, triggerName).WithTriggerConfig(
			NewHTTPTriggerConfig().
				WithAuthType("anonymous").
				WithMethods("GET", "POST"),
		),
	)
	assert.Nil(err)
	s.checkTriggerResponse(
		&updateTriggerOutput.triggerMetadata,
		triggerName,
		TriggerTypeHttp,
		sourceArn,
		invocationRole,
	)
	assert.Equal("anonymous", *updateTriggerOutput.TriggerConfig.(*HTTPTriggerConfig).AuthType)

	listTriggersOutput, err := client.ListTriggers(NewListTriggersInput(serviceName, functionName))
	assert.Nil(err)
	assert.Equal(len(listTriggersOutput.Triggers), 1)

	_, errDelTrigger := client.DeleteTrigger(NewDeleteTriggerInput(serviceName, functionName, triggerName))
	assert.Nil(errDelTrigger)
}

func (s *FcClientTestSuite) checkTriggerResponse(
	triggerResp *triggerMetadata,
	triggerName,
	triggerType,
	sourceArn,
	invocationRole string,
) {
	assert := s.Require()
	assert.Equal(*triggerResp.TriggerName, triggerName)
	assert.Equal(*triggerResp.TriggerType, triggerType)
	if triggerType != TriggerTypeHttp {
		assert.Equal(*triggerResp.SourceARN, sourceArn)
	} else {
		assert.Nil(triggerResp.SourceARN)
	}
	assert.Equal(*triggerResp.InvocationRole, invocationRole)
	assert.NotNil(*triggerResp.CreatedTime)
	assert.NotNil(*triggerResp.LastModifiedTime)
}

func (s *FcClientTestSuite) clearService(client *Client, serviceName string) {
	assert := s.Require()
	// DeleteFunction
	listFunctionsOutput, err := client.ListFunctions(
		NewListFunctionsInput(serviceName).WithLimit(10),
	)
	assert.Nil(err)
	for _, fuc := range listFunctionsOutput.Functions {
		functionName := *fuc.FunctionName
		listTriggersOutput, err := client.ListTriggers(
			NewListTriggersInput(serviceName, functionName),
		)
		assert.Nil(err)
		for _, trigger := range listTriggersOutput.Triggers {
			_, errDelTrigger := client.DeleteTrigger(
				NewDeleteTriggerInput(serviceName, functionName, *trigger.TriggerName),
			)
			assert.Nil(errDelTrigger)
		}

		_, errDelFunc := client.DeleteFunction(
			NewDeleteFunctionInput(serviceName, functionName),
		)
		assert.Nil(errDelFunc)
	}
	// DeleteService
	_, errDelService := client.DeleteService(
		NewDeleteServiceInput(serviceName),
	)
	assert.Nil(errDelService)
}
