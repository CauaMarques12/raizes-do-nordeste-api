package model

type MenuItemDomainInterface interface {
	GetProductID() string
	GetName() string
	GetDescription() string
	GetCategory() string
	GetPriceCents() int64
	GetQuantity() int64
	GetAvailable() bool
}

func NewMenuItemDomain(
	productID, name, description, category string,
	priceCents, quantity int64,
) MenuItemDomainInterface {
	return &menuItemDomain{
		productID:   productID,
		name:        name,
		description: description,
		category:    category,
		priceCents:  priceCents,
		quantity:    quantity,
		available:   quantity > 0,
	}
}

type menuItemDomain struct {
	productID   string
	name        string
	description string
	category    string
	priceCents  int64
	quantity    int64
	available   bool
}

func (md *menuItemDomain) GetProductID() string {
	return md.productID
}

func (md *menuItemDomain) GetName() string {
	return md.name
}

func (md *menuItemDomain) GetDescription() string {
	return md.description
}

func (md *menuItemDomain) GetCategory() string {
	return md.category
}

func (md *menuItemDomain) GetPriceCents() int64 {
	return md.priceCents
}

func (md *menuItemDomain) GetQuantity() int64 {
	return md.quantity
}

func (md *menuItemDomain) GetAvailable() bool {
	return md.available
}
