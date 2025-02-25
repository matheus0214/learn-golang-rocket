package api

import (
	"gobid/internal/jsonutils"
	"gobid/internal/usecases/product"
	"net/http"

	"github.com/google/uuid"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJSON[*product.CreateProductReq](r)
	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	userID, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)
	if !ok {
		_ = jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}

	id, err := api.ProductService.CreateProduct(r.Context(), userID, data.ProductName, data.Description, data.Baseprice, data.AuctionEnd)
	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]string{
			"error": "failed to create product auction",
		})
		return
	}

	_ = jsonutils.EncodeJSON(w, r, http.StatusCreated, map[string]any{
		"message": "product auction created successfuly",
		"id":      id.String(),
	})
}
