package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	type Berry struct {
		Firmness struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"firmness"`
		Flavors []struct {
			Flavor struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"flavor"`
			Potency int `json:"potency"`
		} `json:"flavors"`
		GrowthTime int `json:"growth_time"`
		ID         int `json:"id"`
		Item       struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		MaxHarvest       int    `json:"max_harvest"`
		Name             string `json:"name"`
		NaturalGiftPower int    `json:"natural_gift_power"`
		NaturalGiftType  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"natural_gift_type"`
		Size        int `json:"size"`
		Smoothness  int `json:"smoothness"`
		SoilDryness int `json:"soil_dryness"`
	}

	type Item struct {
		Name string `json:"name"`
		Cost int    `json:"cost"`
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/berry/1/")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var berry Berry
	err = decoder.Decode(&berry)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	// fmt.Println("Berry Data:", berry)
	fmt.Printf("Berry ID: %d\n", berry.ID)
	fmt.Printf("Berry Name: %s\nGrowth Time: %d\n", berry.Name, berry.GrowthTime)
	itemURL := berry.Item.URL

	resp2, err := http.Get(itemURL)
	if err != nil {
		fmt.Println("Error fetching item data:", err)
		return
	}
	defer resp2.Body.Close()
	decoder2 := json.NewDecoder(resp2.Body)

	var item Item
	err = decoder2.Decode(&item)
	if err != nil {
		fmt.Println("Error unmarshalling item JSON:", err)
		return
	}

	fmt.Println("Item Name:", item.Name)
	fmt.Println("Item Cost:", item.Cost)
}