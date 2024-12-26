package model

type GetTodosRequest struct {
	IsCreateModalVisible bool `query:"isCreateModalVisible"`
	IsUpdateModalVisible bool `query:"isUpdateModalVisible"`
	TodoID               int  `query:"id"`
}
