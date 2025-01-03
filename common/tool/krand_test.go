package tool

import (
	"testing"
)

// 测试生成不同类型的随机字符串
func TestKrand(t *testing.T) {
	// 测试生成数字字符串
	t.Run("生成数字字符串", func(t *testing.T) {
		result := Krand(KC_RAND_KIND_NUM, 5)
		if len(result) != 5 {
			t.Errorf("期望长度为 5，实际长度为 %d", len(result))
		}
	})

	// 测试生成小写字母字符串
	t.Run("生成小写字母字符串", func(t *testing.T) {
		result := Krand(KC_RAND_KIND_LOWER, 5)
		if len(result) != 5 {
			t.Errorf("期望长度为 5，实际长度为 %d", len(result))
		}
	})

	// 测试生成大写字母字符串
	t.Run("生成大写字母字符串", func(t *testing.T) {
		result := Krand(KC_RAND_KIND_UPPER, 5)
		if len(result) != 5 {
			t.Errorf("期望长度为 5，实际长度为 %d", len(result))
		}
	})

	// 测试生成混合字符字符串
	t.Run("生成混合字符字符串", func(t *testing.T) {
		result := Krand(KC_RAND_KIND_ALL, 5)
		if len(result) != 5 {
			t.Errorf("期望长度为 5，实际长度为 %d", len(result))
		}
	})

	// 测试生成空字符串（长度为 0）
	t.Run("生成空字符串", func(t *testing.T) {
		result := Krand(KC_RAND_KIND_ALL, 0)
		if len(result) != 0 {
			t.Errorf("期望长度为 0，实际长度为 %d", len(result))
		}
	})

	// 测试生成长度为 1 的字符串
	t.Run("生成长度为 1 的字符串", func(t *testing.T) {
		result := Krand(KC_RAND_KIND_ALL, 1)
		if len(result) != 1 {
			t.Errorf("期望长度为 1，实际长度为 %d", len(result))
		}
	})

	// 测试非法 kind 参数（小于 0）
	t.Run("非法 kind 参数小于 0", func(t *testing.T) {
		result := Krand(-1, 5)
		if len(result) != 5 {
			t.Errorf("期望长度为 5，实际长度为 %d", len(result))
		}
	})

	// 测试非法 kind 参数（大于 3）
	t.Run("非法 kind 参数大于 3", func(t *testing.T) {
		result := Krand(4, 5)
		if len(result) != 5 {
			t.Errorf("期望长度为 5，实际长度为 %d", len(result))
		}
	})
}
