工程化实践               
# 工程项目结构                    
代码如何组织；目录怎么运营；代码怎么组织怎么分层；代码的初始化依赖注入怎么做。                         
依赖反转，依赖注入思想等。                    

[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

## Go目录
### /cmd
本项目的主干。                                 
每个应用程序的目录名应该与你想要的可执行文件的名称相匹配(例如，/cmd/myapp)。                                            
不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。                       
如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中。你会惊讶于别人会怎么做，所以要明确你的意图!                    
通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西。

**cmd应用目录负责程序的：启动、关闭、配置初始化等。**

### /internal
私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。请注意，这个布局模式是由 Go 编译器本身执行的。                                 
有关更多细节，请参阅Go 1.4 release notes 。注意，你并不局限于顶级 internal 目录。在项目树的任何级别上都可以有多个内部目录。                       
你可以选择向 internal 包中添加一些额外的结构，以分隔共享和非共享的内部代码。                             
这不是必需的(特别是对于较小的项目)，但是最好有有可视化的线索来显示预期的包的用途。                                  
你的实际应用程序代码可以放在 /internal/app 目录下(例如 /internal/app/myapp)，这些应用程序共享的代码可以放在 /internal/pkg 目录下(例如 /internal/pkg/myprivlib)。                                     

### /pkg
外部应用程序可以使用的库代码(例如 /pkg/mypubliclib)。其他项目会导入这些库，希望它们能正常工作，所以在这里放东西之前要三思:-)注意。                            
internal 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的。                                  
/pkg 目录仍然是一种很好的方式，可以显式地表示该目录中的代码对于其他人来说是安全使用的好方法。                                                     
由 Travis Jeffery 撰写的 I'll take pkg over internal 博客文章提供了 pkg 和 internal 目录的一个很好的概述，以及什么时候使用它们是有意义的。

当根目录包含大量非Go组件和目录时，这也是一种将Go代码分组到一个位置的方法。
这是一种常见的布局模式，但并不是所有人都接受它，一些 Go 社区的人也不推荐它。
如果你的应用程序项目真的很小，并且额外的嵌套并不能增加多少价值(除非你真的想要:-)，那就不要使用它。当它变得足够大时，你的根目录会变得非常繁琐时(尤其是当你有很多非 Go 应用组件时)，请考虑一下。

### /vendor
应用程序依赖项(手动管理或使用你喜欢的依赖项管理工具，如新的内置 Go Modules 功能)。go mod vendor 命令将为你创建 /vendor 目录。                       
请注意，如果未使用默认情况下处于启用状态的 Go 1.14，则可能需要在 go build 命令中添加 -mod=vendor 标志。                         
如果你正在构建一个库，那么不要提交你的应用程序依赖项。                         

注意，自从 1.13 以后，Go 还启用了模块代理功能(默认使用 https://proxy.golang.org 作为他们的模块代理服务器)。                                
在here 阅读更多关于它的信息，看看它是否符合你的所有需求和约束。如果需要，那么你根本不需要 vendor 目录。                          
国内模块代理功能默认是被墙的，七牛云有维护专门的的模块代理。                                               


## Kit Project Layout
基础kit库为独立项目，公司级的建议只有一个。                             
kit项目必须具备的特点：
- 统一
- 标准库布局
- 高度抽象
- 支持插件

[Develop Your Design Philosophy](https://www.ardanlabs.com/blog/2017/01/develop-your-design-philosophy.html)                                                                
[Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)

## 服务应用程序目录
### /api
OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件。
protobuf文件，以及生成的go文件。通常把API文档直接在proto文件中描述。                     

### /configs
配置文件模板或默认配置。
将你的 confd 或 consul-template 模板文件放在这里。

### /test
额外的外部测试应用程序和测试数据。你可以随时根据需求构造 /test 目录。对于较大的项目，有一个数据子目录是有意义的。
例如，你可以使用 /test/data 或 /test/testdata (如果你需要忽略目录中的内容)。
请注意，Go 还会忽略以“.”或“_”开头的目录或文件，因此在如何命名测试数据目录方面有更大的灵活性。



## 应用程序目录
放置多个微服务APP，也可以按照gitlab的group里简历多个project，每个project对应一个APP。
- 多APP方式
    APP目录内每个微服务按照自己的全局唯一名称来建立目录。如account/vip来建立目录，也可以是APPID。
- 和APP平级的目录pkg存放业务相关的公共库。如果不希望被导出这些目录，也可以放到myAPP/internal/pkg中。                             

微服务中的 app 服务类型分为4类：interface、service、job、admin。                                         
- interface: 对外的BFF服务，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口。                                 
- service: 对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务。                                 
- admin：区别于 service，更多是面向运营侧的服务，通常数据权限更高，隔离带来更好的代码级别安全。                              
- job: 流式任务处理的服务，上游一般依赖 message broker。                              
- task: 定时任务，类似 cronjob，部署到 task 托管平台中。                             

task更注重定时任务类的，job偏常驻执行任务。                                        



## B站实践
### V1
项目依赖路径：model->dao->service->api
model struct串联各个层，直到api需要做DTO对象转换。                              
DO（domain object)：领域对象，就是从现实世界中抽象出来的有形或无形的业务实体。缺乏 DTO -> DO 的对象转换。
    
### V2
app 目录下有 api、cmd、configs、internal 目录，目录里一般还会放置 README、CHANGELOG、OWNERS。
- internal: 是为了避免有同业务下有人跨目录引用了内部的 biz、data、service 等内部 struct。                        
- biz: 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，repo 接口在这里定义，使用依赖倒置的原则。                      
- data: 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起。
    data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。                        
- service: 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，
    同时协同各类 biz 交互，但是不应处理复杂逻辑。                                    

PO(Persistent Object): 持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，
那么数据表中的每个字段（或若干个）就对应 PO 的一个（或若干个）属性。                        
[ent](https://github.com/facebook/ent)


## lifecycle
Lifecycle 需要考虑服务应用的对象初始化以及生命周期的管理，所有 HTTP/gRPC 依赖的前置资源初始化，包括 data、biz、service，之后再启动监听服务。
我们使用 https://github.com/google/wire ，来管理所有资源的依赖注入。为何需要依赖注入？
[wire](https://github.com/google/wire)                          

## wire
[wire blog](https://blog.golang.org/wire)                   
手撸资源的初始化和关闭是非常繁琐，容易出错的。                                     
上面提到我们使用依赖注入的思路 DI，结合google wire，静态的go generate生成静态的代码，可以在很方便诊断和查看，不是在运行时利用reflection实现。


### /web
特定于 Web 应用程序的组件:静态 Web 资产、服务器端模板和 SPAs。  

 
## 通用应用目录
### /init
System init（systemd，upstart，sysv）和 process manager/supervisor（runit，supervisor）配置。

### /scripts
执行各种构建、安装、分析等操作的脚本。
这些脚本保持了根级别的 Makefile 变得小而简单(例如， https://github.com/hashicorp/terraform/blob/master/Makefile )。
有关示例，请参见/scripts 目录。

### /build
打包和持续集成。
将你的云( AMI )、容器( Docker )、操作系统( deb、rpm、pkg )包配置和脚本放在 /build/package 目录下。
将你的 CI (travis、circle、drone)配置和脚本放在 /build/ci 目录中。
请注意，有些 CI 工具(例如 Travis CI)对配置文件的位置非常挑剔。
尝试将配置文件放在 /build/ci 目录中，将它们链接到 CI 工具期望它们的位置(如果可能的话)。

### /deployments
IaaS、PaaS、系统和容器编配部署配置和模板(docker-compose、kubernetes/helm、mesos、terraform、bosh)。
注意，在一些存储库中(特别是使用 kubernetes 部署的应用程序)，这个目录被称为 /deploy。

## 其他目录
### /docs
设计和用户文档(除了 godoc 生成的文档之外)。

### /tools
这个项目的支持工具。注意，这些工具可以从 /pkg 和 /internal 目录导入代码。

### /examples
你的应用程序和/或公共库的示例。

### /third_party
外部辅助工具，分叉代码和其他第三方工具(例如 Swagger UI)。

### /githooks
Git hooks。

### /assets
与存储库一起使用的其他资产(图像、徽标等)。

### /website
如果你不使用 Github 页面，则在这里放置项目的网站数据。

## 你不应该拥有的目录
### /src
有些 Go 项目确实有一个 src 文件夹，但这通常发生在开发人员有 Java 背景，在那里它是一种常见的模式。如果可以的话，尽量不要采用这种 Java 模式。
你真的不希望你的 Go 代码或 Go 项目看起来像 Java:-)

不要将项目级别 src 目录与 Go 用于其工作空间的 src 目录(如 How to Write Go Code 中所述)混淆。
$GOPATH 环境变量指向你的(当前)工作空间(默认情况下，它指向非 windows 系统上的 $HOME/go)。这个工作空间包括顶层 /pkg, /bin 和 /src 目录。
你的实际项目最终是 /src 下的一个子目录，因此，如果你的项目中有 /src 目录，
那么项目路径将是这样的: /some/path/to/workspace/src/your_project/src/your_code.go。
注意，在 Go 1.11 中，可以将项目放在 GOPATH 之外，但这并不意味着使用这种布局模式是一个好主意。



















- API设计


- 配置管理



- 包管理


- 测试






