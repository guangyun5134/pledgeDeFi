package services

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/kucoin"
	"pledge-backend/db"
	"pledge-backend/log"
)

type PriceService struct{}

func NewPrice() *PriceService {
	return &PriceService{}
}

type LatestPrice struct {
	Symbol    string `json:"symbol"`
	Pair      string `json:"pair"`
	Price     string `json:"price"`
	PriceE8   string `json:"price_e8"`
	UpdatedAt int64  `json:"updated_at"`
	Source    string `json:"source"`
}

// GetLatest returns the latest PLGR/USDT price with e8 precision and metadata
func (s *PriceService) GetLatest() (int, LatestPrice) {
	priceStr, err := db.RedisGetString("plgr_price")
	if err != nil || priceStr == "" {
		// fallback to in-memory value maintained by the kucoin stream
		priceStr = kucoin.PlgrPrice
	}

	updatedAtStr, _ := db.RedisGetString("plgr_price_updated_at")
	var updatedAt int64
	if updatedAtStr != "" {
		if v, err := strconv.ParseInt(updatedAtStr, 10, 64); err == nil {
			updatedAt = v
		}
	}
	if updatedAt == 0 {
		updatedAt = time.Now().Unix()
	}

	source, _ := db.RedisGetString("plgr_price_source")
	if source == "" {
		source = "kucoin"
	}

	// calculate e8 integer string
	priceE8 := "0"
	if priceStr != "" {
		if priceDec, err := decimal.NewFromString(priceStr); err == nil {
			e8 := decimal.NewFromInt(100000000)
			priceE8 = priceDec.Mul(e8).Truncate(0).String()
		} else {
			log.Logger.Sugar().Error("price parse error ", err)
		}
	}

	resp := LatestPrice{
		Symbol:    "PLGR",
		Pair:      "PLGR-USDT",
		Price:     priceStr,
		PriceE8:   priceE8,
		UpdatedAt: updatedAt,
		Source:    source,
	}

	return statecode.CommonSuccess, resp
}

