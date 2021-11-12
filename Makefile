.DEFAULT_GOAL := help
CURDATE=$(shell date '+%Y%m%d')
#===================================== git ================================#
gitee: # git 代码提交并推送
	@if [ ! $M ]; then \
		echo "add param: M=<your comment for this commit>"; \
		exit 1; \
	fi
	git commit -a -m "${M}"
	git push origin
tag: # 添加标签（当前日期）
	# 删除当前日期的 tag
	- git tag -d v0.6.5-${CURDATE}
	- git push origin :refs/tags/v0.6.5-${CURDATE}
	git tag -a v0.6.5-${CURDATE} -m "$(shell git log -1 --pretty=%B)" # 最后提交注释作为tag注释
	git push origin v0.6.5-${CURDATE}
	@echo "发布 release 版本号必须为 v0.6.5-yyyymmdd (6.5为CTP版本号)"
github: # 先执行 make gitee
	sed -i "s#gitee.com#github.com#g" go.mod
	sed -i "s#gitee.com#github.com#g" lnx/*.go
	sed -i "s#gitee.com#github.com#g" win/*.go
	sed -i "s#gitee.com#github.com#g" demo/*.go
	git commit -am "$(shell git log -1 --pretty=%B)"
	git push github
	# tag
	- git tag -d v0.6.5-${CURDATE}
	- git push github :refs/tags/v0.6.5-${CURDATE}
	git tag -a v0.6.5-${CURDATE} -m "$(shell git log -1 --pretty=%B)"
	git push github v0.6.5-${CURDATE}
	# 恢复
	git reset --hard HEAD~1

.PHONY: gitee github tag
.PHONY: help
help:
	@echo 'gitee提交并推送:      make gitee M="提交说明"'
	@echo 'github提交并推送:     make github'
	@echo '创建tag(当前日期):  make tag'
	@echo 'make -n 检查语法'
