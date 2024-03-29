package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type ProductSize struct {
	Id              int        `db:"id" json:"id"`
	Size            *string    `db:"size" json:"size" form:"size"`
	AdditionalPrice *int       `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductSize(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productSize" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productSize"`
	result := services.Info{}
	data := []ProductSize{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductSize(id int) (ProductSize, error) {
	sql := `SELECT * FROM "productSize" WHERE "id"=$1`
	data := ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductSizeBySize(size string) (ProductSize, error) {
	sql := `SELECT * FROM "productSize" WHERE "size"=$1`
	data := ProductSize{}
	err := db.Get(&data, sql, size)
	return data, err
}

func FindProductSizeById(id int) (ProductSize, error) {
	sql := `SELECT "size" FROM "productSize" WHERE "id" = $1`
	data := ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductSize(data ProductSize) (ProductSize, error) {
	sql := `
	INSERT INTO "productSize" ("size","additionalPrice") VALUES
	(:size, :additionalPrice)
	RETURNING *`

	result := ProductSize{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProductSize(data ProductSize) (ProductSize, error) {
	sql := `
	UPDATE "productSize" SET
	"size"=COALESCE(NULLIF(:size,''),"size"),
	"additionalPrice"=COALESCE(:additionalPrice,"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := ProductSize{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProductSize(id int) (ProductSize, error) {
	sql := `DELETE FROM "productSize" WHERE "id" = $1 RETURNING *`
	data := ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}
