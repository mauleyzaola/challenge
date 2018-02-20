package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mauleyzaola/challenge/domain"
)

var commands map[string]callback = map[string]callback{
	"help":   help,
	"quit":   quit,
	"use":    useBasket,
	"create": createBasket,
	"scan":   scanProduct,
	"remove": removeBasket,
	"total":  totalBasket,
}

func quit(params ...string) (interface{}, error) {
	os.Exit(0)
	return nil, nil
}

func useBasket(params ...string) (interface{}, error) {
	activeBasket = params[0]
	_, err := totalBasket(params...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func createBasket(params ...string) (interface{}, error) {
	data, err := request("POST", "/basket", nil)
	if err != nil {
		return nil, err
	}
	basket := &domain.Basket{}
	if err = json.Unmarshal(data, basket); err != nil {
		return nil, err
	}
	return basket, nil
}

func scanProduct(params ...string) (interface{}, error) {
	if len(activeBasket) == 0 {
		return nil, fmt.Errorf("there is no basket selected")
	}
	codes := append([]string{}, params...)
	input := &struct {
		Codes []string
	}{
		codes,
	}
	data, err := request("POST", fmt.Sprintf("/basket/%s/scan", activeBasket), input)
	if err != nil {
		return nil, err
	}
	var result domain.Products
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func removeBasket(params ...string) (interface{}, error) {
	if len(activeBasket) == 0 {
		return nil, fmt.Errorf("there is no basket selected")
	}
	_, err := request("DELETE", fmt.Sprintf("/basket/%s", activeBasket), nil)
	return nil, err
}

func totalBasket(params ...string) (interface{}, error) {
	if len(activeBasket) == 0 {
		return nil, fmt.Errorf("there is no basket selected")
	}

	data, err := request("GET", fmt.Sprintf("/basket/%s", activeBasket), nil)
	if err != nil {
		return nil, err
	}
	basket := &domain.Basket{}
	if err = json.Unmarshal(data, basket); err != nil {
		return nil, err
	}
	return basket, nil
}
