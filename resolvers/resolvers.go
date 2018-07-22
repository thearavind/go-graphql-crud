package resolvers

import(
	"github.com/jinzhu/gorm"
)

type Resolver struct{
	DB *gorm.DB
}