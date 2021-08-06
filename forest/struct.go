/**
 * @Author: znzAbl
 * @Date: 2021-08-06 14:42
 */
package forest

type item struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Order            int         `json:"order,omitempty"`
	ParentID         int         `json:"parent_id"`
	IsClassification bool        `json:"-"`
	Data             interface{} `json:"data,omitempty"`
	Child            ItemSort    `json:"child,omitempty"`
}

type ItemSort []*item

func (s ItemSort) Len() int { return len(s) }

func (s ItemSort) Less(i, j int) bool { return s[i].Order < s[j].Order }

func (s ItemSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type forest struct {
	itemList     ItemSort
	auxiliaryMap map[int]*item
}
