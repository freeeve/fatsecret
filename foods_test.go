package fatsecret

import (
	"fmt"
	"os"
	. "gopkg.in/check.v1"
)

type FoodsSuite struct{}

var _ = Suite(&FoodsSuite{})

var (
	apikey = ""
	secret = ""
)

func init() {
	apikey = os.Getenv("FATSECRET_APIKEY")
	secret = os.Getenv("FATSECRET_SECRET")
}

func (s *FoodsSuite) TestFoodSearch(c *C) {
	fs, err := Connect(apikey, secret)
	c.Assert(err, IsNil)
	foods, err := fs.FoodSearch("banana")
	c.Assert(err, IsNil)
	fmt.Println(foods)
	c.Assert(len(foods), Equals, 20)
}

func (s *FoodsSuite) TestMultiwordFoodSearch(c *C) {
	fs, err := Connect(apikey, secret)
	c.Assert(err, IsNil)
	foods, err := fs.FoodSearch("banana pie")
	c.Assert(err, IsNil)
	fmt.Println(foods)
	c.Assert(len(foods), Equals, 20)
}
