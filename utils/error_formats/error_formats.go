package error_formats

import (
	"strings"

	"github.com/Bobby-P-dev/FinalProject3_kel7/utils/error_utils"
)

func ParseError(err error) error_utils.MessageErr {

	if strings.Contains(err.Error(), "no rows in result set") {
		return error_utils.NewNotFoundError("no record found")
	}
	return error_utils.NewInternalServerErrorr("something went wrong")
}
