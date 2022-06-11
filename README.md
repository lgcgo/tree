# tree
A level-order tree traversal for golang.

# Features
- There is no recursive code to realize infinite parent-child data nesting.
- Support sorting by weight field.
- Virtual root node automatic.

# Installation
`go get -u github.com/lgcgo/tree`

# Example
````go
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

if tr, err := tree.NewWithData(data); err != nil {
 panic(err)
}
tjson, _ := json.MarshalIndent(tr.Tree(), "", "  ")
fmt.Print(string(tjson))

````

print
````json
{
  "title": "",
  "key": "root",
  "parent_key": "",
  "value": "",
  "weight": 0,
  "children": [
    {
      "title": "superAdmin",
      "key": "key-1",
      "parent_key": "root",
      "value": "superAdmin",
      "weight": 0,
      "children": [
        {
          "title": "sub1Admin",
          "key": "key-2",
          "parent_key": "key-1",
          "value": "sub1Admin",
          "weight": 0,
          "children": [
            {
              "title": "sub3Admin",
              "key": "key-4",
              "parent_key": "key-2",
              "value": "sub3Admin",
              "weight": 0,
              "children": []
            },
            {
              "title": "sub4Admin",
              "key": "key-5",
              "parent_key": "key-2",
              "value": "sub4Admin",
              "weight": 0,
              "children": [
                {
                  "title": "sub5Admin",
                  "key": "key-6",
                  "parent_key": "key-5",
                  "value": "sub5Admin",
                  "weight": 0,
                  "children": []
                },
                {
                  "title": "sub6Admin",
                  "key": "key-7",
                  "parent_key": "key-5",
                  "value": "sub6Admin",
                  "weight": 0,
                  "children": []
                }
              ]
            }
          ]
        },
        {
          "title": "sub2Admin",
          "key": "key-3",
          "parent_key": "key-1",
          "value": "sub2Admin",
          "weight": 0,
          "children": []
        }
      ]
    }
  ]
}
````


# License
Licensed under the [MIT license](LICENSE). Free & Forever.