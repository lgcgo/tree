# tree
A level-order tree traversal or golang.

# Features
- Support unlimited parent-child nesting of data
- Concise, elegant and efficient

# Installation
`go get -u github.com/lgcgo/tree`

# Example
In json format
````go
var data = []*tree.TreeData{
    {
        Title:     "User",
        Key:       1,
        ParentKey: 0,
        Value:     "/user",
        Weight:    50,
    },
    {
        Title:     "Add user",
        Key:       2,
        ParentKey: 1,
        Value:     "/user/add",
        Weight:    50,
    },
    {
        Title:     "Delete user",
        Key:       3,
        ParentKey: 1,
        Value:     "/user/delete",
        Weight:    49,
    },
}
t, err := tree.NewWithData(data)
tData := t.Tree()
tjson, _ := json.MarshalIndent(tData, "", "  ")
fmt.Println(string(tjson))
````

Console view
```json
{
  "title": "root",
  "key": 0,
  "parent_key": 0,
  "value": "",
  "weight": 0,
  "disabled": false,
  "children": [
    {
      "title": "User",
      "key": 1,
      "parent_key": 0,
      "value": "/user",
      "weight": 50,
      "disabled": false,
      "children": [
        {
          "title": "Add user",
          "key": 2,
          "parent_key": 1,
          "value": "/user/add",
          "weight": 50,
          "disabled": false,
          "children": []
        },
        {
          "title": "Delete user",
          "key": 3,
          "parent_key": 1,
          "value": "/user/delete",
          "weight": 49,
          "disabled": false,
          "children": []
        }
      ]
    }
  ]
}
```

# License
Licensed under the [MIT license](LICENSE). Free & Forever.