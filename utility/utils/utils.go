package utils

import (
	"context"
	"go_ticket/internal/errorcode"

	"github.com/jinzhu/copier"
)

func MyCopy(ctx context.Context, toValue interface{}, fromValue interface{}) (err error) {
	if err = copier.Copy(toValue, fromValue); err != nil {
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
	} else {
		return nil
	}
}
