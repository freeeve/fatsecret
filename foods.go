package fatsecret

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Food struct {
	Id          string `json:"food_id,omitempty"`
	Name        string `json:"food_name,omitempty"`
	Type        string `json:"food_type,omitempty"`
	Url         string `json:"food_url,omitempty"`
	Description string `json:"food_description,omitempty"`
}

type FoodSearchResponse struct {
	Foods *FoodSearchResponseFoods `json:"foods,omitempty"`
	Error *ErrorResponse           `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type FoodSearchResponseFoods struct {
	PageSize     int    `json:"max_results,string"`
	TotalResults int    `json:"total_results,string"`
	PageNumber   int    `json:"page_number,string"`
	Food         []Food `json:"food"`
}

func (fs FatSecretConn) FoodSearch(query string) ([]Food, error) {
	resp, err := fs.get(
		"foods.search",
		map[string]string{"search_expression": query},
	)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	//fmt.Println(string(body))
	foodresp := FoodSearchResponse{}
	err = json.Unmarshal(body, &foodresp)
	if err != nil {
		return nil, err
	}
	if foodresp.Error != nil {
		return nil, errors.New(foodresp.Error.Message)
	}
	return foodresp.Foods.Food, nil
}
