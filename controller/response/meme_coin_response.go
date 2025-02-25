package response

import "time"

type MemeCoinInfoResp struct {
	ID              uint64    `json:"id"`
	Name            string    `json:"name"`
	Description     *string   `json:"description,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	PopularityScore uint      `json:"popularity_score"`
}

func NewMemeCoinInfoResp() MemeCoinInfoResp {
	return MemeCoinInfoResp{}
}
