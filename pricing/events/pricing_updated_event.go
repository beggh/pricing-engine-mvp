package events

import (
	event_models "pricing-engine/pricing/events/models"
)

type PricingUpdatedEvent struct {
	event      Event
	priceModel event_models.Price
}
