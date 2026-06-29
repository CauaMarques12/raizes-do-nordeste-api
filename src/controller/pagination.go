package controller

import (
	"strconv"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func getPagination(c *gin.Context) (int64, int64, *rest_err.RestErr) {
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")
	if pageQuery == "" && limitQuery == "" {
		return 0, 0, nil
	}

	page := int64(1)
	limit := int64(10)

	if pageQuery != "" {
		parsedPage, err := strconv.ParseInt(pageQuery, 10, 64)
		if err != nil || parsedPage < 1 {
			return 0, 0, rest_err.NewBadRequestError("page deve ser maior que zero")
		}
		page = parsedPage
	}

	if limitQuery != "" {
		parsedLimit, err := strconv.ParseInt(limitQuery, 10, 64)
		if err != nil || parsedLimit < 1 || parsedLimit > 100 {
			return 0, 0, rest_err.NewBadRequestError("limit deve estar entre 1 e 100")
		}
		limit = parsedLimit
	}

	return page, limit, nil
}
