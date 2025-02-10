package export

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/yahyaammar-dev/pacebe/types"
	"gorm.io/gorm"
)

type Export struct {
	db *gorm.DB
}

func NewExporter(db *gorm.DB) *Export {
	return &Export{
		db: db,
	}
}

func (e *Export) ExportCSV(w http.ResponseWriter) {
	var products []types.Product
	result := e.db.Find(&products) // Fetch all products from the database
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Set headers for CSV download
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=products.csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write CSV header
	header := []string{"ID", "Name", "Description", "Price", "Stock", "Category", "ImageURL", "CreatedAt", "UpdatedAt"}
	if err := writer.Write(header); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write each product as a CSV row
	for _, product := range products {
		record := []string{
			fmt.Sprintf("%d", product.ID),
			product.Name,
			product.Description,
			fmt.Sprintf("%.2f", product.Price),
			fmt.Sprintf("%d", product.Stock),
			product.Category,
			product.ImageURL,
			product.CreatedAt.Format(time.RFC3339),
			product.UpdatedAt.Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
