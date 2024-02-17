package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Product struct {
	Id           int        `db:"id" json:"id"`
	Name         *string    `db:"name" json:"name" form:"name"`
	BasePrice    *int       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description  *string    `db:"description" json:"description" form:"description"`
	Image        *string    `db:"image" json:"image" form:"image"`
	IsBestSeller *bool      `db:"isBestSeller" json:"isBestSeller" form:"isBestSeller"`
	Discount     *float64   `db:"discount" json:"discount" form:"discount"`
	CreatedAt    *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt    *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProducts(search string, orderBy string, limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "products"
	WHERE "name" ILIKE $1
	ORDER BY "` + orderBy + `" ASC
	LIMIT $2
	OFFSET $3`
	sqlCount := `SELECT COUNT(*) FROM "products" WHERE "name" ILIKE $1`
	fmtSearch := fmt.Sprintf("%%%v%%", search)
	result := services.Info{}
	data := []Product{}
	db.Select(&data, sql, fmtSearch, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, fmtSearch)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProduct(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "id"=$1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductByName(name string) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "name" = $1`
	data := Product{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateProduct(data Product) (Product, error) {
	sql := `
	INSERT INTO "products" ("name","basePrice","description","image","discount","isBestSeller") VALUES
	(:name, :basePrice, :description, :image, :discount, :isBestSeller)
	RETURNING *`

	result := Product{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProduct(data Product) (Product, error) {
	sql := `
	UPDATE "products" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"basePrice"=COALESCE(NULLIF(:basePrice,0),"basePrice"),
	"description"=COALESCE(NULLIF(:description,''),"description"),
	"image"=COALESCE(NULLIF(:image,''),"image"),
	"isBestSeller"=COALESCE(:isBestSeller,false),
	"discount"=COALESCE(NULLIF(:discount,0.0),"discount"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := Product{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProduct(id int) (Product, error) {
	sql := `DELETE FROM "products" WHERE "id" = $1 RETURNING *`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}
