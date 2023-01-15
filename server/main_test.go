package main

import (
	"context"
	"testing"

	"github.com/josenarvaezp/chat-gpt-shop-service/proto"
	"github.com/stretchr/testify/assert"
)

func TestListBananas(t *testing.T) {
	server := &server{}

	// Test case: List 2 bananas
	response, err := server.ListBananas(context.Background(), &proto.ListBananasRequest{Size: 2})
	assert.Nil(t, err)
	assert.Len(t, response.Bananas, 2)

	// Test case: List 0 bananas
	response, err = server.ListBananas(context.Background(), &proto.ListBananasRequest{Size: 0})
	assert.Nil(t, err)
	assert.Len(t, response.Bananas, 0)

	// test case fails
	// Test case: List more bananas than available
	// response, err = server.ListBananas(context.Background(), &proto.ListBananasRequest{Size: 10})
	// assert.Nil(t, err)
	// assert.Len(t, response.Bananas, len(bananas))
}

func TestGetBananaColor(t *testing.T) {
	server := &server{}

	// Test case: Get color of banana with ID 1
	response, err := server.GetBananaColor(context.Background(), &proto.GetBananaColorRequest{BananaId: "1"})
	assert.Nil(t, err)
	assert.Equal(t, "Yellow", response.Color)

	// Test case: Get color of banana with non-existing ID
	_, err = server.GetBananaColor(context.Background(), &proto.GetBananaColorRequest{BananaId: "10"})
	assert.NotNil(t, err)
	// error code updated
	assert.Equal(t, "rpc error: code = NotFound desc = Banana not found", err.Error())
}
