.DEFAULT_GOAL := help
CURDATE=$(shell date '+%Y%m%d')
#===================================== git ================================#
gitpush: # git 代码提交并推送
	@if [ ! $M ]; then \
		echo "add param: M=<your comment for this commit>"; \
		exit 1; \
	fi
	git commit -a -m "${M}"
	git push origin
tag: # 添加标签（当前日期）
	# 删除当前日期的 tag
	- git tag -d v${CURDATE}
	- git push origin :refs/tags/v${CURDATE}
	git tag -a v${CURDATE} -m "$(shell git log -1 --pretty=%B)" # 最后提交注释作为tag注释
	git push origin --tags
	@echo "发布 release 版本号必须为 y.mm.dd (大版本用1位数字)"

.PHONY: gitpush tag
.PHONY: help
help:
	@echo 'git提交并推送:      make gitpush M="提交说明"'
	@echo '创建tag(当前日期):  make tag'
	@echo 'make -n 检查语法'