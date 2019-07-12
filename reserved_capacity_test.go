package fc

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

func TestReservedCapacityStructs(t *testing.T) {
	suite.Run(t, new(ReservedCapacityStructsTestSuite))
}

type ReservedCapacityStructsTestSuite struct {
	suite.Suite
}

func (s *ReservedCapacityStructsTestSuite) TestListReservedCapacities() {
	assert := s.Require()
	input := NewListReservedCapacitiesInput()
	assert.NotNil(input)

	input.WithNextToken("nextToken")
	assert.NotNil(input.NextToken)
	assert.Equal("nextToken", *input.NextToken)

	input.WithLimit(int32(10))
	assert.NotNil(input.Limit)
	assert.Equal(int32(10), *input.Limit)
}