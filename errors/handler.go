package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"

	"gorm.io/gorm"
)

func Handler(c *router.Context, err error) {
	fmt.Printf("%T\n", err)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(http.StatusNotFound, NOT_FOUND)
		return
	}

	// if route not found
	fmt.Printf("%T\n", err)

	c.Error(http.StatusInternalServerError, err.Error())
}
