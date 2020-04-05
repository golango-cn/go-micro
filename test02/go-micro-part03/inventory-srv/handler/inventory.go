package handler

import (
	"context"
	"go-micro/go-micro-part03/inventory-srv/model"
	"go-micro/go-micro-part03/proto/inventory"
)

type InventoryHandler struct {
}

func (i InventoryHandler) Sell(ctx context.Context, requ *inventory.Request, resp *inventory.Response) error {

	inv := model.GetInventoryService()
	id, err := inv.Sell(requ.BookId, requ.UserId)

	if err != nil {
		return err
	}

	resp.Inv = &inventory.Inv{
		Id: id,
	}
	resp.Success = true

	return nil
}

func (i InventoryHandler) Confirm(ctx context.Context, requ *inventory.Request, resp *inventory.Response) error {

	inv := model.GetInventoryService()
	err := inv.Confirm(requ.HistoryId)

	if err != nil {
		return err
	}

	resp.Success = true
	return nil
}
