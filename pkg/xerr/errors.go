package xerr

import (
	"github.com/zeromicro/x/errors"
)
func New(code int,msg string) error {
	return errors.New(code,msg)
}
func NewMsg(msg string) error {
	return errors.New(SERVER_COMMON_ERR,msg)
}
func NewDBErr() error {
	return errors.New(DB_ERR,ErrMsg(DB_ERR))
}
func NewInternalErr() error {
	return errors.New(SERVER_COMMON_ERR,ErrMsg(SERVER_COMMON_ERR))
}
