package model

type GetTodosRequest struct {
	IsUpdateModalVisible bool `query:"isUpdateModalVisible"`
	TodoID               int  `query:"id"`
}
