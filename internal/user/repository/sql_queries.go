package repository

const (
	createUserQuery = `INSERT INTO users (first_name, last_name, email, password, role, avatar, delivery_address) 
		VALUES ($1, $2, $3, $4, $5, COALESCE(NULLIF($6, ''), null), $7)
		RETURNING user_id, first_name, last_name, email, password, avatar, created_at, updated_at, role, delivery_address`

	findByEmailQuery = `SELECT user_id, email, first_name, last_name, role, avatar, password, delivery_address, created_at, updated_at FROM users WHERE email = $1`

	findByIdQuery = `SELECT user_id, email, first_name, last_name, role, avatar, password, delivery_address, created_at, updated_at FROM users WHERE user_id = $1`

	findAllQuery = `SELECT user_id, email, first_name, last_name, role, avatar, password, delivery_address, created_at, updated_at FROM users LIMIT $1 OFFSET $2`

	updateByIdQuery = `UPDATE users SET first_name = $2, last_name = $3, email = $4, password = $5, role = $6, avatar = $7, delivery_address = $8 WHERE user_id = $1
		RETURNING user_id, first_name, last_name, email, password, avatar, delivery_address, created_at, updated_at, role`

	deleteByIdQuery = `DELETE FROM users WHERE user_id = $1`
)
