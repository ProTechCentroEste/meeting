// server/server_test.go
package main

import (
	"context"
	"testing"

	pb "github.com/ProTechCentroEste/meeting/proto"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	s := &server{}
	req := &pb.LoginRequest{
		Auth_provider: "Google",
		Token:         "sample_token",
	}
	res, err := s.Login(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "12345", res.User_id)
	assert.True(t, res.Success)
}
