package campaign

import (
	"testing"
	"time"

	"github.com/eylmzer/campaingdemo/pkg/product"
)

func TestStatus(t *testing.T) {

	startTime := time.Now().Truncate(24 * time.Hour)

	campaing := &Campaign{
		StartTime: startTime,
		EndTime:   startTime.Add(time.Hour * time.Duration(3)),
	}

	campaing.CurrentTime = startTime.Add(time.Hour * -2)

	// Test case 1: Campaign is not started
	status := campaing.Status(campaing.CurrentTime)
	expectedStatus := "Not Started"
	if status != expectedStatus {
		t.Errorf("Expected status %s, but got %s", expectedStatus, status)
	}

	// Test case 2: Campaign is active
	status = campaing.Status(campaing.StartTime)
	expectedStatus = "Active"
	if status != expectedStatus {
		t.Errorf("Expected status %s, but got %s", expectedStatus, status)
	}

	campaing.CurrentTime = campaing.EndTime.Add(time.Hour * 3)
	// Test case 3: Campaign has ended
	status = campaing.Status(campaing.CurrentTime)
	expectedStatus = "Ended"
	if status != expectedStatus {
		t.Errorf("Expected status %s, but got %s", expectedStatus, status)
	}
}

func TestCalculateProductPrice(t *testing.T) {
	product := &product.Product{
		Code:  "P1",
		Price: 100,
	}

	startTime := time.Now().Truncate(24 * time.Hour)

	campaign := &Campaign{
		Product:      product,
		Duration:     20,
		PriceLimit:   10,
		StartTime:    startTime,
		CurrentTime:  startTime,
		EndTime:      startTime.Add(time.Hour * time.Duration(20)),
		CurrentSales: 50,
	}

	// Test case 1: Campaign not started
	price := campaign.CalculateProductPrice(-2, product.Price)
	expectedPrice := product.Price
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}

	// Test case 2: Campaign ended
	campaign.CurrentTime = campaign.EndTime.Add(3 * time.Hour)
	price = campaign.CalculateProductPrice(15, product.Price)
	expectedPrice = product.Price
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}

	// Test case 3: Campaign active
	campaign.StartTime = time.Now().Truncate(24 * time.Hour)
	campaign.CurrentTime = campaign.StartTime
	price = campaign.CalculateProductPrice(1, product.Price)
	expectedPrice = 99.5
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}

	// Test case 4: Price reaches price limit
	campaign.StartTime = time.Now().Truncate(24 * time.Hour)
	campaign.CurrentTime = campaign.StartTime
	campaign.PriceLimit = 90.0
	price = campaign.CalculateProductPrice(10, product.Price)
	expectedPrice = 90.0 // Price should not go below the price limit
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}
}

func TestCalculateTurnover(t *testing.T) {
	product := &product.Product{
		Code:  "P1",
		Price: 100,
	}

	campaign := &Campaign{
		Product:      product,
		CurrentSales: 50,
	}

	turnover := campaign.CalculateTurnover()
	expectedTurnover := 5000.0 // 50 sales * $100 price
	if turnover != expectedTurnover {
		t.Errorf("Expected turnover %.2f, but got %.2f", expectedTurnover, turnover)
	}
}
