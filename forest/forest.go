/**
 * @Author: znzAbl
 * @Date: 2021-08-06 14:42
 */
package forest

import "sort"

/**
 * 渲染森林结构数据
 * 兼容子节点先于父节点插入
 * parentID等于0 表示为根节点
 * parentID大于0 但是最后没有找到挂载节点将会被抛弃
 * 重复如果两条数据的id和parentID重复时将会被抛后来者
 */
func NewForest() *forest {
	f := new(forest)
	f.auxiliaryMap = make(map[int]*item)
	f.itemList = make(ItemSort, 0)
	return f
}

// 装载森林结构数据
func (f *forest) PushSonInTree(id, parentID, order int, name string, data interface{}) {
	if _, ok := f.auxiliaryMap[id]; !ok {
		f.auxiliaryMap[id] = &item{
			ID:       id,
			Name:     name,
			Order:    order,
			ParentID: parentID,
			Data:     data,
		}
	}
	if f.auxiliaryMap[id].IsClassification {
		return
	}
	if parentID == 0 {
		return
	}
	k, ok := f.auxiliaryMap[parentID]
	if !ok || k == f.auxiliaryMap[id] {
		return
	}
	f.auxiliaryMap[id].IsClassification = true
	if len(k.Child) == 0 {
		k.Child = make(ItemSort, 0)
		k.Child = append(k.Child, f.auxiliaryMap[id])
	} else {
		k.Child = append(k.Child, f.auxiliaryMap[id])
		sort.Sort(k.Child)
	}
	return
}

// 刷新
func (f *forest) Refresh() *forest {
	f.itemList = make(ItemSort, 0)
	return f
}

// 装载所有根节点
func (f *forest) GetTree() ItemSort {
	if len(f.itemList) > 0 {
		return f.itemList
	}
	for _, v := range f.auxiliaryMap {
		if v.IsClassification {
			// 已挂载的一定不是根节点
			continue
		}
		if v.ParentID > 0 {
			// 不是根节点 尝试挂载
			f.PushSonInTree(v.ID, v.ParentID, v.Order, v.Name, v.Data)
			continue
		}
		f.itemList = append(f.itemList, v)
	}
	if len(f.itemList) > 1 {
		sort.Sort(f.itemList)
	}
	return f.itemList
}

