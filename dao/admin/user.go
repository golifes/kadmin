package admin

import (
	"context"
	"kadmin/model/admin"
)

type PsParams struct {
	Ps int `json:"ps,default=10" binding:"gte=10,lte=20"` //pageSize(每页显示多少)
	Pn int `json:"pn,default=1" binding:"gte=1,lte=100"`  //pageNum(页码)
}

//用户接口
type Users interface {
	//查找单个用户信息
	FindUserById(ctx context.Context, id int64) (u *admin.User, err error)

	//查询用户集合 查询条件
	FindUserList(ctx context.Context) (data interface{})
	//创建用户
	Create(ctx context.Context, u admin.User) (err error)
	//修改用户信息 #TODO 这里要id还是不要id参数
	Update(ctx context.Context, id int64, u admin.User) (err error)
	//更新密码
	UpdatePwd(ctx context.Context, id int64, pwd string) (err error)
	//删除用户
	Delete(ctx context.Context, id int64) (err error)
}
