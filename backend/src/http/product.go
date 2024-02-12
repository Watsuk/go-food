package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Watsuk/go-food/src/auth"
	"github.com/Watsuk/go-food/src/permissions"
	"github.com/Watsuk/go-food/src/product"
	"github.com/go-chi/chi"
)

func CreateProductEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Restaurateur, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		var newProduct newProduct // On crée une variable de type Product
		err = json.NewDecoder(r.Body).Decode(&newProduct)

		if err != nil {
			log.Printf("Erreur de décodage JSON : %v", err)
			http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
			return
		}

		product, err := product.CreateProduct(db, newProduct.TruckID, newProduct.Name, newProduct.Label, newProduct.Description, newProduct.Price)

		if err != nil {
			log.Printf("Erreur lors la création du product : %v", err)
			http.Error(w, "Erreur lors la création du product", http.StatusBadRequest)
			return
		}

		jsonProduct, err := json.Marshal(product)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonProduct)
	}
}

func GetProductByIdEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.User, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		productIDString := chi.URLParam(r, "productID")

		productID, err := strconv.Atoi(productIDString)

		product, err := product.GetProduct(db, int64(productID))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonProduct, err := json.Marshal(product)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonProduct)
	}
}

func GetProductsByTruckEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		truckIDString := chi.URLParam(r, "truckID")

		truckID, err := strconv.Atoi(truckIDString)

		products, err := product.GetProductsByTruckID(db, int64(truckID))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonProducts, err := json.Marshal(products)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonProducts)
	}
}

func DeleteProductEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Restaurateur, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		productIDString := chi.URLParam(r, "productID")

		productID, err := strconv.Atoi(productIDString)

		err = product.DeleteProduct(db, int64(productID))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Product deleted"))
	}
}

type newProduct struct {
	TruckID     int    `json:"truckID"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
