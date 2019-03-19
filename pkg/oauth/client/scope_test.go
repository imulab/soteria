package client

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestEqualityScopeStrategy(t *testing.T) {
	suite.Run(t, new(EqualityScopeStrategyTestSuite))
}

type EqualityScopeStrategyTestSuite struct {
	suite.Suite
}

func (suite *EqualityScopeStrategyTestSuite) TestAccept() {
	for _, t := range []struct{
		num 			int
		ignoreCase 		bool
		clientScopes 	[]string
		testScope		string
		expectPass 		bool
	}{
		{
			num: 1,
			ignoreCase: false,
			clientScopes: []string{"foo", "bar"},
			testScope: "foo",
			expectPass: true,
		},
		{
			num: 2,
			ignoreCase: false,
			clientScopes: []string{"foo", "bar"},
			testScope: "FOO",
			expectPass: false,
		},
		{
			num: 3,
			ignoreCase: true,
			clientScopes: []string{"foo", "bar"},
			testScope: "FOO",
			expectPass: true,
		},
		{
			num: 4,
			ignoreCase: false,
			clientScopes: []string{"foo", "bar"},
			testScope: "baz",
			expectPass: false,
		},
	}{
		// given
		strategy := EqualityScopeStrategy{IgnoreCase:t.ignoreCase}
		client := &DefaultClient{Scopes:t.clientScopes}

		// when
		accepts := strategy.Accepts(client, t.testScope)

		// then
		suite.Assert().Equal(t.expectPass, accepts, fmt.Sprintf("test #%d", t.num))
	}
}

func (suite *EqualityScopeStrategyTestSuite) TestAcceptAll() {
	for _, t := range []struct{
		num 			int
		ignoreCase 		bool
		clientScopes 	[]string
		testScopes		[]string
		expectPass 		bool
	}{
		{
			num: 1,
			ignoreCase: false,
			clientScopes: []string{"foo", "bar"},
			testScopes: []string{"foo", "bar"},
			expectPass: true,
		},
		{
			num: 2,
			ignoreCase: false,
			clientScopes: []string{"foo", "bar"},
			testScopes: []string{"foo", "baz"},
			expectPass: false,
		},
	}{
		// given
		strategy := EqualityScopeStrategy{IgnoreCase:t.ignoreCase}
		client := &DefaultClient{Scopes:t.clientScopes}

		// when
		accepts := strategy.AcceptsAll(client, t.testScopes)

		// then
		suite.Assert().Equal(t.expectPass, accepts, fmt.Sprintf("test #%d", t.num))
	}
}