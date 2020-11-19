package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type shell struct {
	S []string
}
type UserReq struct {
	Name       string           `json:"name"`
	Account    string           `json:"account"`
	Password   string           `json:"password"`
	Roles      map[string][]int `json:"roles"`       // 角色列表
	CreatorId  int64            `json:"creator_id"`  // 操作人的Id（谁添加的）
	Comment    string           `json:"comment"`     // 备注，添加信息
	Status     string           `json:"status"`      // 当前状态，是否审批通过
	VerifierId int64            `json:"verifier_id"` // 审核人员ID
}

type CalloutAgentConfig struct {
	AgentID int64  `json:"agent_id"`
	Sip     string `json:"sip"`
	Phone   string `json:"phone"`
}

type CalloutEntConfig struct {
	CalloutSwitch string               `json:"callout_switch"`
	Configs       []CalloutAgentConfig `json:"configs"`
}

type ChangePasswordReq struct {
	Account      string `json:"account"`
	ID           int64  `json:"id"`
	PrePassword  string `json:"pre_password"`
	NewPassword1 string `json:"new_password1"`
	NewPassword2 string `json:"new_password2"`
}

type UserSearchReq struct {
	Name    string `json:"name"`
	Account string `json:"account"`
	Product string `json:"product"` // 产品线，来鼓还是美洽
	Roles   []int  `json:"roles"`   // 角色列表
	Status  string `json:"status"`  // 当前状态，是否审批通过
}

type EntAuditAddReq struct {
	EntIds     []int64 `json:"ent_ids"`     // 企业ID
	Status     string  `json:"status"`      // 状态，黑，灰等,枚举类型
	Reason     []int64 `json:"reason"`      // 加入的原因，使用枚举类型
	Source     int64   `json:"source"`      // 封禁来源
	Comment    string  `json:"comment"`     // 备注
	AddExtra   bool    `json:"add_extra"`   // 是否把附加资源也加入
	AddComment bool    `json:"add_comment"` // 是否把附加资源也加入
	ForceAdd   bool    `json:"force_add"`   // 是否是强制添加
}
type EntAuditSearchReq struct {
	EntIds  []int64   `json:"ent_id"`  // 企业ID
	Status  string    `json:"status"`  // 状态，黑，灰等,枚举类型
	Reason  int64     `json:"reason"`  // 加入的原因，使用枚举类型
	Source  int64     `json:"source"`  // 封禁来源
	Product string    `json:"product"` // 产品线，来鼓还是美洽
	Comment string    `json:"comment"` // 备注
	Token   string    `json:"token"`   // 企业token
	Name    string    `json:"name"`    // 企业名字
	Phone   string    `json:"phone"`   // 电话
	Email   string    `json:"email"`
	StartAt time.Time `json:"start_at"` // 搜索开始时间
	EndAt   time.Time `json:"end_at"`   // 搜索结束时间

	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type Request struct {
	EntID         int64  `json:"ent_id"`
	Start         string `json:"start"`
	End           string `json:"end"`
	ExportContent bool   `json:"export_content"`
	ShowClue      bool   `json:"show_clue"`
	MaxRow        int    `json:"max_row"`
	WorkerCount   int    `json:"worker_count"`
	DataSource    int    `json:"data_source"`
	StartEntId    int64  `json:"start_ent_id"`
	Asyn          bool   `json:"asyn"`
}

func main() {

	u := Request{
		EntID:         1,
		Start:         "2006-01-02",
		End:           "2006-01-02",
		ExportContent: true,
		ShowClue:      true,
		MaxRow:        5000,
		WorkerCount:   1,
		DataSource:    2,
		StartEntId:    0,
		Asyn:          true,
	}
	bytes, _ := json.Marshal(u)

	fmt.Println(string(bytes))
}
