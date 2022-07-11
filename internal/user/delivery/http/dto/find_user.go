package dto

import "github.com/dinorain/checkoutaja/pkg/utils"

type FindUserResponseDto struct {
	Meta utils.PaginationMetaDto `json:"meta"`
	Data interface{}             `json:"data"`
}
