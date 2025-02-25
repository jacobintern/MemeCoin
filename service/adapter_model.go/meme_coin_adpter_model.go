package adapter_model

import (
	"time"

	"github.com/jacobintern/meme_coin/controller/response"
)

type MemeCoinInfo struct {
	ID              uint64
	Name            string
	Description     *string
	CreatedAt       time.Time
	PopularityScore uint
}

func NewMemeCoinInfo() (e MemeCoinInfo) {
	return e
}

func (e MemeCoinInfo) ToMemeCoinInfoResp() (res response.MemeCoinInfoResp) {
	res.ID = e.ID
	res.Name = e.Name
	if e.Description != nil {
		res.Description = e.Description
	}
	res.CreatedAt = e.CreatedAt
	res.PopularityScore = e.PopularityScore

	return res
}
