package admin

import "time"

type User struct {
	ID        int64  //主键id
	Account   string //账号
	Name      string //姓名
	Password  string //密码
	Email     string //邮箱
	TokenTime int    //token有效期
	Status    int    //状态 010101
	Ctime     time.Duration
	Utime     time.Duration
}
