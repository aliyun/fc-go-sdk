package fc

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestServiceStructs(t *testing.T) {
	suite.Run(t, new(ServiceStructsTestSuite))
}

type ServiceStructsTestSuite struct {
	suite.Suite
}

func (s *ServiceStructsTestSuite) TestCreateService() {
	assert := s.Require()

	tracingConfigInput := NewTracingConfig()
	tracingConfigInput.WithType(TracingTypeJaeger)
	assert.Equal(TracingTypeJaeger, *tracingConfigInput.Type)
	tracingConfigInput.WithParams(NewJaegerConfig().WithEndpoint("mock-jaeger-endpoint"))
	jaegerConfig, ok := tracingConfigInput.Params.(*JaegerConfig)
	assert.True(ok)
	assert.Equal("mock-jaeger-endpoint", *jaegerConfig.Endpoint)

	tracingConfigInput = NewTracingConfig()
	tracingConfigInput.WithJaegerConfig(NewJaegerConfig().WithEndpoint("mock-jaeger-endpoint"))
	jaegerConfig, ok = tracingConfigInput.Params.(*JaegerConfig)
	assert.True(ok)
	assert.Equal("mock-jaeger-endpoint", *jaegerConfig.Endpoint)
	assert.Equal(TracingTypeJaeger, *tracingConfigInput.Type)

	input := NewCreateServiceInput()
	assert.NotNil(input)

	input.WithServiceName("mock-service")
	assert.NotNil(input.ServiceName)
	assert.Equal("mock-service", *input.ServiceName)

	input.WithTracingConfig(tracingConfigInput)
	assert.NotNil(input.TracingConfig)
	assert.Equal(tracingConfigInput, input.TracingConfig)
}

func (s *ServiceStructsTestSuite) TestUpdateService() {
	assert := s.Require()

	tracingConfigInput := NewTracingConfig()
	tracingConfigInput.WithJaegerConfig(NewJaegerConfig().WithEndpoint("mock-jaeger-endpoint"))
	jaegerConfig, ok := tracingConfigInput.Params.(*JaegerConfig)
	assert.True(ok)
	assert.Equal("mock-jaeger-endpoint", *jaegerConfig.Endpoint)
	assert.Equal(TracingTypeJaeger, *tracingConfigInput.Type)

	input := NewUpdateServiceInput("mock-service-name")
	assert.NotNil(input)
	input.WithTracingConfig(tracingConfigInput)
	assert.NotNil(input.TracingConfig)
	assert.Equal(tracingConfigInput, input.TracingConfig)
}
