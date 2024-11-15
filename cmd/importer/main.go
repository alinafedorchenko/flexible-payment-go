package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// move to config
	connStr := "user=myuser dbname=mydb password=newpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Perform a sample query
	rows, err := db.Query("SELECT id, email FROM merchants")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var id string
		var email string
		//var live_on string
		//var disbursement_frequency string
		//var minimum_monthly_fee_cents string
		//var created_at string
		if err := rows.Scan(&id, &email); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %s, Email: %s\n", id, email)
	}
}

//schema
//CREATE TYPE disbursement_frequency_type AS ENUM ('daily', 'weekly');
//

//CREATE TABLE merchants (
//id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
//email VARCHAR(255) NOT NULL,
//live_on DATE NOT NULL,
//disbursement_frequency disbursement_frequency_type NOT NULL DEFAULT 'daily',
//minimum_monthly_fee_cents BIGINT NOT NULL,
//created_at TIMESTAMP DEFAULT NOW()
//);
//CREATE TABLE disbursements (
//id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
//amount_cents BIGINT NOT NULL,
//fee_cents BIGINT NOT NULL,
//monthly_fee_cents BIGINT NOT NULL DEFAULT 0,
//merchant_id UUID NOT NULL,
//created_at TIMESTAMP DEFAULT NOW(),
//updated_at TIMESTAMP DEFAULT NOW(),
//FOREIGN KEY (merchant_id) REFERENCES merchants(id)
//);
//
//CREATE TABLE orders (
//id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
//amount_cents BIGINT NOT NULL,
//disbursement_id UUID,
//merchant_id UUID NOT NULL,
//created_at TIMESTAMP DEFAULT NOW() NOT NULL,
//updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
//FOREIGN KEY (disbursement_id) REFERENCES disbursements(id),
//FOREIGN KEY (merchant_id) REFERENCES merchants(id)
//);
//CREATE INDEX index_orders_on_disbursement_id ON orders (disbursement_id);
//CREATE INDEX index_orders_on_merchant_id ON orders (merchant_id);
//CREATE INDEX index_disbursements_on_merchant_id ON disbursements (merchant_id);
