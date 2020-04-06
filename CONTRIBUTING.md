# 贡献说明
本项目采用前后端分离的模式进行开发。前端采用的是react+ant design，后端采用golang+gin框架开发。

## 后端

后端启动方式为：go generate && go run main.go

后端采用的是 [gin-admin](https://github.com/LyricTian/gin-admin) 模板开发的，因此启动方式是一样的。
具体可参考gin-admin文档。如果只进行前端开发的话，也可以直接启动可执行文件。

## 前端

前端采用的 [gin-admin-web](https://github.com/LyricTian/gin-admin-react/tree/ts-master)模板开发的。具体启动方式可以参考gin-admin-web文档。


## 分支管理
- master为保护分支
- dev为开发分支
- 新特性的开发需要从dev拉取，命名规则为：feature_xxx
- 文档更新，命名规则为：patch_xxx
- 问题修复，命名规则为：fix_xxx

## PR说明

- 功能新增：add：具体的新增说明
- 功能修改：change：具体的修改说明
- 问题修复：fix：具体的问题说明
