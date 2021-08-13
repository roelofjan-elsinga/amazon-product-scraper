package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly/v2"
	"github.com/roelofjan-elsinga/lambda/models"
	"net/http"
)

type LambdaRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type LambdaResponse struct {
	SalePrice     string `json:"sale_price"`
	OriginalPrice string `json:"original_price"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}

func main() {
	lambda.Start(HttpHandler)
}

func HttpHandler(ctx context.Context, event models.GatewayRequest) (models.GatewayResponse, error) {

	var body LambdaRequest

	err := json.Unmarshal([]byte(event.Body), &body)

	if err != nil {
		return models.GatewayResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusInternalServerError,
			Body:            err.Error(),
		}, err
	}

	responseBody, err := Action(ctx, body)

	if err != nil {
		return models.GatewayResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusInternalServerError,
			Body:            err.Error(),
		}, err
	}

	jsonBody, err := json.Marshal(responseBody)

	if err != nil {
		return models.GatewayResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusInternalServerError,
			Body:            err.Error(),
		}, err
	}

	return models.GatewayResponse{
		IsBase64Encoded: false,
		StatusCode:      http.StatusOK,
		Body:            string(jsonBody),
	}, nil
}

func Action(ctx context.Context, event LambdaRequest) (*LambdaResponse, error) {
	c := colly.NewCollector()

	product := LambdaResponse{}

	// Find and visit all links
	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		product.Title = e.Text
	})

	c.OnHTML("#productDescription", func(e *colly.HTMLElement) {
		product.Title = e.Text
	})

	c.OnHTML("#priceblock_saleprice", func(e *colly.HTMLElement) {
		product.SalePrice = e.Text
	})

	c.OnHTML("#priceBlockStrikePriceString a-text-strike", func(e *colly.HTMLElement) {
		product.OriginalPrice = e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(event.URL)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
