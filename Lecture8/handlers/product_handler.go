package handlers

import (
	"Lecture8/db"
	"Lecture8/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var _, err = db.DB.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	handleError(c, err)

	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, price FROM products")
	handleError(c, err)
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		handleError(c, err)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	err := db.DB.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	handleError(c, err)

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", updatedProduct.Name, updatedProduct.Price, id)
	handleError(c, err)

	c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM products WHERE id = $1", id)
	handleError(c, err)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
