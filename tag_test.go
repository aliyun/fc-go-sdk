package fc

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTagStructs(t *testing.T) {
	suite.Run(t, new(TagStructsTestSuite))
}

type TagStructsTestSuite struct {
	suite.Suite
}

func (s *TagStructsTestSuite) TestTagResource() {
	assert := s.Require()

	input := NewTagResourceInput("soureArn", map[string]string{})
	assert.NotNil(input)
	assert.Equal("soureArn", *input.ResourceArn)
	assert.True(reflect.DeepEqual(map[string]string{}, input.Tags))

	input.WithResourceArn("services/foo")
	assert.NotNil(input.ResourceArn)
	assert.Equal("services/foo", *input.ResourceArn)

	input.WithTags(map[string]string{"k": "v"})
	assert.NotNil(input.Tags)
	assert.True(reflect.DeepEqual(map[string]string{
		"k": "v",
	}, input.Tags))
}

func (s *TagStructsTestSuite) TestUnTagResource() {
	assert := s.Require()

	input := NewUnTagResourceInput("soureArn")
	assert.NotNil(input)
	assert.Equal("soureArn", *input.ResourceArn)

	input.WithResourceArn("services/foo")
	assert.NotNil(input.ResourceArn)
	assert.Equal("services/foo", *input.ResourceArn)

	input.WithTagKeys([]string{"k2", "k2"})
	assert.NotNil(input.TagKeys)
	assert.True(reflect.DeepEqual([]string{"k2", "k2"}, input.TagKeys))

	input.WithAll(true)
	assert.Equal(true, *input.All)
}

func (s *TagStructsTestSuite) TestGetResourceTags() {
	assert := s.Require()

	input := NewGetResourceTagsInput("soureArn")
	assert.NotNil(input)
	assert.Equal("soureArn", *input.ResourceArn)

	input.WithResourceArn("services/foo")
	assert.NotNil(input.ResourceArn)
	assert.Equal("services/foo", *input.ResourceArn)
}
