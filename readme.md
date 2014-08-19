# fatsecret for golang

[![Build Status](https://travis-ci.org/wfreeman/fatsecret.svg?branch=master)](https://travis-ci.org/wfreeman/fatsecret)
[![Coverage Status](https://img.shields.io/coveralls/wfreeman/fatsecret.svg)](https://coveralls.io/r/wfreeman/fatsecret?branch=master)
[![Waffle](https://badge.waffle.io/wfreeman/fatsecret.png?label=ready)](https://waffle.io/wfreeman/fatsecret)
[![Gitter chat](https://badges.gitter.im/wfreeman/fatsecret.png)](https://gitter.im/wfreeman/fatsecret)

Currently experimental. Only planning to support the queries that aren't user-oriented initially (signed but not delegated): [authentication docs](http://platform.fatsecret.com/api/Default.aspx?screen=rapiauth)

You do need an API key. To run unit tests, two environment variables are required: `FATSECRET_APIKEY`, and `FATSECRET_SECRET`, with the appropriate values (the consumer key and the shared secret).

## minimum viable snippet
```
import (
  "github.com/wfreeman/fatsecret"
)

func main() {
  fs, err := fatsecret.Connect(apikey, secret)
  if err != nil {
    panic(err)
  }
  
  foods, err := fs.FoodSearch("banana")
  if err != nil {
    panic(err)
  }
  
  fmt.Println(foods[0])
}
```

outputs:
````
{35755 Bananas Generic http://www.fatsecret.com/calories-nutrition/usda/bananas Per 100g - Calories: 89kcal | Fat: 0.33g | Carbs: 22.84g | Protein: 1.09g []}
```
