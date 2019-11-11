package api

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"mmo_game/pb"
)

//登录  路由业务
type TokenApi struct {
	znet.BaseRouter
}

func (*TokenApi) Handle(request ziface.IRequest) {
	data := request.GetData()
	fmt.Println("************", string(data))
	msgD := &pb.Token{}
	err := proto.Unmarshal(data, msgD)
	if err != nil {
		fmt.Println("----unmarshal err-----", err)
		return
	}
	fmt.Println("解析后的数据：", msgD.Token)
	if msgD.Token != "8c9f08032f2f47d889b0ae99fdd8485e" { //模拟token验证
		fmt.Println("token验证失败")
		request.GetConnection().Stop()
		return
	}
	fmt.Println("--------------------------token---------------")
}
