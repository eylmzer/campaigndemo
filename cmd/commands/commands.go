package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	campaign "github.com/eylmzer/campaingdemo/pkg/campaing"
	"github.com/eylmzer/campaingdemo/pkg/campaingscenario"
	"github.com/eylmzer/campaingdemo/pkg/order"
	prd "github.com/eylmzer/campaingdemo/pkg/product"
)

type Command struct {
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

func ExecuteCommand(cmds string, cs *campaingscenario.CampaingScenario) (string, error) {
	parts := strings.Split(cmds, " ")
	switch parts[0] {
	case "create_product":
		productCode := parts[1]
		price, _ := strconv.ParseFloat(parts[2], 64)
		stock, _ := strconv.Atoi(parts[3])
		product := prd.NewProduct(productCode, price, stock)
		cs.Products[product.Code] = *product
		return fmt.Sprintf("Product created; code %s, price %.2f, stock %d", productCode, price, stock), nil

	case "create_order":
		productCode := parts[1]
		quantity, err := strconv.Atoi(parts[2])
		if err != nil {
			return "", errors.New("invalid quantity value")
		}

		product, ok := cs.Products[productCode]
		if !ok {
			return "", errors.New("product not found")
		}

		if quantity > product.Stock {
			return "", errors.New("insufficient stock")
		}

		price := product.Price
		order, err := order.NewOrder(&product, quantity)
		if err != nil {
			return "", errors.New("failed to create order")
		}
		cs.Orders[order.Product.Code] = *order

		// Update product stock
		product.Stock -= quantity

		return fmt.Sprintf("Order created; product %s, quantity %d, price %.2f", productCode, quantity, price), nil

	case "create_campaign":
		campaignName := parts[1]
		productCode := parts[2]
		duration, _ := strconv.Atoi(parts[3])
		priceLimit, _ := strconv.ParseFloat(parts[4], 64)
		targetSales, _ := strconv.Atoi(parts[5])

		product, ok := cs.Products[productCode]
		if !ok {
			return "", errors.New("product not found")
		}

		campaign, err := campaign.NewCampaign(campaignName, &product, duration, priceLimit, targetSales)
		if err != nil {
			return "", errors.New("failed to create campaign")
		}

		cs.Campaigns[campaign.Name] = *campaign
		return fmt.Sprintf("Campaign created; name %s, product %s, duration %d, limit %.2f, target sales count %d",
			campaignName, productCode, duration, priceLimit, targetSales), nil

	case "get_product_info":
		productCode := parts[1]
		product, ok := cs.Products[productCode]
		if !ok {
			return "", errors.New("product not found")
		}
		return fmt.Sprintf("Product %s info; price %.2f, stock %d", productCode, product.Price, product.Stock), nil

	case "increase_time":
		hours, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("invalid time value")
		}

		for _, c := range cs.Campaigns {
			p := cs.Products[c.Product.Code]
			price := c.CalculateProductPrice(hours, p.Price)
			p.Price = price
			cs.Products[c.Product.Code] = p
		}

		cs.IncreaseTime(hours)
		return fmt.Sprintf("Time is %s", cs.GetCurrentTime()), nil

	case "get_campaign_info":
		campaignName := parts[1]
		campaign, ok := cs.Campaigns[campaignName]
		if !ok {
			return "", errors.New("campaign not found")
		}
		return fmt.Sprintf("Campaign %s info; Status %s, Target Sales %d, Total Sales %d, Turnover %.2f, Average Item Price %.2f",
			campaignName, campaign.Status(campaign.StartTime), campaign.TargetSales, campaign.CurrentSales, campaign.CalculateTurnover(), campaign.CalculateAverageItemPrice()), nil

	default:
		return "", errors.New("invalid command")
	}
}
