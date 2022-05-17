/*
 * @Author: Jimmy.liu
 * @Date: 2022-05-13 14:58:17
 * @Last Modified by: Jimmy.liu
 * @Last Modified time: 2022-05-13 15:37:17
 */
package tree

import (
	"errors"
	"sort"
)

// Tree struct
type tree struct {
	RootNode *node          // Root node
	NodeList map[uint]*node // Node list
}

// Data structure(Refer to antd tree requirements)
type TreeData struct {
	Title     string      `json:"title"`
	Key       uint        `json:"key"`
	ParentKey uint        `json:"parent_key"`
	Value     string      `json:"value"`
	Weight    int         `json:"weight"`
	Disabled  bool        `json:"disabled"`
	Children  []*TreeData `json:"children"`
}

// Implement Inoder abstract component
func (d TreeData) CalSortValue() int {
	return d.Weight
}

// Create with treeData
func NewWithData(data []*TreeData) (*tree, error) {
	var (
		insTree  = new(tree)
		rootNode = NewNode(&TreeData{Title: "root"}) // Virtual root node
		nodeMap  = make(map[*node]*node, 0)          // Relationship chain
		nodeList = make(map[uint]*node, 0)           // Node list
	)
	// Set node list
	nodeList[0] = rootNode
	for _, item := range data {
		nodeList[item.Key] = NewNode(item)
	}
	// Set relationship list
	for _, v := range nodeList {
		if v == rootNode {
			nodeMap[v] = nil
			continue
		}
		pk := v.Noder.(*TreeData).ParentKey
		// Detect parent node
		if _, ok := nodeList[pk]; !ok {
			return nil, errors.New("missing parent data")
		}
		nodeMap[v] = nodeList[pk]
	}
	// Set child nodes
	for _, n := range nodeList {
		var (
			nodes = make(nodes, 0)
			datas = make([]*TreeData, 0)
		)
		// Get child nodes
		for k, v := range nodeMap {
			if n == v {
				nodes = append(nodes, k)
			}
		}
		sort.Sort(nodes)
		// data conversion
		for _, cn := range nodes {
			datas = append(datas, cn.Noder.(*TreeData))
		}
		n.Noder.(*TreeData).Children = datas
	}
	insTree.RootNode = rootNode
	insTree.NodeList = nodeList
	return insTree, nil
}

// Get the specified node tree data
func (t *tree) GetNodeTree(n *node) *TreeData {
	return n.Noder.(*TreeData)
}

// Gets the specified node
func (t *tree) GetNode(key uint) (*node, error) {
	n, ok := t.NodeList[key]
	if !ok {
		return nil, errors.New(`node not exist`)
	}
	return n, nil
}

// Get the number of nodes
func (t *tree) CountNode() int {
	return len(t.NodeList)
}

// Get whole tree data
func (t *tree) Tree() *TreeData {
	return t.GetNodeTree(t.RootNode)
}
