package fc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	serviceName  = "sn"
	qualifier    = "q"
	functionName = "fn"
	target       = int64(10)
)

func TestProvisionConfigStruct(t *testing.T) {
	suite.Run(t, new(ProvisionConfigStructsTestSuite))
}

type ProvisionConfigStructsTestSuite struct {
	suite.Suite
}

func (s *ProvisionConfigStructsTestSuite) TestPutProvisionConfig_Success() {
	assert := s.Require()

	input := NewPutProvisionConfigInput(serviceName, qualifier, functionName)
	assert.NotNil(input)
	assert.Equal(serviceName, *input.ServiceName)
	assert.Equal(qualifier, *input.Qualifier)
	assert.Equal(functionName, *input.FunctionName)
	assert.NotNil(input.PutProvisionConfigObject)
	assert.Nil(input.Target)
	assert.Nil(input.IfMatch)
	assert.Equal(fmt.Sprintf(provisionConfigWithQualifierPath, pathEscape(serviceName), pathEscape(qualifier), pathEscape(functionName)), input.GetPath())
	assert.Equal(input.PutProvisionConfigObject, input.GetPayload())

	input.WithTarget(target)
	assert.NotNil(input.PutProvisionConfigObject)
	assert.NotNil(input.Target)
	assert.Equal(target, *input.Target)
	assert.Equal(input.PutProvisionConfigObject, input.GetPayload())

	input.WithIfMatch(ifMatch)
	assert.NotNil(input.IfMatch)
	assert.Equal(ifMatch, *input.IfMatch)

	err := input.Validate()
	assert.Nil(err)
}

func (s *ProvisionConfigStructsTestSuite) TestPutProvisionConfig_Fail() {
	assert := s.Require()

	input := NewPutProvisionConfigInput("", qualifier, functionName)
	err := input.Validate()
	assert.NotNil(err)
	assert.Equal("Service name is required but not provided", err.Error())

	input = NewPutProvisionConfigInput(serviceName, "", functionName)
	err = input.Validate()
	assert.NotNil(err)
	assert.Equal("Qualifier is required but not provided", err.Error())

	input = NewPutProvisionConfigInput(serviceName, qualifier, "")
	err = input.Validate()
	assert.NotNil(err)
	assert.Equal("Function name is required but not provided", err.Error())
}

func (s *ProvisionConfigStructsTestSuite) TestGetProvisionConfig_Success() {
	assert := s.Require()

	input := NewGetProvisionConfigInput(serviceName, qualifier, functionName)
	assert.NotNil(input)
	assert.Equal(serviceName, *input.ServiceName)
	assert.Equal(qualifier, *input.Qualifier)
	assert.Equal(functionName, *input.FunctionName)
	assert.Equal(fmt.Sprintf(provisionConfigWithQualifierPath, pathEscape(serviceName), pathEscape(qualifier), pathEscape(functionName)), input.GetPath())

	err := input.Validate()
	assert.Nil(err)
}

func (s *ProvisionConfigStructsTestSuite) TestGetProvisionConfig_Fail() {
	assert := s.Require()

	input := NewGetProvisionConfigInput("", qualifier, functionName)
	err := input.Validate()
	assert.NotNil(err)
	assert.Equal("Service name is required but not provided", err.Error())

	input = NewGetProvisionConfigInput(serviceName, "", functionName)
	err = input.Validate()
	assert.NotNil(err)
	assert.Equal("Qualifier is required but not provided", err.Error())

	input = NewGetProvisionConfigInput(serviceName, qualifier, "")
	err = input.Validate()
	assert.NotNil(err)
	assert.Equal("Function name is required but not provided", err.Error())
}

func (s *ProvisionConfigStructsTestSuite) TestListProvisionConfigs_Success() {
	assert := s.Require()

	input := NewListProvisionConfigsInput()
	assert.NotNil(input)
	assert.Equal(provisionConfigPath, input.GetPath())

	input.WithServiceName(serviceName)
	assert.Equal(serviceName, *input.ServiceName)
	err := input.Validate()
	assert.Nil(err)

	input.WithQualifier(qualifier)
	assert.Equal(qualifier, *input.Qualifier)
	err = input.Validate()
	assert.Nil(err)

	input.WithNextToken("does-not-matter")
	assert.Equal("does-not-matter", *input.NextToken)

	input.WithLimit(int32(2))
	assert.Equal(int32(2), *input.Limit)

	params := input.GetQueryParams()
	assert.Equal(serviceName, params.Get("serviceName"))
	assert.Equal(qualifier, params.Get("qualifier"))
	assert.Equal("does-not-matter", params.Get("nextToken"))
	assert.Equal("2", params.Get("limit"))
}

func (s *ProvisionConfigStructsTestSuite) TestListProvisionConfigs_Fail() {
	assert := s.Require()

	input := NewListProvisionConfigsInput()
	input.WithQualifier(qualifier)
	assert.Equal(qualifier, *input.Qualifier)
	err := input.Validate()
	assert.NotNil(err)
	assert.Equal("Service name is required if qualifier is provided", err.Error())
}
