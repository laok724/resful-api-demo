package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/laok724/resful-api-demo/apps/host"
	"github.com/laok724/resful-api-demo/apps/host/impl"
)

var (
	// 定义对象满足该接口实例
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {

	// 初始化化全局 Logger，不设置默认打印，影响性能
	zap.DevelopmentSetup()
	// host service的具体实现
	service = impl.NewHostServiceImpl()
}
