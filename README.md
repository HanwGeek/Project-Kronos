# Project-Kronos

Project Kronos with a brilliant team.

## 参考资料

参考: <https://github.com/sshuair/awesome-gis>
WFS标准: <https://www.ogc.org/standards/wfs>

## Front End

* Vue + Vuetify + Openlayers/Mapbox (已遵照OGC标准)

## Back End

* Server: 
  * 处理来自Client的请求
  * 打开并解析 `.shp`文件或`PostGreSQL`数据库
  * 返回`JSON`
* Core:
  * 将数据解析为自主定义的数据结构
  * 可能的算法扩展

## Database

* Shapefile
* PostgreSQL

## How to commit

基本原则：

- 采用`git`作为版本控制系统，大家应该都已被添加为成员，所以不用再`fork`了。

- 充分利用`git`的分支管理策略，即`master`作为主分支，用于发布新版本，不在上面干活；自己干活时新建分支，再将修改提交到主分支上即可。

大致流程：(可参考[廖雪峰的官网网站](https://www.liaoxuefeng.com/wiki/896043488029600))

- `git clone https://github.com/HanwGeek/Project-Kronos`：克隆项目至本地文件夹
- `cd Project-Kronos`：进入本地项目文件夹中，可以发现当前分支是`master`
- `git checkout -b <branch_name>`：创建并切换至新分支(如：ygm_dev)，干活都在该分支上进行
- 干活ing...
- `git add <filename>`：添加修改的文件至缓存区
- `git commit -m <info>`：向本地库提交，并描述所做修改
- `git push -u origin <branch_name>`：向远程库(github)提交即可
- 坐等review...

## Contributors
@[HanwGeek](https://github.com/HanwGeek/) @[Nithouson](https://github.com/Nithouson) @[sunshineYin](https://github.com/sunshineYin) @[changxiaoyin](https://github.com/changxiaoyin/)

