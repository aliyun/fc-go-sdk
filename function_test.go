package fc

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func getStringPointer(str string) *string {
	return &str
}

type FunctionStructsTestSuite struct {
	suite.Suite
}

func (s *FunctionStructsTestSuite) TestHeaders() {
	assert := s.Require()

	input := NewInvokeFunctionInput("service", "func")
	assert.Equal("service", *input.ServiceName)
	assert.Equal("func", *input.FunctionName)

	input.WithAsyncInvocation()
	headers := input.GetHeaders()
	assert.Equal("Async", headers["X-Fc-Invocation-Type"])

	input.WithHeader("X-Fc-Invocation-Code-Version", "Latest")
	headers = input.GetHeaders()
	assert.Equal("Latest", headers["X-Fc-Invocation-Code-Version"])
}

func (s *FunctionStructsTestSuite) TestEnvironmentVariables() {
	assert := s.Require()

	{
		input := NewCreateFunctionInput("service")
		assert.Equal("service", *input.ServiceName)
		assert.Nil(input.EnvironmentVariables)

		input.WithEnvironmentVariables(map[string]string{})
		assert.NotNil(input.EnvironmentVariables)
		assert.Len(input.EnvironmentVariables, 0)

		input.WithEnvironmentVariables(map[string]string{"a": "b"})
		assert.NotNil(input.EnvironmentVariables)
		assert.Equal(map[string]string{"a": "b"}, input.EnvironmentVariables)
	}

	{
		input := NewUpdateFunctionInput("service", "func")
		assert.Equal("service", *input.ServiceName)
		assert.Equal("func", *input.FunctionName)
		assert.Nil(input.EnvironmentVariables)

		input.WithEnvironmentVariables(map[string]string{})
		assert.NotNil(input.EnvironmentVariables)
		assert.Len(input.EnvironmentVariables, 0)

		input.WithEnvironmentVariables(map[string]string{"a": "b"})
		assert.NotNil(input.EnvironmentVariables)
		assert.Equal(map[string]string{"a": "b"}, input.EnvironmentVariables)
	}

	output := &GetFunctionOutput{}
	assert.Nil(output.EnvironmentVariables)
}

func (s *FunctionStructsTestSuite) TestInstanceConcurrency() {
	assert := s.Require()

	{
		input := NewCreateFunctionInput("service")
		assert.Equal("service", *input.ServiceName)
		assert.Nil(input.InstanceConcurrency)

		input.WithInstanceConcurrency(int32(0))
		assert.NotNil(input.InstanceConcurrency)
		assert.Equal(int32(0), *input.InstanceConcurrency)

		input.WithInstanceConcurrency(int32(1))
		assert.NotNil(input.InstanceConcurrency)
		assert.Equal(int32(1), *input.InstanceConcurrency)
	}
	{
		input := NewUpdateFunctionInput("service", "func")
		assert.Equal("service", *input.ServiceName)
		assert.Equal("func", *input.FunctionName)
		assert.Nil(input.InstanceConcurrency)

		input.WithInstanceConcurrency(int32(0))
		assert.NotNil(input.InstanceConcurrency)
		assert.Equal(int32(0), *input.InstanceConcurrency)

		input.WithInstanceConcurrency(int32(1))
		assert.NotNil(input.InstanceConcurrency)
		assert.Equal(int32(1), *input.InstanceConcurrency)
	}

	output := &GetFunctionOutput{}
	assert.Nil(output.InstanceConcurrency)
}

func (s *FunctionStructsTestSuite) TestCustomContainerArgs() {
	assert := s.Require()

	{
		input := NewCreateFunctionInput("service")
		assert.Equal("service", *input.ServiceName)
		assert.Nil(input.CustomContainerConfig)
		assert.Nil(input.CAPort)

		input.WithCustomContainerConfig(&CustomContainerConfig{})
		assert.NotNil(input.CustomContainerConfig)
		assert.Nil(input.CustomContainerConfig.Image)
		assert.Nil(input.CustomContainerConfig.Command)
		assert.Nil(input.CustomContainerConfig.Args)

		port := int32(9000)
		input.WithCAPort(port)
		assert.NotNil(input.CAPort)
		assert.Equal(int32(9000), *input.CAPort)

		input.WithCustomContainerConfig(&CustomContainerConfig{
			Image:   getStringPointer("registry.cn-hangzhou.aliyuncs.com/fc-test/busybox"),
			Command: getStringPointer(`["python", "server.py"]`),
			Args:    getStringPointer(`["9000"]`),
		})
		assert.NotNil(input.CustomContainerConfig)
		assert.Equal("registry.cn-hangzhou.aliyuncs.com/fc-test/busybox", *input.CustomContainerConfig.Image)
		assert.Equal(`["python", "server.py"]`, *input.CustomContainerConfig.Command)
		assert.Equal(`["9000"]`, *input.CustomContainerConfig.Args)
	}

	{
		input := NewUpdateFunctionInput("service", "func")
		assert.Equal("service", *input.ServiceName)
		assert.Equal("func", *input.FunctionName)
		assert.Nil(input.CustomContainerConfig)
		assert.Nil(input.CAPort)

		input.WithCustomContainerConfig(&CustomContainerConfig{})
		assert.NotNil(input.CustomContainerConfig)
		assert.Nil(input.CustomContainerConfig.Image)
		assert.Nil(input.CustomContainerConfig.Command)
		assert.Nil(input.CustomContainerConfig.Args)

		port := int32(9000)
		input.WithCAPort(port)
		assert.NotNil(input.CAPort)
		assert.Equal(int32(9000), *input.CAPort)

		input.WithCustomContainerConfig(&CustomContainerConfig{
			Image:   getStringPointer("registry.cn-hangzhou.aliyuncs.com/fc-test/busybox"),
			Command: getStringPointer(`["python", "server.py"]`),
			Args:    getStringPointer(`["9000"]`),
		})
		assert.NotNil(input.CustomContainerConfig)
		assert.Equal("registry.cn-hangzhou.aliyuncs.com/fc-test/busybox", *input.CustomContainerConfig.Image)
		assert.Equal(`["python", "server.py"]`, *input.CustomContainerConfig.Command)
		assert.Equal(`["9000"]`, *input.CustomContainerConfig.Args)
	}

	output := &GetFunctionOutput{}
	assert.Nil(output.CustomContainerConfig)
	assert.Nil(output.CAPort)
}

func TestFunctionStructs(t *testing.T) {
	suite.Run(t, new(FunctionStructsTestSuite))
}
