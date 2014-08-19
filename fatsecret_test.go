package fatsecret

import (
	"os"
	"testing"
	. "gopkg.in/check.v1"
)

type FatSecretSuite struct{}

var _ = Suite(&FatSecretSuite{})

func Test(t *testing.T) {
	TestingT(t)
}

func (s *FatSecretSuite) TestConnect(c *C) {
	apikey := os.Getenv("FATSECRET_APIKEY")
	secret := os.Getenv("FATSECRET_SECRET")
	_, err := Connect(apikey, secret)
	c.Assert(err, IsNil)
}
