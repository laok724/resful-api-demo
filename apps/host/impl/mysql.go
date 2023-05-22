package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/laok724/resful-api-demo/apps/host"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子 logger
		// 封装 zap 使其满足 logger 接口
		// 为什么要封装：
		//   1.Logger全局实例
		//   2.Logger Level 动态调整，Logrus 不支持 Level 共同调整
		//   3.加入日志轮转功能的集合
		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}
