package model

func AddCartItem(userID, bookID uint64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "insert ignore into cart(user_id, book_id) values(?, ?)"
	result, err := tx.Exec(query, userID, bookID)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		query = `update cart
		set quantity = quantity + 1
		where user_id = ? and book_id = ?`
		tx.Exec(query, userID, bookID)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
