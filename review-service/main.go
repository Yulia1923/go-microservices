package main

import (
	"fmt"
	"log"
)

func main() {
	userClient, conn, err := DialUserService()
	if err != nil {
		log.Fatalf("连接用户服务失败: %v", err)
	}
	defer CloseUserConnection(conn)

	service := NewReviewService(userClient)

	review, author, err := service.GetReviewWithAuthor(5001)
	if err != nil {
		log.Fatalf("查询评价失败: %v", err)
	}

	fmt.Printf("评价 ID：%d\n", review.ID)
	fmt.Printf("商品 ID：%d\n", review.ProductID)
	fmt.Printf("评价内容：%s\n", review.Content)
	fmt.Printf("作者：%s\n", author)
}
