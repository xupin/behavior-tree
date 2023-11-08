package main

import (
	"fmt"

	"github.com/xupin/behavior-tree/bt"
)

func main() {
	// 定义根节点
	t := bt.NewSequence(
		bt.NewXuLi(),
		bt.NewCondition(func(db bt.IBlackboard) bool {
			result := db.Get("can_use_skill")
			if result == nil || !result.(bool) {
				return false
			}
			return true
		}),
		bt.NewRepeater(bt.NewJueZhao(), 3),
	)
	db := bt.NewBlackboard()
	result := t.Exec(db)
	fmt.Print("结果: ")
	switch result {
	case bt.Success:
		fmt.Println("成功")
	case bt.Failure:
		fmt.Println("失败")
	}
}
