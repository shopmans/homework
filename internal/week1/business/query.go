package business

import (
	// 标准包
	"database/sql"

	// 第三方包

	// 内部包
	user "gotraining3/internal/week1/business/models"
)

func Query(db *sql.DB, id int) (*user.User, error) {
	u := user.User{}
	row := db.QueryRow("select `id`, `name` from `users` where `id` = ?", id)
	err := row.Scan(&u.ID, &u.Name)
	if sql.ErrNoRows == err {
		return nil, nil
	}

	return &u, nil
}
