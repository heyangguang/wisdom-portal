# wisdomPortal系统介绍

该系统由IBM团队开发与维护，使用Golang语言开发后端，Vue开发前端。
该系统致力于解决运维中的业务关键指标集合监控运维，业务中间表状态数据条目与自动发现问题，解决问题等。
使用Golang语言的Gin框架开发，项目刚刚起步，后面会逐步完善功能，最终达到我们想要的智慧系统。

# 系统权限控制
本项目采用casbin用于权限管控，PERM模型理论，但依旧使用的是RBAC的权限控制模型。有兴趣的小伙伴可以更改模型采用面向未来的ABAC权限控制模型。
系统权限使用权限模板抽象化了接口和请求动作，让用户在前端选择上有了很大的灵活性和更好的体验。
多个接口和请求动作可以合并成一个权限模板，一个用户或用户组可以对应多个权限模板。
用户权限的组成：用户权限 = 用户权限 + 用户组权限

# 多因子登录
由于系统内部维护大多数的核心业务集合与核心数据校验，所以在登录上加入了Google的双因子认证，使用TOTP的方式进行双因子认证，采用SHA1的加密算法，30S刷新一次。
想体验的小伙伴可以直接在AppStore里面下载谷歌的验证器或者RedHat的FreeOTP验证器。

# JWT登录
在我个人的想法里，JWT不需要存储在任何地方，因为它本身就是无状态的。如果非要存储，那为什么不直接使用Session + Cookie呢？
所以，我们的JWT没有刷新，默认12个小时就过期了。

项目功能：
- [ ] 用户功能
    - [X] 用户登录 
    - [X] 用户注册 
    - [ ] 用户注销 
    - [X] 添加用户
    - [ ] 删除用户
    - [ ] 修改用户
    - [ ] 用户列表
- [ ] 用户组功能
    - [X] 添加用户组
    - [ ] 删除用户组
    - [ ] 修改用户组
    - [ ] 用户组列表
- [ ] 权限功能
    - [X] 添加权限模板
    - [ ] 修改权限模板
    - [ ] 删除权限模板
    - [ ] 权限模板列表
    - [X] 添加用户权限模板
    - [ ] 删除用户权限模板
    - [ ] 修改用户权限模板
    - [ ] 添加用户组权限模板
    - [ ] 修改用户组权限模板
    - [ ] 删除用户组权限模板
- [ ] JWT登录功能
    - [X] JWT生成
    - [X] JWT验证
    - [ ] JWT删除
- [X] 双因子认证功能
    - [X] 创建双因子秘钥
    - [X] 校验双因子验证码
    - [X] 双因子认证二维码生成
    - [X] 双因子认证二维码查看
- [X] 接口文档
- [X] Gin日志模块

