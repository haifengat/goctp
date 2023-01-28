package def

type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2
)

[[ range .]]// [[ .Comment ]]
type [[ .TypeName ]] [[ .TypeType ]]
[[ range .Consts ]]const [[ .Name ]]  = [[ if eq (len .Value) 1 ]]'[[ .Value ]]'[[ else ]]"[[ .Value ]]"[[ end ]] // [[ .Comment ]]
[[ end ]]
[[ end ]]