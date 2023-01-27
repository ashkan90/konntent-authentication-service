package utils

import (
	"context"
	"konntent-authentication-service/pkg/validation"
)

func ValidateRequestWithCtx(c context.Context, r interface{}) map[string]string {
	v := c.Value(Validator).(validation.IValidator)
	if errs := v.Validate(r); len(errs) > 0 {
		return errs
	}

	return nil
}
