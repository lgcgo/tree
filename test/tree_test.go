package tree

import (
	"encoding/json"
	"testing"

	"github.com/lgcgo/tree"
)

var data = []*tree.TreeData{
	{
		Title:     "superAdmin",
		Key:       1,
		ParentKey: 0,
		Value:     "superAdmin",
		Weight:    0,
	},
	{
		Title:     "sub1Admin",
		Key:       2,
		ParentKey: 1,
		Value:     "sub1Admin",
		Weight:    0,
	},
	{
		Title:     "sub2Admin",
		Key:       3,
		ParentKey: 1,
		Value:     "sub2Admin",
		Weight:    0,
	},
	{
		Title:     "sub3Admin",
		Key:       4,
		ParentKey: 2,
		Value:     "sub3Admin",
		Weight:    0,
	},
	{
		Title:     "sub4Admin",
		Key:       5,
		ParentKey: 2,
		Value:     "sub4Admin",
		Weight:    0,
	},
	{
		Title:     "sub5Admin",
		Key:       6,
		ParentKey: 5,
		Value:     "sub5Admin",
		Weight:    0,
	},
	{
		Title:     "sub6Admin",
		Key:       7,
		ParentKey: 5,
		Value:     "sub6Admin",
		Weight:    0,
	},
}

func TestTree(t *testing.T) {
	var (
		err error
		tr  *tree.Tree
	)
	if tr, err = tree.NewWithData(data); err != nil {
		t.Error(err.Error())
	}
	tjson, _ := json.MarshalIndent(tr.Tree(), "", "  ")
	t.Error(string(tjson))
}

func TestGetAllChildKey(t *testing.T) {
	var (
		err error
		tr  *tree.Tree
	)
	if tr, err = tree.NewWithData(data); err != nil {
		t.Error(err.Error())
	}
	node, err := tr.GetNode(5)
	if err != nil {
		t.Error(err.Error())
	}
	res := tr.GetAllChildKey(node)

	t.Error(res)
}
