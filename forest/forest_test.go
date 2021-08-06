/**
 * @Author: znzAbl
 * @Date: 2021-08-06 15:08
 */
package forest

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewForest(t *testing.T) {
	forest := NewForest()
	forest.PushSonInTree(15,14,1,"第三个根菜单的子菜单-的子菜单", "备注一下8")
	forest.PushSonInTree(12,1,2,"第一个根菜单的子菜单2", "备注一下5")

	forest.PushSonInTree(1,0,1,"第一个根菜单", "备注一下1")
	forest.PushSonInTree(2,0,2,"第二个根菜单", "备注一下2")
	forest.PushSonInTree(3,0,3,"第三个根菜单", "备注一下3")

	forest.PushSonInTree(11,1,1,"第一个根菜单的子菜单", "备注一下4")
	forest.PushSonInTree(13,2,1,"第二个根菜单的子菜单", "备注一下6")
	forest.PushSonInTree(14,3,1,"第三个根菜单的子菜单", "备注一下7")

	str, _ := json.Marshal(forest.GetTree())
	fmt.Println(string(str))

	forest.PushSonInTree(16,15,1,"第三个根菜单的子菜单-的子菜单-的子菜单", "备注一下9")
	str, _ = json.Marshal(forest.Refresh().GetTree())
	fmt.Println(string(str))
}
