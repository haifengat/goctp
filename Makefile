.DEFAULT_GOAL := help
CURDATE=$(shell date '+%Y%m%d')
#===================================== git ================================#
gitpush: # git 代码提交并推送
	@if [ ! $M ]; then \
		echo "add param: M=<your comment for this commit>"; \
		exit 1; \
	fi
	git commit -a -m "${M}"
	git push origin v6.3.15
tag: # 回滚到 6.3.15
	# \cp v6.3.15_20190220/ctp*.go ctpdefine/
	# \cp v6.3.15_20190220/*.so lnx/
	# \cp v6.3.15_20190220/*.dll win/
	- git tag -d v0.6.3-${CURDATE}
	- git push origin v6.3.15 :refs/tags/v0.6.5-${CURDATE}
	git tag -a v0.6.3-${CURDATE} -m "$(shell git log -1 --pretty=%B)" # 最后提交注释作为tag注释
	git push origin v6.3.15 v0.6.3-${CURDATE} # 分支
	@echo "发布 release 版本号必须为 v0.6.3-yyyymmdd (6.3为CTP版本号)"
.PHONY: gitpush tag
.PHONY: help
help:
	@echo 'git提交并推送:      make gitpush M="提交说明"'
	@echo '创建tag(当前日期):  make tag'
	@echo 'make -n 检查语法'