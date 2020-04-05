package model

import (
	"fmt"
	"github.com/micro/go-micro/v2/logger"
)

type InventoryService struct {
	keys map[int64]int64
}

// 销减库存
func (m *InventoryService) Sell(bookId int64, userId int64) (int64, error) {

	// 操作数据库销减库存
	logger.Info("库存扣减成功", bookId, userId)

	//id := int64(rand.Intn(1000))
	m.keys[bookId] = bookId

	return bookId, nil
}

//确认
func (m *InventoryService) Confirm(id int64) error {

	if _, ok := m.keys[id]; ok {
		//delete(m.keys, id)

		logger.Info("库存确认成功")
		return nil
	}

	return fmt.Errorf("没有找到编号：%d", id)
}

var inventoryService *InventoryService

func Init()  {
	inventoryService = &InventoryService{keys: make(map[int64]int64, )}
}

func GetInventoryService() *InventoryService {
	return inventoryService
}
