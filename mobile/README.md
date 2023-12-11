<h1 align="center">osstp-go-hive</h1>
<div align="center">

Go + Flutter开发的管理系统脚手架, 分包合理, 精简易于扩展。

</div>

## 特性
### 服务端
- `Jwt` 使用JWT轻量级安全认证, 并提供活跃用户Token刷新功能
- `Casbin` 强大的、高效的开源访问控制框架，其权限管理机制支持多种访问控制模型
### 移动端
- `Flutter` Google移动端的跨平台开发框架，易于开发、维护，高效安全
- `Getx` 轻量且强大的状态管理、智能的依赖注入和便捷的路由管理解决方案

## 项目截图

![登录](https://github.com/gnimli/osstp-go-hive-ui/blob/main/src/assets/GithubImages/login.PNG)
![用户管理](https://github.com/gnimli/osstp-go-hive-ui/blob/main/src/assets/GithubImages/user.PNG)
![角色管理](https://github.com/gnimli/osstp-go-hive-ui/blob/main/src/assets/GithubImages/role.PNG)
![角色权限](https://github.com/gnimli/osstp-go-hive-ui/blob/main/src/assets/GithubImages/rolePermission.PNG)
![菜单管理](https://github.com/gnimli/osstp-go-hive-ui/blob/main/src/assets/GithubImages/menu.PNG)
![API管理](https://github.com/gnimli/osstp-go-hive-ui/blob/main/src/assets/GithubImages/api.PNG)

## 项目结构概览

```
├─initialize # casbin mysql zap validator 等公共资源
├─config # viper读取配置
├─controller # controller层，响应路由请求的方法
├─dto # 返回给前端的数据结构
├─middleware # 中间件
├─model # 结构体模型
├─dao # 数据库操作
├─response # 常用返回封装，如Success、Fail
├─routes # 所有路由
├─util # 工具方法
└─vo # 接收前端请求的数据结构

```

## 服务端Go项目

## 前端Vue项目
    osstp-go-hive-ui 
<https://github.com/gnimli/osstp-go-hive-ui.git>

## TODO

- 增加图片服务器
- 增加promtail-loki-grafana日志监控系统
- 增加swagger文档

## MIT License

    Copyright (c) 2021 gnimli

