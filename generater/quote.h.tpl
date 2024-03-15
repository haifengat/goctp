#include "../[[APIpath]]/ThostFtdcMdApi.h"

class Quote: CThostFtdcMdSpi{
public:
	[[ range .On ]]// [[ .Comment ]]    
    typedef void [[ .Name ]]Type(void*[[ range .Params ]], [[ .Type ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ end ]]);
    void *_[[ .Name ]];
    virtual void [[ .Name ]]([[ range $idx, $param := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Type ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ end ]]){
        if (_[[ .Name ]]) {
			(([[ .Name ]]Type*)_[[ .Name ]])(this[[ range .Params ]], [[ .Var ]][[ end ]]);
		}
    }
	[[ end ]]
};