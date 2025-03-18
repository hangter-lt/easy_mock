package utils

import (
	"reflect"
)

// 计算两个map[string]any的包含关系
func IsAInB(a, b map[string]any) bool {

	for kA, vA := range a {
		// 如果b不存在该key,或不满足包含关系,返回false
		vB, ok := b[kA]
		if !ok || !deepContains(vA, vB) {
			return false
		}
	}

	return true
}

// 深度判断 valueA 是否被 valueB 包含
func deepContains(valueA, valueB any) bool {
	// 如果两者类型不同，直接返回 false
	if reflect.TypeOf(valueA) != reflect.TypeOf(valueB) {
		return false
	}

	switch vA := valueA.(type) {
	case map[string]any:
		vB, ok := valueB.(map[string]any)
		if !ok || len(vA) > len(vB) {
			return false
		}
		// 递归检查 mapA z的所有键值对是否都在 mapB 中
		for key, valA := range vA {
			valB, exists := vB[key]
			if !exists || !deepContains(valA, valB) {
				return false
			}
		}
		return true
	case []any:
		vB, ok := valueB.([]any)
		if !ok || len(vA) > len(vB) {
			return false
		}
		// 检查 sliceA 的所有元素是否都在 sliceB 中
		for _, itemA := range vA {

			found := false
			for _, itemB := range vB {
				if deepContains(itemA, itemB) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	default:
		// 对于基本类型，直接使用 reflect.DeepEqual 比较
		return reflect.DeepEqual(valueA, valueB)
	}
}
