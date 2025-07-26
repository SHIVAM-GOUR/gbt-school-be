package handlers

import (
	"net/http"

	"github.com/SHIVAM-GOUR/gbt-school-be/internal/models"
	"github.com/SHIVAM-GOUR/gbt-school-be/internal/utils"
	"gorm.io/gorm"
)

type InquiryHandler struct {
	DB *gorm.DB
}

func (i *InquiryHandler) GetInquiries(w http.ResponseWriter, r *http.Request) {
	var inquiries []models.Inquiry
	i.DB.Find(&inquiries)

	utils.SendJSONResponse(w, http.StatusOK, inquiries)
}
