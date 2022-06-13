package tree

import (
	"encoding/json"
	"testing"

	"github.com/lgcgo/tree"
)

var data = []*tree.TreeData{
	{
		Title:     "superAdmin",
		Key:       "key-1",
		ParentKey: "",
		Value:     "superAdmin",
		Weight:    0,
	},
	{
		Title:     "sub1Admin",
		Key:       "key-2",
		ParentKey: "key-1",
		Value:     "sub1Admin",
		Weight:    0,
	},
	{
		Title:     "sub2Admin",
		Key:       "key-3",
		ParentKey: "key-1",
		Value:     "sub2Admin",
		Weight:    0,
	},
	{
		Title:     "sub3Admin",
		Key:       "key-4",
		ParentKey: "key-2",
		Value:     "sub3Admin",
		Weight:    0,
	},
	{
		Title:     "sub4Admin",
		Key:       "key-5",
		ParentKey: "key-2",
		Value:     "sub4Admin",
		Weight:    0,
	},
	{
		Title:     "sub5Admin",
		Key:       "key-6",
		ParentKey: "key-5",
		Value:     "sub5Admin",
		Weight:    0,
	},
	{
		Title:     "sub6Admin",
		Key:       "key-7",
		ParentKey: "key-5",
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
	tjson, _ := json.MarshalIndent(tr.GetTreeData(), "", "  ")
	t.Errorf("tree json: %s", string(tjson))
}

func TestGetAllChildKey(t *testing.T) {
	var (
		err error
		tr  *tree.Tree
		res []string
	)
	if tr, err = tree.NewWithData(data); err != nil {
		t.Error(err.Error())
	}
	res, err = tr.GetSpecChildKeys("key-5")
	if err != nil {
		t.Error(err.Error())
	}
	t.Errorf("child keys: %v", res)
}
