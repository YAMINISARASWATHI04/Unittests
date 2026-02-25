package database

import "database/sql"

type Repository interface{
	FindCurrentDiscount() int
}

type DiscountRepository struct{
	db *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository{
	return &DiscountRepository{
		db:db,
	}
}

func (dc *DiscountRepository) FindCurrentDiscount()(discount int){
	sql:="select value from discount order by RAND() LIMIT 1"
	row:=dc.db.QueryRow(sql)

	row.Scan(&discount)
	return discount
}