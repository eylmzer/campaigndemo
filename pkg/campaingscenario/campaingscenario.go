package campaingscenario

import (
	"log"
	"time"

	campaign "github.com/eylmzer/campaingdemo/pkg/campaing"
	"github.com/eylmzer/campaingdemo/pkg/order"
	"github.com/eylmzer/campaingdemo/pkg/product"
)

type CampaingScenario struct {
	Products    map[string]product.Product
	Campaigns   map[string]campaign.Campaign
	Orders      map[string]order.Order
	Logger      *log.Logger
	CurrentTime time.Time
}

func NewCampaingScenario(logger *log.Logger) *CampaingScenario {
	defaultTime := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	return &CampaingScenario{
		Products:    make(map[string]product.Product),
		Campaigns:   make(map[string]campaign.Campaign),
		Orders:      make(map[string]order.Order),
		Logger:      logger,
		CurrentTime: defaultTime,
	}
}

func (cs *CampaingScenario) IncreaseTime(hours int) {
	cs.CurrentTime = cs.CurrentTime.Add(time.Hour * time.Duration(hours))
}

func (cs *CampaingScenario) GetCurrentTime() string {
	return cs.CurrentTime.Format("15:54")
}
