package main

import (
	"di_example/mocks"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	service mocks.ServiceRepo
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	suite.service.On("GetNamesService").Return("Hey!!")
	actual := suite.service.GetNamesService()
	assert.Equal(suite.T(), actual, "Hey!!")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}