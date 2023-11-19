package storage

import "market/model"

type StorageI interface {
	Category() CategoryRepoI
	Product() ProductRepoI
	Branch() BranchRepoI
	Client() ClientRepoI
	Order() OrderRepoI
	OrderProduct() OrderProductsRepoI
}

type CategoryRepoI interface {
	Create(model.CategoryCreate) (*model.Category, error)
	GetById(model.CategoryPrimaryKey) (*model.Category, error)
	GetList(model.GetListCategoryRequest) (*model.GetListCategoryResponse, error)
	Update(req model.CategoryUpdate) (string, error)
	Delete(req model.CategoryPrimaryKey) (string, error)
}

type ProductRepoI interface {
	Create(model.ProductCreate) (*model.Product, error)
	GetById(model.ProductPrimaryKey) (*model.Product, error)
	GetList(model.GetListProductRequest) (*model.GetListProductResponse, error)
	Update(req model.ProductUpdate) (string, error)
	Delete(req model.ProductPrimaryKey) (string, error)
}

type BranchRepoI interface {
	Create(model.BranchCreate) (*model.Branch, error)
	GetById(model.BranchPrimaryKey) (*model.Branch, error)
	GetList(model.GetListBranchRequest) (*model.GetListBranchResponse, error)
	Update(req model.BranchUpdate) (string, error)
	Delete(req model.BranchPrimaryKey) (string, error)
}

type ClientRepoI interface {
	Create(model.ClientCreate) (*model.Client, error)
	GetById(model.ClientPrimaryKey) (*model.Client, error)
	GetList(model.GetListClientRequest) (*model.GetListClientResponse, error)
	Update(req model.ClientUpdate) (string, error)
	Delete(req model.ClientPrimaryKey) (string, error)
}

type OrderRepoI interface {
	Create(model.OrderCreate) (*model.Order, error)
	GetById(model.OrderPrimaryKey) (*model.Order, error)
	GetList(model.GetListOrderRequest) (*model.GetListOrderResponse, error)
	Update(req model.OrderUpdate) (string, error)
	StatusUpdate(req model.OrderStatusUpdate) (string, error)
	Delete(req model.OrderPrimaryKey) (string, error)
}

type OrderProductsRepoI interface {
	Create(model.AddOrderProduct) (*model.OrderProducts, error)
	GetByOrderId(req model.GetOrderProductByOrderId) (*[]model.OrderProducts, error)
	Delete(req model.GetOrderProductByOrderId) (string, error)
}
