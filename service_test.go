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

	input.WithVPCConfig(NewVPCConfig().
		WithVPCID("vpc-id").
		WithVSwitchIDs([]string{"sw-1", "sw-2"}).
		WithSecurityGroupID("sg-id"))
	assert.NotNil(input.VPCConfig)
	assert.Equal("vpc-id", *input.VPCConfig.VPCID)
	assert.Equal([]string{"sw-1", "sw-2"}, input.VPCConfig.VSwitchIDs)
	assert.Equal("sg-id", *input.VPCConfig.SecurityGroupID)
	
	logStore := NewLogConfig().WithProject("mock-log-project").WithLogstore("mock-log-store")
	logStore.WithEnableRequestMetrics(true)
	input.WithLogConfig(logStore)
	assert.NotNil(logStore)
	assert.Equal(true, *logStore.EnableRequestMetrics)

	logStore.WithEnableInstanceMetrics(true)
	input.WithLogConfig(logStore)
	assert.NotNil(logStore)
	assert.Equal(true, *logStore.EnableInstanceMetrics)
}

func (s *ServiceStructsTestSuite) TestUpdateService() {
	assert := s.Require()

	input := NewUpdateServiceInput("mock-service-name")
	assert.NotNil(input)

	logConfig := NewLogConfig().WithProject("mock-log-project").WithLogstore("mock-log-store").WithEnableRequestMetrics(true).WithEnableInstanceMetrics(true)
	input.WithLogConfig(logConfig)
	assert.NotNil(logConfig)
	assert.Equal(true, *input.LogConfig.EnableRequestMetrics)
	assert.Equal(true, *input.LogConfig.EnableInstanceMetrics)

}
