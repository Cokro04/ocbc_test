package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	SKU             string  `json:"sku"`
	ProductName     string  `json:"productName"`
	QuantityInStock int     `json:"quantityInStock"`
	Price           float64 `json:"price"`
	Category        string  `json:"category"`
}

// Allowed categories
var allowedCategories = map[string]bool{
	"Electronics": true,
	"Books":       true,
	"Apparel":     true,
	"Home Goods":  true,
}

// Validator
func ValidateProduct(p Product) []string {
	var errs []string

	// 1. SKU mandatory
	if p.SKU == "" {
		errs = append(errs, "The sku is a mandatory field")
	} else {
		// 2. SKU format
		match, _ := regexp.MatchString(`^SKU-\d{8}$`, p.SKU)
		if !match {
			errs = append(errs, "The sku must be in the format SKU-XXXXXXXX")
		}
	}

	// 3. productName mandatory
	if p.ProductName == "" {
		errs = append(errs, "The productName is a mandatory field")
	}

	// 4. quantityInStock cannot be negative
	if p.QuantityInStock < 0 {
		errs = append(errs, "The quantityInStock cannot be negative")
	}

	// 5. price > 0
	if p.Price <= 0 {
		errs = append(errs, "The price must be greater than zero")
	}

	// 6. category mandatory
	if p.Category == "" {
		errs = append(errs, "The category is a mandatory field")
	} else {
		// 7. category must be valid
		if !allowedCategories[p.Category] {
			errs = append(errs, "Invalid product category")
		}
	}

	return errs
}

// Handler
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"Invalid JSON format"})
		return
	}

	// Validate
	errors := ValidateProduct(p)

	// Return 400 if any validation errors
	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	// Otherwise return success
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("Starting server on :8080")
	r := chi.NewRouter()

	r.Post("/products", ProductHandler)

	http.ListenAndServe(":8080", r)
}
