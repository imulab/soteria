package token

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth/request"
	"github.com/imulab/soteria/pkg/utility"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestHmacSha256AuthorizeCodeStrategy(t *testing.T) {
	signingKey, err := utility.RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}

	codeStrategy, err := NewHmacSha256AuthorizeCodeStrategy(16, signingKey)
	if err != nil {
		t.Fatal(err)
	}

	suite.Run(t, &HmacShaAuthorizeCodeStrategyTestSuite{strategy:codeStrategy})
}

func TestHmacSha384AuthorizeCodeStrategy(t *testing.T) {
	signingKey, err := utility.RandomBytes(48)
	if err != nil {
		t.Fatal(err)
	}

	codeStrategy, err := NewHmacSha384AuthorizeCodeStrategy(16, signingKey)
	if err != nil {
		t.Fatal(err)
	}

	suite.Run(t, &HmacShaAuthorizeCodeStrategyTestSuite{strategy:codeStrategy})
}

func TestHmacSha512AuthorizeCodeStrategy(t *testing.T) {
	signingKey, err := utility.RandomBytes(64)
	if err != nil {
		t.Fatal(err)
	}

	codeStrategy, err := NewHmacSha512AuthorizeCodeStrategy(16, signingKey)
	if err != nil {
		t.Fatal(err)
	}

	suite.Run(t, &HmacShaAuthorizeCodeStrategyTestSuite{strategy:codeStrategy})
}

func TestMemoryAuthorizeCodeRepository(t *testing.T) {
	suite.Run(t, new(MemoryAuthorizeCodeRepositoryTestSuite))
}

type HmacShaAuthorizeCodeStrategyTestSuite struct {
	suite.Suite
	strategy 	AuthorizeCodeStrategy
}

func (s *HmacShaAuthorizeCodeStrategyTestSuite) TestNewCode() {
	for _, t := range []struct{
		num 		int
		ctxGen 		func() context.Context
		expectCode	bool
	}{
		{
			num: 1,
			ctxGen: func() context.Context {
				return context.Background()
			},
			expectCode: true,
		},
		{
			num: 2,
			ctxGen: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			expectCode: false,
		},
	}{
		code, err := s.strategy.NewCode(nil, t.ctxGen())
		if t.expectCode {
			s.Assert().NotEmpty(code, "code in #%d should not be empty", t.num)
			s.Assert().Nil(err, "err in #%d should be nil", t.num)
		} else {
			s.Assert().Empty(code, "code in #%d should be empty", t.num)
			s.Assert().NotNil(err, "err in #%d should be non-nil", t.num)
		}
	}
}

func (s *HmacShaAuthorizeCodeStrategyTestSuite) TestValidateCode()  {
	for _, t := range []struct{
		num 		int
		codeGen 	func() string
		ctxGen 		func() context.Context
		expectValid bool
	}{
		{
			num: 1,
			codeGen: func() string {
				code, _ := s.strategy.NewCode(nil, context.Background())
				s.Require().NotEmpty(code)
				return code
			},
			ctxGen: func() context.Context {
				return context.Background()
			},
			expectValid: true,
		},
		{
			num: 2,
			codeGen: func() string {
				code, _ := s.strategy.NewCode(nil, context.Background())
				s.Require().NotEmpty(code)
				return code + "tempered"	// tempered with signature
			},
			ctxGen: func() context.Context {
				return context.Background()
			},
			expectValid: false,
		},
		{
			num: 3,
			codeGen: func() string {
				code, _ := s.strategy.NewCode(nil, context.Background())
				s.Require().NotEmpty(code)
				return code
			},
			ctxGen: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			expectValid: false,
		},
	}{
		err := s.strategy.ValidateCode(t.codeGen(), nil, t.ctxGen())
		if t.expectValid {
			s.Assert().Nil(err, "err in #%d should be nil", t.num)
		} else {
			s.Assert().NotNil(err, "err in #%d should be non-nil", t.num)
		}
	}
}

type MemoryAuthorizeCodeRepositoryTestSuite struct {
	suite.Suite
}

func (s *MemoryAuthorizeCodeRepositoryTestSuite) TestCRUD() {
	var (
		err error
		repo = NewMemoryAuthorizeCodeRepository()
		foo = "foo"
		bar = "bar"
	)

	// save
	err = repo.Save(
		foo,
		request.NewOAuthAuthorizeRequest(),
		context.Background(),
		)
	s.Assert().Nil(err)

	// get
	_, err = repo.GetSession(foo, context.Background())
	s.Assert().Nil(err)

	_, err = repo.GetSession(bar, context.Background())
	s.Assert().NotNil(err)

	// del
	err = repo.Delete(foo, context.Background())
	s.Assert().Nil(err)
	_, err = repo.GetSession(foo, context.Background())
	s.Assert().NotNil(err)
}