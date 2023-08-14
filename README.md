# lipgloss-text
Lipgloss helper libraly for making styled text with ease 
    
![Made with VHS](https://vhs.charm.sh/vhs-5AYoZ7xWGYNVishewRhD5q.gif)


# how to use 

## text/line
the two main types you should be using are `Line` and `Text`.
> there is also `Span`type but its not nececary most conviniet one to use  to to use 
Line from can take `string int float32/64 bool span ` and slices of ever of this types 
`Text` takes all types that `Line` takes and in additon it can take `Line []Line`
Non supported Types will return `type (Your Invalid Type) not suported!`


Line 
```go 
line := LineFrom("line from string", myStyle)
fmt.Println(line.Render())
```
Text 
```
text := TextFrom(line, myStyle)
fmt.Println(text.Render())
```
## options 
There are two options as for now (you can make an issue if you wan't me to add more) 
that can be used in `Render()` method. if non are passed defaults will be aplied.
If non values a
- `Styled`  bool value telling if style will be aplied  `def true`
- `Delimiter` delimiter bettwen strings for  `Line` `def " "` for `Text` `def "\n"`
