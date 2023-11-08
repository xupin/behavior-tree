package bt

import "fmt"

// 绝招
type JueZhao struct {
}

func NewJueZhao() *JueZhao {
	return &JueZhao{}
}

func (r *JueZhao) Exec(db IBlackboard) Status {
	fmt.Println("[绝招]done")
	db.Set("can_use_skill", false)
	return Success
}

// 蓄力
type XuLi struct {
}

func NewXuLi() *XuLi {
	return &XuLi{}
}

func (r *XuLi) Exec(db IBlackboard) Status {
	db.Set("can_use_skill", true)
	fmt.Println("[蓄力]done")
	return Success
}
