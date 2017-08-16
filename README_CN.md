
[English](https://github.com/go-xorm/cmd/blob/master/README.md)

# xorm 工具

xorm 是一组数据库操作命令行工具。 

## 源码安装

`go get github.com/go-xorm/cmd/xorm`

同时你需要安装如下依赖:

* github.com/go-xorm/xorm

* Mysql: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

* MyMysql: [github.com/ziutek/mymysql/godrv](https://github.com/ziutek/mymysql/godrv)

* Postgres: [github.com/lib/pq](https://github.com/lib/pq)

* SQLite: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) 

* MSSQL: [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)

** 对于sqlite3的支持，你需要自己进行编译 `go build -tags sqlite3` 因为sqlite3需要cgo的支持。

## 命令列表

有如下可用的命令：

* **reverse**     反转一个数据库结构，生成代码
* **shell**       通用的数据库操作客户端，可对数据库结构和数据操作
* **dump**        Dump数据库中所有结构和数据到标准输出
* **source**      从标注输入中执行SQL文件
* **driver**      列出所有支持的数据库驱动

## 反转

Reverse 命令让你根据数据库的表来生成结构体或者类代码文件。安装好工具之后，可以通过

`xorm help reverse`

获得帮助。

例子:

首先要进入到当前项目的目录下，主要是后面的命令最后一个参数中用到的模版存放在当前项目的目录下

`cd $GOPATH/github.com/go-xorm/cmd/xorm`

sqlite:
`xorm reverse sqite3 test.db templates/goxorm`

mysql:
`xorm reverse mysql root:@/xorm_test?charset=utf8 templates/goxorm`

mymysql:
`xorm reverse mymysql xorm_test2/root/ templates/goxorm`

postgres:
`xorm reverse postgres "dbname=xorm_test sslmode=disable" templates/goxorm`

之后将会生成代码 generated go files in `./model` directory

### 模版和配置

当前，默认支持Go，C++ 和 objc 代码的生成。具体可以查看源码下的 templates 目录。在每个模版目录中，需要放置一个配置文件来控制代码的生成。如下：

```
lang=go
genJson=1
```

`lang` 目前支持 go， c++ 和 objc。
`genJson` 可以为0或者1，如果是1则结构会包含json的tag，此项配置目前仅支持Go语言。

## Shell

Shell command provides a tool to operate database. For example, you can create table, alter table, insert data, delete data and etc.

`xorm shell sqlite3 test.db` will connect to the sqlite3 database and you can type `help` to list all the shell commands.

## Dump

Dump command provides a tool to dump all database structs and data as SQL to your standard output.

`xorm dump sqlite3 test.db` could dump sqlite3 database test.db to standard output. If you want to save to file, just
type `xorm dump sqlite3 test.db > test.sql`.

## Source

`xorm source sqlite3 test.db < test.sql` will execute sql file on the test.db.

## Driver

List all supported drivers since default build will not include sqlite3.

## LICENSE

 BSD License
 [http://creativecommons.org/licenses/BSD/](http://creativecommons.org/licenses/BSD/)
