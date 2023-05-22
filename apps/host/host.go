package host

import (
	"context"
)

// host app service 的接口定义
type Service interface {
	// 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// 查询主机详情
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	// 主机更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// 删除主机,删除为啥要返回，比如前端需要打印当前删除的主机ip 等信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

// HOST模型定义
type Host struct {
	*Resource // 公共属性
	*Describe // 独有属性
}

const (
	PRIVATE_IDC Vendor = iota
	ALIYUN
	TXYUN
)

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

type Vendor int

// 资源公共属性
type Resource struct {
	Id          string            `json:"id"  validate:"required"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                      // 厂商
	Region      string            `json:"region"  validate:"required"` // 地域
	CreateAt    int64             `json:"create_at"`                   // 创建时间
	ExpireAt    int64             `json:"expire_at"`                   // 过期时间
	Type        string            `json:"type"  validate:"required"`   // 规格
	Name        string            `json:"name"  validate:"required"`   // 名称
	Description string            `json:"description"`                 // 描述
	Status      string            `json:"status"`                      // 服务商中的状态
	Tags        map[string]string `json:"tags"`                        // 标签
	UpdateAt    int64             `json:"update_at"`                   // 更新时间
	SyncAt      int64             `json:"sync_at"`                     // 同步时间
	Account     string            `json:"account"`                     // 资源的所属账号
	PublicIP    string            `json:"public_ip"`                   // 公网IP
	PrivateIP   string            `json:"private_ip"`                  // 内网IP
}

// 资源独有属性
type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}

type QueryHostRequest struct{}

type DescribeHostRequest struct{}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct{}
