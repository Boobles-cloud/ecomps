package handlers

import (
	"net/http"
)

func (p *ProductHandler) HandleGettingProductById(w http.ResponseWriter, r *http.Request) {}

func (p *ProductHandler) HandleGettingAllProductsByTenantId(w http.ResponseWriter, r *http.Request) {}

func (p *ProductHandler) HandleGettingPictureById(w http.ResponseWriter, r *http.Request) {}
