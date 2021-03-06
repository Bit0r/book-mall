package model

import "strings"

func AddOrder(userID uint64, addressID uint, booksID []uint64) (err error) {
	if len(booksID) == 0 {
		return nil
	}

	txn, _ := db.Begin()

	query := `insert into orders(total, user_id, address_id, status)
	values(0, ?, ?, ?)`
	result, err := txn.Exec(query, userID, addressID, "unpaid")
	if err != nil {
		return err
	}
	orderID, _ := result.LastInsertId()

	holders := strings.Repeat(",?", len(booksID))[1:]
	query = `insert into order_item(order_id, book_id, price, quantity)
	select ?, book_id, price, quantity
	from book, cart
	where id = book_id and user_id = ? and book_id in (` + holders + ")"

	args := make([]interface{}, len(booksID)+2)
	args[0], args[1] = orderID, userID
	for i, v := range booksID {
		args[i+2] = v
	}
	_, err = txn.Exec(query, args...)
	if err != nil {
		return err
	}

	query = `update orders
	set total = (select sum(subtotal)
			from order_item
			where order_id = ?)
	where id = ?`
	_, err = txn.Exec(query, orderID, orderID)
	if err != nil {
		return err
	}

	query = `delete from cart
	where user_id = ? and book_id in (` + holders + ")"
	_, err = txn.Exec(query, args[1:]...)
	if err != nil {
		return err
	}

	txn.Commit()
	return
}
