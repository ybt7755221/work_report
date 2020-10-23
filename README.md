## &nbsp;&nbsp;&nbsp;&nbsp;work_report是一个基于gin的api微服务封装


#### 介绍:

##### 目录结构:

- controllers -- 处理入参，出参
- entities -- 数据库实例
- models -- 数据库操作
- service -- 处理业务逻辑
- router -- 路由
- middlewares -- 中间件
- libraries -- 工具库

##### libraries:

+ apolloCli - apollo连接库
+ database - 数据库连接库
+ efile - 文件操作库
+ elog - 错误日志库
+ redis - redis链接库
+ verify - 数据校验库
+ wmail - 邮件库
+ mongo - mongo连接库
---
#### 安装说明：

##### 安装
&nbsp;&nbsp;&nbsp;&nbsp;下载项目，在项目根目录运行：（本地测试推荐）
    
    go mod tidy
    go mod download
    go run main.go

##### 修改apollo配置

    //如果要使用apollo，修改配置
    const (
      	AppId = ""  //apollo app id
      	IpFAT = ""  //测试 apollo ip
      	IpPRO = ""  //正式 apollo ip
      	NameSpacename = "application"           //默认spacename
      	BackUpFile = "/etc/application.agollo"  //本地备份文件地址
    )

##### 自动生成entities
&nbsp;&nbsp;&nbsp;&nbsp;参见 [xorm工具](http://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-13/index.html)

&nbsp;&nbsp;&nbsp;&nbsp;通过工具生成后，将需要的表的实例放入项目下的entities,修改包名（自动生成的包名为model）

##### 自动生成项目文件
&nbsp;&nbsp;&nbsp;&nbsp; Linux/Mac:

    ./gtool -t GinContents -r contents -d Gin -c true -f ./

&nbsp;&nbsp;&nbsp;&nbsp; Windows:

    ./gtool.exe -t GinContents -r contents -d Gin -c true -f ./
    
Windows需要下载 [gtool](https://github.com/ybt7755221/gtool) 自己编译exe文件

---
#### 部分apollo配置说明：

&nbsp;&nbsp;&nbsp;&nbsp;环境变量
    
    ENVIRONMENT = fat   //判断apollo是测试还是正式
    
&nbsp;&nbsp;&nbsp;&nbsp;如需要判断项目优先使用：
    
    os.Getenv("ACTIVE") //判断docker环境是正式测试

&nbsp;&nbsp;&nbsp;&nbsp;是否开启prof（性能监控工具建议正式性能优化时开启，其他时间关闭）： 
    
    PPROF_STATUS = start
    
&nbsp;&nbsp;&nbsp;&nbsp; authentication中间件里的验签的secret
    
    SECRET = Dl*sCKW7C{SfYiPtYX*O5/71vG9&sm?2U
    
---
#### 参考文档地址

+ [xorm数据库操作文档](http://gobook.io/read/gitea.com/xorm/manual-zh-CN/#)
+ [xormplus](https://www.kancloud.cn/xormplus/xorm/167093)

