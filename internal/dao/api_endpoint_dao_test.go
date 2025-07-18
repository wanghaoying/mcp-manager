package dao_test

import (
	"context"
	"mcp-manager/internal/dao"
	"mcp-manager/internal/model"
	_ "mcp-manager/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIEndpointDAO_Create_GetByID_Update_Delete_List(t *testing.T) {
	d := dao.NewAPIEndpointDAO(nil)
	ctx := context.Background()

	// Create
	endpoint := &model.APIEndpoint{Path: "/test", SwaggerID: 1}
	err := d.Create(ctx, endpoint)
	assert.NoError(t, err)
	assert.NotZero(t, endpoint.ID)

	// GetByID
	got, err := d.GetByID(ctx, endpoint.ID)
	assert.NoError(t, err)
	assert.Equal(t, endpoint.Path, got.Path)

	// Update
	endpoint.Path = "/updated"
	err = d.Update(ctx, endpoint)
	assert.NoError(t, err)
	got, err = d.GetByID(ctx, endpoint.ID)
	assert.NoError(t, err)
	assert.Equal(t, "/updated", got.Path)

	// List
	endpoints, err := d.List(ctx, endpoint.SwaggerID)
	assert.NoError(t, err)
	assert.Len(t, endpoints, 1)

	// Delete
	err = d.Delete(ctx, endpoint.ID)
	assert.NoError(t, err)
	got, err = d.GetByID(ctx, endpoint.ID)
	assert.Error(t, err)
	assert.Nil(t, got)
}
