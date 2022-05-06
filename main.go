package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"net/http"
	"os"
)

type Product struct {
	Name    string `mapstructure:"name" json:"name"`
	Url     string `mapstructure:"url" json:"url"`
	InStock bool   `json:"in_stock"`
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config file path")
	flag.Parse()

	if configPath == "" {
		fmt.Print(`Config not passed!

Usage:

  ./is-ui-available -config=config.yaml

Sample config.yaml:

  products:
    - name: "Dream Router"
      url: "https://store.ui.com/collections/unifi-network-unifi-os-consoles/products/dream-router"

`)
		os.Exit(1)
	}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("sample_config.yaml")
	if err != nil {
		panic(err)
	}

	var products []Product
	err = config.BindStruct("products", &products)
	if err != nil {
		panic(err)
	}

	var inStockProducts []Product

	for _, product := range products {
		resp, err := http.Get(product.Url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			continue
		}

		selection := doc.Find(".comProduct .comProduct__title--wrapper .comProductTile__soldOut")
		inStock := selection.Index() < 0

		if inStock {
			inStockProducts = append(inStockProducts, Product{Name: product.Name, Url: product.Url, InStock: inStock})
		}
	}

	jsonString, err := json.Marshal(inStockProducts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", jsonString)
}
