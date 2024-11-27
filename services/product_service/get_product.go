package product_service

import (
	"context"
	"os"

	"github.com/machinebox/graphql"
)

func FetchSingleProduct(id string) (string, error) {
	product_service_url := os.Getenv("PRODUCT_SERVICE_URL");
	// Initialize GraphQL client
	client := graphql.NewClient(product_service_url)

	// Define query`
	req := graphql.NewRequest(`
		query SingleProduct {
        singleProduct(id: $id) {
        status
        message
        data{
            id
            name
            quantity
            rating
            description
        }
    }
    }
	`)
	req.Var("id", id);

	// Execute the query
	var resqData struct {
		Product struct {
			Id string
			Name string `json:"name"`
		} `json:"product"`
	}

	// Make the request
	if err := client.Run(context.Background(), req, &resqData); err != nil {
		return "", err
	}
	return resqData.Product.Name, nil

}