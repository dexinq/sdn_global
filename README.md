#### 结构
整个程序通过api向外提供服务，包括数据服务，控制服务等
api是整个程序的入口
api 提供对后台service的服务整合能力

程序整个通过consul作为registry


#### srv 文件夹中提供service服务

- appcron 中提供的服务是用于提供cron定时任务
- appmail 中的服务提供发送邮件的接口
- appmanage 用于操作app相关数据结构
- dataservice 数据服务用于提供数据相关的服务，查询修改等，是其他服务依赖的基础服务之一
- ovncontroller 用于提供操作ovn的接口，是其他服务依赖的基础服务之一
- ovscontroller 用于提供配置ovs的接口，是其他服务依赖的基础服务之一


#### api 文件夹中提供api服务

- app 主要提供将app相关service整合并向外提供服务的接口

#### util 中文件夹提供通用的各个srv通用的工具库

