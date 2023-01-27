package middlewarepkg

import (
	"context"
	"konntent-authentication-service/pkg/utils"
)

func GetHeaderMap(c context.Context) map[string]string {
	var val = c.Value(utils.HeaderMapCtx)
	if conv, ok := val.(map[string]string); ok {
		return conv
	}

	return val.(map[string]string)
}
