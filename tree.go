/*
 * @Author: lgcgo
 * @Date: 2022-06-10 20:56:41
 * @Last Modified by: lgcgo
 * @Last Modified time: 2022-06-10 20:57:02
 */
package tree

import (
	"errors"
	"sort"
)

// Tree struct
type Tree struct {
	RootNode *Node          // Root node
	NodeList map[uint]*Node // Node list
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
func NewWithData(data []*TreeData) (*Tree, error) {
	var (
		insTree  = new(Tree)
		rootNode = NewNode(&TreeData{Title: "root"}) // Virtual root node
		nodeMap  = make(map[*Node]*Node, 0)          // Relationship chain
		nodeList = make(map[uint]*Node, 0)           // Node list
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

// Get all child key of node
func (t *Tree) GetAllChildKey(key uint) ([]uint, error) {
	var (
		keys     = make([]uint, 0)
		children = make(map[uint]*TreeData, 0)
		treeData *TreeData
		node     *Node
		err      error
	)
	if node, err = t.GetNode(key); err != nil {
		return nil, err
	}
	treeData = t.GetNodeTree(node)
	// Save self key
	keys = append(keys, treeData.Key)
	// Initialize queue
	for _, v := range treeData.Children {
		children[v.Key] = v
	}
	// level-traversal
	for {
		if len(children) == 0 {
			break
		}
		temp := children
		for k, v := range temp {
			keys = append(keys, v.Key)
			delete(children, k)

			if len(v.Children) > 0 {
				for _, vv := range v.Children {
					children[vv.Key] = vv
				}
			}
		}
	}

	return keys, nil
}

// Get the specified node tree data
func (t *Tree) GetNodeTree(n *Node) *TreeData {
	return n.Noder.(*TreeData)
}

// Gets the specified node
func (t *Tree) GetNode(key uint) (*Node, error) {
	n, ok := t.NodeList[key]
	if !ok {
		return nil, errors.New(`node not exist`)
	}
	return n, nil
}

// Get the number of nodes
func (t *Tree) CountNode() int {
	return len(t.NodeList)
}

// Get whole tree data
func (t *Tree) Tree() *TreeData {
	return t.GetNodeTree(t.RootNode)
}
