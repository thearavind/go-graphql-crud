package resolvers

import (
	"context"
)

var message = "World"

type newProductArgs struct{ 
	Name *string 
	Stock *int32 
	Price *string 
}

type stocksArgs struct{ 
	Unit *int32 
	ID int32
}

type priceArgs struct{ 
	Price *string 
	ID int32
}

func (r *Resolver) Products(ctx context.Context) (*[]*Item, error) {
	var foundUser []*Item
	r.DB.Find(&foundUser)
	return &foundUser, nil
}

func (r *Resolver) Product(ctx context.Context, args struct{ ID int32}) (*Item, error) {
	foundUser := &Item{}
	r.DB.First(foundUser, Item{Id: args.ID})
	if(foundUser.Id == 0) {
		return nil, nil
	}
	return foundUser, nil
}

func (r *Resolver) NewProduct(ctx context.Context, args newProductArgs) (*Item, error) {
	var newProduct = &Item{Name: *args.Name, Stock: *args.Stock, Price: *args.Price}
	r.DB.Create(newProduct)
	return newProduct, nil
}

func (r *Resolver) AddStock(ctx context.Context, args stocksArgs) (*Item, error) {
	foundUser := &Item{}
	r.DB.First(foundUser, Item{Id: args.ID})
	foundUser.Stock += *args.Unit
	r.DB.Save(foundUser)
	return foundUser, nil
}

func (r *Resolver) RemoveStock(ctx context.Context, args stocksArgs) (*Item, error) {
	foundUser := &Item{}
	r.DB.First(foundUser, Item{Id: args.ID})
	foundUser.Stock -= *args.Unit
	r.DB.Save(foundUser)
	return foundUser, nil
}

func (r *Resolver) UpdatePrice(ctx context.Context, args priceArgs) (*Item, error) {
	foundUser := &Item{}
	r.DB.First(foundUser, Item{Id: args.ID})
	foundUser.Price = *args.Price
	r.DB.Save(foundUser)
	return foundUser, nil
}

func (r *Resolver) DeleteProduct(ctx context.Context, args struct{ ID int32 }) (*DeleteResponse, error) {
	r.DB.Delete(&Item{Id: args.ID})
	return &DeleteResponse{Success: true}, nil
}