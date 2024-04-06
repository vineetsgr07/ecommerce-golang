package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vineetsrg07/ecommerce/golang/db"
	"github.com/vineetsrg07/ecommerce/golang/env"
	"github.com/vineetsrg07/ecommerce/golang/errors"
	"github.com/vineetsrg07/ecommerce/golang/server/write"
)

func CreateProduct(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
			return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	p := &db.ProductProduct{} // Assuming you have a Product struct similar to the Post struct
	err := decoder.Decode(p)
	if err != nil || p == nil {
			return write.Error(errors.NoJSONBody)
	}

	// Assuming there's no direct author for a product, but you might want to track who created it
	// This could be a "CreatedBy" field or similar
	// p.CreatedBy = user.ID

	return write.JSONorErr(env.DB().CreateProduct(r.Context(), db.CreateProductParams{
			Name:        p.Name,
			Description: p.Description,
			Status:      p.Status, // Assuming this is a valid product status like "available"
			ImageUrl:    p.ImageUrl,
	}))
}

func GetProduct(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
			return write.Error(errors.RouteUnauthorized)
	}

	id, err := getID(r) // Assuming getID extracts the product ID from the request path
	if err != nil {
			return write.Error(errors.RouteNotFound)
	}

	ProductProduct, err := env.DB().FindProductByID(r.Context(), id)
	if err != nil {
			if isNotFound(err) {
					return write.Error(errors.ProductNotFound) // Assuming you have a ProductNotFound error defined
			}
			return write.Error(err)
	}

	return write.JSON(ProductProduct)
}

func GetProducts(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
			return write.Error(errors.RouteUnauthorized)
	}

	// Optional: Retrieve query parameters from the request to filter products, e.g., by status
	// status := r.URL.Query().Get("status")

	// Assuming you have a method to find products possibly filtered by some criteria.
	// The example below doesn't apply filtering, but you could modify it to include it.
	products := env.DB().FindAllProducts(r.Context())

	return write.JSON(products)
}

func UpdateProduct(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
			return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	p := &db.ProductProduct{} // Assuming you have a Product struct similar to Post
	err := decoder.Decode(p)
	if err != nil || p == nil {
			return write.Error(errors.NoJSONBody)
	}

	// Optional: Check if the user has the rights to update the product
	// This could depend on your application's logic and user roles/permissions

	return write.JSONorErr(env.DB().UpdateProduct(r.Context(), db.UpdateProductParams{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Status:      p.Status, // Assuming this is a field in your Product struct
			ImageUrl:    p.ImageUrl, // Assuming products have images
	}))
}

func DeleteProduct(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
			return write.Error(errors.RouteUnauthorized)
	}

	id, err := getID(r) // Assuming getID extracts the product ID from the request path
	if err != nil {
			return write.Error(errors.RouteNotFound)
	}

	// Optional: Check if the user has the rights to delete the product
	// This could depend on your application's logic and user roles/permissions

	return write.SuccessOrErr(env.DB().DeleteProductByID(r.Context(), id))
}
