package token

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

//登录  路由业务
type TokenApi struct {
	znet.BaseRouter
}

func (*TokenApi) Handle(request ziface.IRequest) {
	data := request.GetData()
	fmt.Println("************", string(data))
	fmt.Println("--------------------------token---------------")
}
