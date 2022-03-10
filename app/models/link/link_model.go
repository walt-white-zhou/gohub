//Package link 模型
package link

import (
	"gohub/app/models"
)

type Link struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	models.CommonTimestampsField
}
