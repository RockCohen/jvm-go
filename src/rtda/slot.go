package rtda

import "rtda/heap"

// Slot
// 对数据结构槽的定义
/**
	num 	存放基本数据类型
    ref		存放引用数据类型
*/
type Slot struct {
	num int32
	ref *heap.Object
}
