package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-microservices/user-service/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	//创建与用户服务的连接
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	//根据连接创建grpc客户端
	client := userpb.NewUserServiceClient(conn)

	//分别查询一个存在的用户和一个不存在的用户
	for _, id := range []int64{1001, 9999} {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			3*time.Second,
		)

		resp, err := client.GetUser(ctx, &userpb.GetUserRequest{
			Id: id,
		})
		cancel()

		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				fmt.Printf(
					"查询用户%d失败: code=%s,message=%s\n",
					id,
					st.Code(),
					st.Message(),
				)
			} else {
				fmt.Printf("查询用户%d失败: %v\n", id, err)
			}
			continue
		}
		fmt.Printf(
			"查询成功: ID=%d, Name=%s, Email=%s\n",
			resp.GetId(),
			resp.GetName(),
			resp.GetEmail(),
		)
	}
}
