package errors

import (
	"errors"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"

	"gorm.io/gorm"
)

func Handler(c *router.Context, err error) {
	if appErr, ok := err.(AppError); ok {
		c.Error(appErr.Code(), appErr.Message())
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(http.StatusNotFound, NOT_FOUND)
		return
	}

	c.Error(http.StatusInternalServerError, err.Error())
}
