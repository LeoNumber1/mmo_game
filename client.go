package main

import (
	"fmt"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"io"
	"mmo_game/pb"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	dp := znet.NewDataPack()
	msg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx Client Test Message")))
	fmt.Println("---msg---err---", msg, "---", err, "---")
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}
	time.Sleep(5 * time.Second)

	return

	for {
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}

		//msg := &pb.BroadCast{}
		//err = proto.Unmarshal(headData, msg)
		//if err != nil {
		//	fmt.Println("-------unmarshal err :", err)
		//	//return
		//}
		//
		//fmt.Println("msg =", msg)

		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", msg.Data)
			if msg.Id == 1 {
				msgD := &pb.SyncPid{}
				err = proto.Unmarshal(msg.Data, msgD)
				if err != nil {
					fmt.Println("-------unmarshal err :", err)
					return
				}
				fmt.Println("msg.Data =", msgD)
			} else if msg.Id == 200 {
				msgD := &pb.BroadCast{}
				err = proto.Unmarshal(msg.Data, msgD)
				if err != nil {
					fmt.Println("-------unmarshal err :", err)
					return
				}
				fmt.Println("msg.Data =", msgD)
			}
		}

		time.Sleep(2 * time.Second)
	}
}
