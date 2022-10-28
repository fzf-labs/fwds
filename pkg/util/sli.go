package util

import (
	"bytes"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Sli = newSli()

type sli struct {
}

func newSli() *sli {
	return &sli{}
}

// StringSliceReflectEqual 判断 string和slice 是否相等
// 因为使用了反射，所以效率较低，可以看benchmark结果
func (us *sli) StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

// StringSliceEqual 判断 string和slice 是否相等
// 使用了传统的遍历方式
func (us *sli) StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// reflect.DeepEqual的结果保持一致
	if (a == nil) != (b == nil) {
		return false
	}

	// bounds check 边界检查
	// 避免越界
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// SliceShuffle shuffle a slice
func (us *sli) SliceShuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// Uint64SliceReverse 对uint64 slice 反转
func (us *sli) Uint64SliceReverse(a []uint64) []uint64 {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

// IsInSlice 判断某一值是否在slice中
// 因为使用了反射，所以时间开销比较大，使用中根据实际情况进行选择
func (us *sli) IsInSlice(value interface{}, sli interface{}) bool {
	switch reflect.TypeOf(sli).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(sli)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// Uint64ShuffleSlice 对slice进行随机
func (us *sli) Uint64ShuffleSlice(a []uint64) []uint64 {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// see: https://yourbasic.org/golang/

// Uint64DeleteElemInSlice 从slice删除元素
// fast version, 会改变顺序
// i：slice的索引值
// s: slice
func (us *sli) Uint64DeleteElemInSlice(i int, s []uint64) []uint64 {
	if i < 0 || i > len(s)-1 {
		return s
	}
	// Remove the element at index i from s.
	s[i] = s[len(s)-1] // Copy last element to index i.
	s[len(s)-1] = 0    // Erase last element (write zero value).
	s = s[:len(s)-1]   // Truncate slice.

	return s
}

// Uint64DeleteElemInSliceWithOrder 从slice删除元素
// slow version, 保持原有顺序
// i：slice的索引值
// s: slice
func (us *sli) Uint64DeleteElemInSliceWithOrder(i int, s []uint64) []uint64 {
	if i < 0 || i > len(s)-1 {
		return s
	}
	// Remove the element at index i from s.
	copy(s[i:], s[i+1:]) // Shift s[i+1:] left one index.
	s[len(s)-1] = 0      // Erase last element (write zero value).
	s = s[:len(s)-1]     // Truncate slice.

	return s
}

// StringToInt64 string切片转int64切片
func (us *sli) StringToInt64(s []string) ([]int64, error) {
	int64s := make([]int64, len(s))
	for i, item := range s {
		parseInt, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, err
		}
		int64s[i] = parseInt
	}
	return int64s, nil
}

// StringToInt32 string切片转int32切片
func (us *sli) StringToInt32(s []string) ([]int32, error) {
	int32s := make([]int32, len(s))
	for i, item := range s {
		parseInt, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return nil, err
		}
		int32s[i] = int32(parseInt)
	}
	return int32s, nil
}

// ArrayDuplication String切片去重
func (us *sli) ArrayDuplication(arr []string) []string {
	var out []string
	tmp := make(map[string]byte)
	for _, v := range arr {
		tmpLen := len(tmp)
		tmp[v] = 0
		if len(tmp) != tmpLen {
			out = append(out, v)
		}
	}
	return out
}

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInt64 format int64 slice to string, eg: n1,n2,n3.
func (us *sli) JoinInt64(is []int64) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// SplitInt64 split string into int64 slice.
func (us *sli) SplitInt64(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

//切片快捷操作汇总：
//a := []int{1, 2, 3}
//b := []int{4, 5, 6}
//i := 1
//j := 3
//1.将切片 b 的元素追加到切片 a 之后：a = append(a, b...)
//2.删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
//3.切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
//4.为切片 a 扩展 j 个元素长度：a = append(a, make([]int, j)...)
//5.在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
//6.在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]int, j), a[i:]...)...)
//7.在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
//8.取出位于切片 a 最末尾的元素 x：x, a := a[len(a)-1:], a[:len(a)-1]

// DeleteSliceByPos 删除切片指定位置元素
func (us *sli) DeleteSliceByPos(slice interface{}, index int) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return slice, errors.New("not slice")
	}
	if v.Len() == 0 || index < 0 || index > v.Len()-1 {
		return slice, errors.New("index error")
	}
	return reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())).Interface(), nil
}

// InsertSliceByIndex 在指定位置插入元素
func (us *sli) InsertSliceByIndex(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return slice, errors.New("not slice")
	}
	if index < 0 || index > v.Len() || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return slice, errors.New("index error")
	}
	if index == v.Len() {
		return reflect.Append(v, reflect.ValueOf(value)).Interface(), nil
	}
	v = reflect.AppendSlice(v.Slice(0, index+1), v.Slice(index, v.Len()))
	v.Index(index).Set(reflect.ValueOf(value))
	return v.Interface(), nil
}

// UpdateSliceByIndex 更新指定位置元素
func (us *sli) UpdateSliceByIndex(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return slice, errors.New("not slice")
	}
	if index > v.Len()-1 || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return slice, errors.New("index error")
	}
	v.Index(index).Set(reflect.ValueOf(value))

	return v.Interface(), nil
}

// ContainsInterface 是否包含指定interface
func (us *sli) ContainsInterface(sl []interface{}, v interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsInt 是否包含指定int
func (us *sli) ContainsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsInt64 是否包含指定int64
func (us *sli) ContainsInt64(sl []int64, v int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsString 是否包含指定string
func (us *sli) ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// UniqueInt64 int64切片去重
func (us *sli) UniqueInt64(s []int64) []int64 {
	size := len(s)
	if size == 0 {
		return []int64{}
	}

	m := make(map[int64]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]int64, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// UniqueInt int切片去重
func (us *sli) UniqueInt(s []int) []int {
	size := len(s)
	if size == 0 {
		return []int{}
	}

	m := make(map[int]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]int, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// UniqueString string切片去重
func (us *sli) UniqueString(s []string) []string {
	size := len(s)
	if size == 0 {
		return []string{}
	}

	m := make(map[string]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]string, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// SumInt64 int64切片求和
func (us *sli) SumInt64(intSlice []int64) (sum int64) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// SumInt int切片求和
func (us *sli) SumInt(intSlice []int) (sum int) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// MaxInt64 int64切片中的最大值
func (us *sli) MaxInt64(int64Slice []int64) (max int64) {

	for _, v := range int64Slice {
		if v > max {
			max = v
		}
	}
	return
}

// MaxInt int切片中的最大值
func (us *sli) MaxInt(intSlice []int) (max int) {
	for _, v := range intSlice {
		if v > max {
			max = v
		}
	}
	return
}

// SumFloat64 float64切片求和
func (us *sli) SumFloat64(intSlice []float64) (sum float64) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// DescByField 根据切片中map的指定字段降序排序
func (us *sli) DescByField(list []map[string]interface{}, field string) {
	sort.Slice(list, func(i, j int) bool {
		return list[i][field].(int64) > list[j][field].(int64)
	})
}

// AscByField 根据切片中map的指定字段升序排序
func (us *sli) AscByField(list []map[string]interface{}, field string) {
	sort.Slice(list, func(i, j int) bool {
		return list[i][field].(int64) < list[j][field].(int64)
	})
}
