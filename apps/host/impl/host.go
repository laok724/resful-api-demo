package impl

import (
	"context"

	"github.com/laok724/resful-api-demo/apps/host"
)

// 负责具体接口的实现

// 录入主机
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Debug("create host")
	i.l.Debugf("create host %s", ins.Name)
	return ins, nil
}

// 查询主机列表
func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

// 查询主机详情
func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

// 更新主机
func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

// 删除主机
func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
