package Week02

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	Name string `json:"name"`
}

var db = sql.DB{}

// 我在意这个错误，那么就把堆栈信息加上自己的信息抛上去。
func GetNameWithMessage(name string) (*User, error) {
	return &User{}, errors.WithMessage(sql.ErrNoRows, "这是一个在意的错误")
}

// 我只是为了查询下有没有，我不在意这个错误。那么就吞下去苦果，返回nil。上面拿到nil知道没有，ok结束。
// 只是举例，实际开发中不会这么用。
func GetNameNotCare(name string) (*User, error) {
	// 执行sql查询,只是举个栗子
	_, err := db.Exec("select name from User where name=?", name)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &User{}, err
}
