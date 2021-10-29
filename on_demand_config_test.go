package fc

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	maximumInstanceCount = int64(10)
)

func TestOnDemandConfigStruct(t *testing.T) {
	suite.Run(t, new(OnDemandConfigStructsTestSuite))
}

type OnDemandConfigStructsTestSuite struct {
	suite.Suite
}

func (s *OnDemandConfigStructsTestSuite) TestPutOnDemandConfig_Success() {
	assert := s.Require()

	input := NewPutOnDemandConfigInput(serviceName, qualifier, functionName)
	assert.NotNil(input)
	assert.Equal(serviceName, *input.ServiceName)
	assert.Equal(qualifier, *input.Qualifier)
	assert.Equal(functionName, *input.FunctionName)
	assert.NotNil(input.PutOnDemandConfigObject)
	assert.Nil(input.MaximumInstanceCount)
	assert.Nil(input.IfMatch)
	assert.Equal(fmt.Sprintf(onDemandConfigWithQualifierPath, pathEscape(serviceName), pathEscape(qualifier), pathEscape(functionName)), input.GetPath())
	assert.Equal(input.PutOnDemandConfigObject, input.GetPayload())

	input.WithMaximumInstanceCount(maximumInstanceCount)
	assert.NotNil(input.PutOnDemandConfigObject)
	assert.NotNil(input.MaximumInstanceCount)
	assert.Equal(maximumInstanceCount, *input.MaximumInstanceCount)
	assert.Equal(input.PutOnDemandConfigObject, input.GetPayload())

	input.WithIfMatch(ifMatch)
	assert.NotNil(input.IfMatch)
	assert.Equal(ifMatch, *input.IfMatch)

	err := input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestPutOnDemandConfig_EmptyParams() {
	assert := s.Require()

	input := NewPutOnDemandConfigInput("", qualifier, functionName)
	err := input.Validate()
	assert.Nil(err)

	input = NewPutOnDemandConfigInput(serviceName, "", functionName)
	err = input.Validate()
	assert.Nil(err)

	input = NewPutOnDemandConfigInput(serviceName, qualifier, "")
	err = input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestGetOnDemandConfig_Success() {
	assert := s.Require()

	input := NewGetOnDemandConfigInput(serviceName, qualifier, functionName)
	assert.NotNil(input)
	assert.Equal(serviceName, *input.ServiceName)
	assert.Equal(qualifier, *input.Qualifier)
	assert.Equal(functionName, *input.FunctionName)
	assert.Equal(fmt.Sprintf(onDemandConfigWithQualifierPath, pathEscape(serviceName), pathEscape(qualifier), pathEscape(functionName)), input.GetPath())

	err := input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestGetOnDemandConfig_EmptyParam() {
	assert := s.Require()

	input := NewGetOnDemandConfigInput("", qualifier, functionName)
	err := input.Validate()
	assert.Nil(err)

	input = NewGetOnDemandConfigInput(serviceName, "", functionName)
	err = input.Validate()
	assert.Nil(err)

	input = NewGetOnDemandConfigInput(serviceName, qualifier, "")
	err = input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestDeleteOnDemandConfig_Success() {
	assert := s.Require()

	input := NewDeleteOnDemandConfigInput(serviceName, qualifier, functionName)
	assert.NotNil(input)
	assert.Equal(serviceName, *input.ServiceName)
	assert.Equal(qualifier, *input.Qualifier)
	assert.Equal(functionName, *input.FunctionName)
	assert.Equal(fmt.Sprintf(onDemandConfigWithQualifierPath, pathEscape(serviceName), pathEscape(qualifier), pathEscape(functionName)), input.GetPath())
	assert.Nil(input.IfMatch)

	input.WithIfMatch(ifMatch)
	assert.NotNil(input.IfMatch)
	assert.Equal(ifMatch, *input.IfMatch)

	err := input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestDeleteOnDemandConfig_EmptyParams() {
	assert := s.Require()

	input := NewDeleteOnDemandConfigInput("", qualifier, functionName)
	err := input.Validate()
	assert.Nil(err)

	input = NewDeleteOnDemandConfigInput(serviceName, "", functionName)
	err = input.Validate()
	assert.Nil(err)

	input = NewDeleteOnDemandConfigInput(serviceName, qualifier, "")
	err = input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestListOnDemandConfigs_Success() {
	assert := s.Require()

	limit := 100
	prefix := "prefix"
	nextToken := "nextToken"
	startKey := "startKey"

	input := NewListOnDemandConfigsInput().WithPrefix(prefix).WithLimit(limit).WithNextToken(nextToken).WithStartKey(startKey)
	assert.NotNil(input)
	assert.Equal(limit, input.Limit)
	assert.Equal(fmt.Sprintf("%d", input.Limit), input.GetQueryParams().Get("limit"))
	assert.Equal(prefix, input.GetQueryParams().Get("prefix"))
	assert.Equal(nextToken, input.GetQueryParams().Get("nextToken"))
	assert.Equal(startKey, input.GetQueryParams().Get("startKey"))

	assert.Equal(onDemandConfigPath, input.GetPath())

	expectedValues := map[string][]string{
		"limit":     {"100"},
		"prefix":    {"prefix"},
		"nextToken": {"nextToken"},
		"startKey":  {"startKey"},
	}
	values := input.GetQueryParams()
	assert.Equal(url.Values(expectedValues), values)

	err := input.Validate()
	assert.Nil(err)
}

func (s *OnDemandConfigStructsTestSuite) TestListOnDemandConfigs_InvalidLimit() {
	assert := s.Require()

	input := NewListOnDemandConfigsInput()
	input.WithLimit(-1)
	err := input.Validate()
	assert.Nil(err)
}
