package scrumdb

import (
	"fmt"

	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/sdk/order"
)

var orderByFields = map[string]string{
	scrumbus.OrderByID: "scrum_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
