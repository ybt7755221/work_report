## &nbsp;&nbsp;&nbsp;&nbsp;work_report是一个基于gin的日志汇总系统。源于本人不愿每日手动重复收集人员日报，并汇总周报而制作的自动生成系统


#### 介绍:
    此项目为日志系统后端代码，
    前端页面参见： [wp-front](https://github.com/ybt7755221/wp-front)

##### 目录结构:
参见 [gpi](https://github.com/ybt7755221/gpi)

##### 安装
&nbsp;&nbsp;&nbsp;&nbsp;下载项目，在项目根目录运行：（本地测试推荐）
    
    go mod tidy
    go mod download
    go run main.go

##### 修改apollo配置 
  
此处直接使用config文件，未使用apollo，如需使用apollo 参见 [gpi](https://github.com/ybt7755221/gpi)

##### 自动生成entities
&nbsp;&nbsp;&nbsp;&nbsp;参见 [xorm工具](http://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-13/index.html)

&nbsp;&nbsp;&nbsp;&nbsp;通过工具生成后，将需要的表的实例放入项目下的entities,修改包名（自动生成的包名为model）

##### 自动生成项目文件
&nbsp;&nbsp;&nbsp;&nbsp; Linux/Mac:

    ./vtool -t WrWorks -r work -d Gin -c true -f ./

&nbsp;&nbsp;&nbsp;&nbsp; Windows:

    ./gtool.exe -t WrWorks -r work -d Gin -c true -f ./
    
Windows暂不提供自动生成代码工具

---
#### 部分apollo配置说明：--暂未使用apollo

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

+ [gpi](https://github.com/ybt7755221/gpi)
+ [xorm数据库操作文档](http://gobook.io/read/gitea.com/xorm/manual-zh-CN/#)
+ [xormplus](https://www.kancloud.cn/xormplus/xorm/167093)

