package elog

import (
	"testing"

	"github.com/gin-gonic/gin"
)

//测试不同skip对获取文件影响
func TestGetFileInfo(t *testing.T) {
	for i := -1; i < 6; i++ {
		t.Logf("%d => %v", i, GetFileInfo(i))
	}
}

func TestGetAllInfo(t *testing.T) {
	t.Log(GetAllInfo(&gin.Context{}))
}
