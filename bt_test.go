package bt_test

import (
	"log"
	"testing"

	"github.com/xupin/behavior-tree/bt"
)

func TestSkillTree(t *testing.T) {
	t.Run("Skill Ready", func(t *testing.T) {
		tree := bt.NewSequence(
			bt.NewXuLi(), // 会设置 can_use_skill = true
			bt.NewCondition(func(db bt.IBlackboard) bool {
				result := db.Get("can_use_skill")
				return result != nil && result.(bool)
			}),
			bt.NewRepeater(bt.NewJueZhao(), 3),
		)

		db := bt.NewBlackboard()
		result := tree.Exec(db)

		switch result {
		case bt.Success:
			log.Print("Skill Ready Test: ok")
		case bt.Failure:
			t.Error("Skill Ready Test: fail")
		}
	})

	t.Run("Skill Not Ready", func(t *testing.T) {
		tree := bt.NewSequence(
			// bt.NewXuLi() 不执行，can_use_skill 默认 nil
			bt.NewCondition(func(db bt.IBlackboard) bool {
				result := db.Get("can_use_skill")
				return result != nil && result.(bool)
			}),
			bt.NewRepeater(bt.NewJueZhao(), 3),
		)

		db := bt.NewBlackboard()
		result := tree.Exec(db)

		switch result {
		case bt.Success:
			t.Error("Skill Not Ready Test: fail")
		case bt.Failure:
			log.Print("Skill Not Ready Test: ok")
		}
	})
}
