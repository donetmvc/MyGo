# eland-eoas-service
Run api

"# MyGo"

参考beego

获取beego安装包 $ go get -u github.com/astaxie/beego $ go get -u github.com/beego/bee

项目目录

workspace -bin -pkg -src

项目clone到src下面 执行 bee run

添加swagger

第一开启应用内文档开关，在配置文件中设置：EnableDocs = true, 这样你就已经内置了 docs 在你的 API 应用中，
然后你就使用 bee run -gendoc=true -downdoc=true,让我们的 API 应用跑起来，-gendoc=true 表示每次自动化的 build 文档，
-downdoc=true 就会自动的下载 swagger 文档查看器


mySql
go get -u github.com/go-sql-driver/mysql

xorm
go get github.com/go-xorm/xorm