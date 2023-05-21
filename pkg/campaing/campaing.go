package campaign

import (
	"errors"
	"time"

	"github.com/eylmzer/campaingdemo/pkg/product"
)

type Campaign struct {
	Name         string
	Product      *product.Product
	Duration     int
	PriceLimit   float64
	TargetSales  int
	TotalSales   int
	StartTime    time.Time
	EndTime      time.Time
	CurrentTime  time.Time
	CurrentSales int
}

func NewCampaign(name string, product *product.Product, duration int, priceLimit float64, targetSales int) (*Campaign, error) {
	if duration <= 0 {
		return nil, errors.New("invalid duration")
	}
	if priceLimit < 0 {
		return nil, errors.New("invalid price limit")
	}
	if targetSales <= 0 {
		return nil, errors.New("invalid target sales")
	}

	startTime := time.Now().Truncate(24 * time.Hour)
	endTime := startTime.Add(time.Hour * time.Duration(duration))

	campaign := &Campaign{
		Name:        name,
		Product:     product,
		Duration:    duration,
		PriceLimit:  priceLimit,
		TargetSales: targetSales,
		StartTime:   startTime,
		CurrentTime: startTime,
		EndTime:     endTime,
	}

	return campaign, nil
}

func (c *Campaign) Status(t time.Time) string {
	currentTime := t
	if currentTime.Before(c.StartTime) {
		return "Not Started"
	} else if currentTime.After(c.EndTime) {
		return "Ended"
	} else {
		return "Active"
	}
}

func (c *Campaign) CalculateProductPrice(hours int, price float64) float64 {

	c.CurrentTime = c.CurrentTime.Add(time.Hour * time.Duration(hours))
	timeElapsedHours := int(c.CurrentTime.Sub(c.StartTime).Hours())

	status := c.Status(c.CurrentTime)
	if status == "Not Started" || status == "Ended" {
		return price
	}

	priceChangePercentage := (float64(timeElapsedHours) / float64(c.Duration)) * c.PriceLimit
	priceChange := (priceChangePercentage / 100) * price

	price = price - priceChange

	if price <= c.PriceLimit {
		return c.PriceLimit
	} else {
		return price
	}
}

func (c *Campaign) CalculateTurnover() float64 {
	return float64(c.CurrentSales) * c.Product.Price
}

func (c *Campaign) CalculateAverageItemPrice() float64 {
	if c.CurrentSales == 0 {
		return 0
	}
	return c.CalculateTurnover() / float64(c.CurrentSales)
}
