package ctpdefine

type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2
)

[[ range .]]// [[ .Comment ]]
type [[ .FuncName ]] [[ .FuncTypeName|baseType ]]
[[ range .FuncFields ]]const [[ .FieldType ]] = [[ if eq (len .FieldName) 1 ]]'[[ .FieldName ]]'[[ else ]]"[[ .FieldName ]]"[[ end ]] // [[ .Comment ]]
[[ end ]]
[[ end ]]