package validate

import (
	"regexp"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
)

func AccountName() pluginsdk.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile(`^[a-zA-Z0-9][-a-zA-Z0-9]{1,61}[a-zA-Z0-9]$`),
		"The Purview account name must be between 3 and 63 characters long, it can contain only letters, numbers and hyphens, and the first and last characters must be a letter or number.")
}
