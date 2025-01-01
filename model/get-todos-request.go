package model

type GetTodosRequest struct {
	IsUpdateModalVisible bool  `query:"isUpdateModalVisible"`
	IsDeleteModalVisible bool  `query:"isDeleteModalVisible"`
	TodoID               int64 `query:"id"`
}
