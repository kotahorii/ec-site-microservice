package loader

import (
	"github.com/graph-gophers/dataloader"
)

type Loaders struct {
	OrderLoader *dataloader.Loader
}

func NewLoaders(
	orderLoader *dataloader.Loader,
) *Loaders {
	return &Loaders{
		OrderLoader: orderLoader,
	}
}
