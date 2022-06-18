package tree_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lgcgo/tree"
)

var data = []*tree.TreeData{
	{
		Title:     "superAdmin",
		Key:       "key-1",
		ParentKey: "",
		Value:     "superAdmin",
		Weight:    50,
	},
	{
		Title:     "sub1Admin",
		Key:       "key-2",
		ParentKey: "key-1",
		Value:     "sub1Admin",
		Weight:    49,
	},
	{
		Title:     "sub2Admin",
		Key:       "key-3",
		ParentKey: "key-1",
		Value:     "sub2Admin",
		Weight:    48,
	},
	{
		Title:     "sub3Admin",
		Key:       "key-4",
		ParentKey: "key-2",
		Value:     "sub3Admin",
		Weight:    47,
	},
	{
		Title:     "sub4Admin",
		Key:       "key-5",
		ParentKey: "key-2",
		Value:     "sub4Admin",
		Weight:    46,
	},
	{
		Title:     "sub5Admin",
		Key:       "key-6",
		ParentKey: "key-5",
		Value:     "sub5Admin",
		Weight:    45,
	},
	{
		Title:     "sub6Admin",
		Key:       "key-7",
		ParentKey: "key-5",
		Value:     "sub6Admin",
		Weight:    44,
	},
}

func ExampleGetTreeData() {
	var (
		err error
		tr  *tree.Tree
	)

	if tr, err = tree.NewWithData(data); err != nil {
		fmt.Println(err.Error())
	}
	tjson, _ := json.MarshalIndent(tr.GetTreeData(), "", "  ")
	fmt.Println(string(tjson))

	// Output:
	// {
	//   "title": "Root",
	//   "key": "root",
	//   "parent_key": "",
	//   "value": "",
	//   "weight": 0,
	//   "children": [
	//     {
	//       "title": "superAdmin",
	//       "key": "key-1",
	//       "parent_key": "root",
	//       "value": "superAdmin",
	//       "weight": 50,
	//       "children": [
	//         {
	//           "title": "sub1Admin",
	//           "key": "key-2",
	//           "parent_key": "key-1",
	//           "value": "sub1Admin",
	//           "weight": 49,
	//           "children": [
	//             {
	//               "title": "sub3Admin",
	//               "key": "key-4",
	//               "parent_key": "key-2",
	//               "value": "sub3Admin",
	//               "weight": 47,
	//               "children": []
	//             },
	//             {
	//               "title": "sub4Admin",
	//               "key": "key-5",
	//               "parent_key": "key-2",
	//               "value": "sub4Admin",
	//               "weight": 46,
	//               "children": [
	//                 {
	//                   "title": "sub5Admin",
	//                   "key": "key-6",
	//                   "parent_key": "key-5",
	//                   "value": "sub5Admin",
	//                   "weight": 45,
	//                   "children": []
	//                 },
	//                 {
	//                   "title": "sub6Admin",
	//                   "key": "key-7",
	//                   "parent_key": "key-5",
	//                   "value": "sub6Admin",
	//                   "weight": 44,
	//                   "children": []
	//                 }
	//               ]
	//             }
	//           ]
	//         },
	//         {
	//           "title": "sub2Admin",
	//           "key": "key-3",
	//           "parent_key": "key-1",
	//           "value": "sub2Admin",
	//           "weight": 48,
	//           "children": []
	//         }
	//       ]
	//     }
	//   ]
	// }
}

func ExampleGetAllChildKey() {
	var (
		err error
		tr  *tree.Tree
		res []string
	)

	if tr, err = tree.NewWithData(data); err != nil {
		fmt.Println(err.Error())
	}
	res, err = tr.GetSpecChildKeys("key-5")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(strings.Join(res, ","))

	// Output:
	// key-5,key-6,key-7
}
