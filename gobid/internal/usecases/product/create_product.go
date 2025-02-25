package product

import (
	"context"
	"gobid/internal/validator"
	"time"
)

type CreateProductReq struct {
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

const minAuctionDuration = 2 * time.Hour

func (cpr *CreateProductReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(cpr.ProductName), "product_name", "this field is required")

	eval.CheckField(validator.NotBlank(cpr.Description), "description", "this field is required")
	eval.CheckField(
		validator.MinChars(cpr.Description, 10) && validator.MaxChars(cpr.Description, 255),
		"description", "this field must be between 10 and 255 characters",
	)

	eval.CheckField(cpr.Baseprice > 0, "baseprice", "must be greater than 0")
	eval.CheckField(time.Until(cpr.AuctionEnd) >= minAuctionDuration, "auction_end", "must be at least two hours duration")

	return eval
}
