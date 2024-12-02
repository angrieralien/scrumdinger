package scrumapp

import (
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
)

var orderByFields = map[string]string{
	"scrum_id": scrumbus.OrderByID,
}
