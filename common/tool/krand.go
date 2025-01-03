package tool

import (
	"math/rand"
	"time"
)

//
const (
	KC_RAND_KIND_NUM   = 0 //数字
	KC_RAND_KIND_LOWER = 1 //小写字母
	KC_RAND_KIND_UPPER = 2 //大写字母
	KC_RAND_KIND_ALL   = 3 //数字、小写字母、大写字母
)

//随机数生成器
func Krand(kind int, length int) string {
    // 初始化 ikind 为传入的 kind 参数
    // 初始化 kinds 为一个二维数组，存储每种字符类型的范围和对应的 ASCII 基础值
    // 初始化 result 为一个长度为 length 的 byte 类型切片，用于存储生成的随机字符
    ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, length)

    // 判断是否生成混合字符
    // 如果 kind 大于 2 或小于 0，则表示要生成包含数字、小写字母和大写字母的混合字符串
    is_all := kind > 2 || kind < 0

    // 初始化随机数生成器，确保每次调用 Krand 函数时生成的随机数序列都是不同的
    rand.Seed(time.Now().UnixNano())

    // 循环生成每个随机字符
    for i := 0; i < length; i++ {
        // 如果生成混合字符，则每次循环随机选择一种字符类型
        if is_all {
            ikind = rand.Intn(3)
        }

        // 获取当前字符类型的范围和对应的 ASCII 基础值
        scope, base := kinds[ikind][0], kinds[ikind][1]

        // 生成一个范围在 [0, scope) 之间的随机整数，并加到 base 上得到一个有效的 ASCII 值
        // 将该 ASCII 值转换为 byte 类型，并存储到 result 切片中
        result[i] = byte(base + rand.Intn(scope))
    }

    // 将 result 切片转换为字符串并返回
    return string(result)
}
