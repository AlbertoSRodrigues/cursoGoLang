package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	selectedCampaign, err := h.CampaignService.GetBy(id)
	return selectedCampaign, 200, err

}
