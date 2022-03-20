package controller

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	gin "github.com/gin-gonic/gin"
	model "github.com/huyanhvn/quick-rest-go/model"
	assert "github.com/stretchr/testify/assert"
)

func setup(tb testing.TB) (func(tb testing.TB), *httptest.ResponseRecorder, *gin.Context, *Controller) {
	log.Println("Setup for unit tests.")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	c := NewController()
	model.ConnectDatabase("unit_test.db")

	return func(tb testing.TB) {
		log.Println("Tear down setup.")
	}, w, ctx, c
}

func TestController_GetAllItems(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{}
			c.GetAllItems(tt.args.ctx)
		})
	}
}

func TestController_GetItemByID(t *testing.T) {
	// setup & tear down
	teardown, w, ctx, c := setup(t)
	defer teardown(t)

	// test definitions
	tests := []struct {
		name     string
		expected gin.H
	}{
		// Example: a non-existent item should return 404
		{
			name:     "non-existent item",
			expected: gin.H{"code": float64(404), "message": "record not found"},
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
			c.GetItemByID(ctx)
			var got gin.H
			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestController_CreateItem(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{}
			c.CreateItem(tt.args.ctx)
		})
	}
}

func TestController_UpdateItem(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{}
			c.UpdateItem(tt.args.ctx)
		})
	}
}

func TestController_DeleteItem(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{}
			c.DeleteItem(tt.args.ctx)
		})
	}
}
