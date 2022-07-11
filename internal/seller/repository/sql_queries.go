package repository

const (
	createSellerQuery = `INSERT INTO sellers (first_name, last_name, email, password, avatar) 
		VALUES ($1, $2, $3, $4, $5, COALESCE(NULLIF($6, ''), null)) 
		RETURNING seller_id, first_name, last_name, email, password, avatar, created_at, updated_at`

	findByEmailQuery = `SELECT seller_id, email, first_name, last_name, avatar, password, created_at, updated_at FROM sellers WHERE email = $1`

	findByIDQuery = `SELECT seller_id, email, first_name, last_name, avatar, password, created_at, updated_at FROM sellers WHERE seller_id = $1`

	findAllQuery = `SELECT seller_id, email, first_name, last_name, avatar, password, created_at, updated_at FROM sellers LIMIT $1 OFFSET $2`

	updateByIDQuery = `UPDATE sellers SET first_name = $2, last_name = $3, email = $4, password = $5, avatar = $6 WHERE seller_id = $1
		RETURNING seller_id, first_name, last_name, email, password, avatar, created_at, updated_at`

	deleteByIDQuery = `DELETE FROM sellers WHERE seller_id = $1`
)
