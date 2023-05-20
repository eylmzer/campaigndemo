package campaign

import (
	"testing"
	"time"

	"github.com/eylmzer/campaingdemo/pkg/product"
)

func TestStatus(t *testing.T) {
	campaign := &Campaign{
		StartTime: time.Now().Add(-time.Hour),
		EndTime:   time.Now().Add(time.Hour),
	}

	// Test case 1: Campaign is not started
	campaign.StartTime = time.Now().Add(5 * time.Hour)
	status := campaign.Status()
	expectedStatus := "Not Started"
	if status != expectedStatus {
		t.Errorf("Expected status %s, but got %s", expectedStatus, status)
	}

	// Test case 2: Campaign is active
	campaign.StartTime = time.Now().Add(-time.Minute)
	campaign.EndTime = time.Now().Add(time.Minute)
	status = campaign.Status()
	expectedStatus = "Active"
	if status != expectedStatus {
		t.Errorf("Expected status %s, but got %s", expectedStatus, status)
	}

	// Test case 3: Campaign has ended
	campaign.StartTime = time.Now().Add(-time.Hour)
	campaign.EndTime = time.Now().Add(-time.Minute)
	status = campaign.Status()
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

	campaign := &Campaign{
		Product:      product,
		Duration:     10,
		PriceLimit:   10,
		StartTime:    time.Now().Add(-time.Hour),
		EndTime:      time.Now().Add(time.Hour),
		CurrentSales: 50,
	}

	// Test case 1: Campaign not started
	campaign.StartTime = time.Now().Add(5 * time.Hour)
	price := campaign.CalculateProductPrice(2, product.Price)
	expectedPrice := product.Price
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}

	// Test case 2: Campaign ended
	campaign.StartTime = time.Now().Add(-2 * time.Hour)
	campaign.EndTime = time.Now().Add(-time.Hour)
	price = campaign.CalculateProductPrice(2, product.Price)
	expectedPrice = product.Price
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}

	// Test case 3: Campaign active
	campaign.StartTime = time.Now().Add(-time.Hour)
	campaign.EndTime = time.Now().Add(time.Hour)
	price = campaign.CalculateProductPrice(2, product.Price)
	expectedPrice = 80.0 // Assuming price decreases by 20% every hour
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, price)
	}

	// Test case 4: Price reaches price limit
	campaign.StartTime = time.Now().Add(-time.Hour)
	campaign.EndTime = time.Now().Add(time.Hour)
	campaign.PriceLimit = 90.0
	price = campaign.CalculateProductPrice(2, product.Price)
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
