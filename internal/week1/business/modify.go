package business

import (
	// 标准包
	"database/sql"

	// 第三方包
	"github.com/pkg/errors"

	// 内部包
	user "gotraining3/internal/week1/business/models"
)

func Modify(db *sql.DB, modifyUser *user.User) (*user.User, error) {
	u := user.User{}
	row := db.QueryRow("select `id`, `name` from `users` where `id` = ?", modifyUser.ID)
	err := row.Scan(&u.ID, &u.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(err, "修改用户失败，用户不存在")
	}

	// 修改用户省略...

	return &u, nil
}
