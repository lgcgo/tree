/*
 * @Author: Jimmy.liu
 * @Date: 2022-05-13 14:58:48
 * @Last Modified by: Jimmy.liu
 * @Last Modified time: 2022-05-13 14:59:24
 */
package tree

// Abstract interface
type INoder interface {
	CalSortValue() int
}

// Decorator
type node struct {
	Noder INoder
}

// Implement sorting algorithm
type nodes []*node

// Sort by sort value
func (ns nodes) Len() int { return len(ns) }
func (ns nodes) Less(i, j int) bool {
	return ns[i].Noder.CalSortValue() > ns[j].Noder.CalSortValue()
}
func (ns nodes) Swap(i, j int) { ns[i], ns[j] = ns[j], ns[i] }
func (ns *nodes) Push(x interface{}) {
	*ns = append(*ns, x.(*node))
}
func (ns *nodes) Pop() interface{} {
	old := *ns
	n := len(old)
	x := old[n-1]
	*ns = old[0 : n-1]
	return x
}

// Create node
func NewNode(n INoder) *node {
	return &node{Noder: n}
}
