package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry
/**
CompositeEntry由更小的Entry组成，正好可以表示
成[]Entry。在Go语言中，数组属于比较低层的数据结构，很少直接
使用。大部分情况下，使用更便利的slice类型。
*/
type CompositeEntry []Entry

func newCompositeEntry(paths string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(paths, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
