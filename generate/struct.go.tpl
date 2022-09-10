package ctpdefine

[[ range .]]
// 信息分发
type [[ .FuncTypeName ]] struct {
    [[ range .FuncFields ]]// [[ .Comment ]]
	[[ .FieldName ]] [[ .FieldType]]
    [[ end ]]
}
[[ end ]]