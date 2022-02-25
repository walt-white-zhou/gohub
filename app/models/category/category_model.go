//Package category 模型
package category

import (
	"gohub/app/models"
)

type Category struct {
	models.BaseModel

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}
