# 基于Gin框架的web项目

## 组织结构
|-- conf: 配置文件  
|-- controllers: 控制器层  
|-- libs: 项目依赖的第三方组件  
|-- models: 数据库模型层  
|-- utils: 各个项目能够通用的方法集合  
|-- main.go：项目入口文件

## 运行
默认配置见conf/dev.toml，项目启动命令`go run main.go`

测试分支配置见conf/test.toml，项目启动时指定环境变量`GOLANG_ENVIRONMENT=test go run main.go`