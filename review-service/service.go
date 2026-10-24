package main

import (
	"fmt"

	"google.golang.org/grpc"
)

// ReviewService 负责管理评价数据，并查询评价作者。
type ReviewService struct {
	reviews    map[int64]Review
	userClient *UserClient
}

// NewReviewService 创建带有示例数据的评价服务。
func NewReviewService(userClient *UserClient) *ReviewService {
	return &ReviewService{
		reviews: map[int64]Review{
			5001: {
				ID:        5001,
				UserID:    1001,
				ProductID: 2001,
				Content:   "商品质量很好",
			},
			5002: {
				ID:        5002,
				UserID:    1002,
				ProductID: 2002,
				Content:   "物流速度很快",
			},
		},
		userClient: userClient,
	}
}

// GetReview 根据评价 ID 查询评价。
func (s *ReviewService) GetReview(id int64) (Review, bool) {
	review, exists := s.reviews[id]
	return review, exists
}

// GetReviewWithAuthor 查询评价及其作者信息。
func (s *ReviewService) GetReviewWithAuthor(id int64) (Review, string, error) {
	review, exists := s.GetReview(id)
	if !exists {
		return Review{}, "", fmt.Errorf("review %d not found", id)
	}

	user, err := s.userClient.GetUser(review.UserID)
	if err != nil {
		return Review{}, "", fmt.Errorf("get review author: %w", err)
	}

	return review, user.GetName(), nil
}

// CloseUserConnection 关闭用户服务的 gRPC 连接。
func CloseUserConnection(conn *grpc.ClientConn) {
	if conn != nil {
		_ = conn.Close()
	}
}
