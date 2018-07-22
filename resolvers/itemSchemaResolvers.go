package resolvers


import(
	"context"
	"strconv"
	graphql "github.com/graph-gophers/graphql-go"
)

type Item struct{
	Id int32 `gorm:"primary_key"`
	Name string
	Price string
	Stock int32
}

type DeleteResponse struct {
	Success bool
}

func (p *Item) NAME(ctx context.Context) (*string, error) {
	return &p.Name, nil
}

func (p *Item) ID(ctx context.Context) (graphql.ID, error) {
	id := strconv.Itoa(int(p.Id))
	return graphql.ID(id), nil
}

func (p *Item) PRICE(ctx context.Context) *string {
	return &p.Price
}

func (p *Item) STOCK(ctx context.Context) int32 {
	return int32(p.Stock)
}

func (d *DeleteResponse) SUCCESS(ctx context.Context) (*bool, error) {
	return &d.Success, nil
}