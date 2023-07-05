# ChatGPT-AccessToken

简介：主要用于批量动态生成 ChatGPT的Access Token，支持集群大批量生成过期时间校验。


### 项目产生的原因

由于潘多拉项目生产的Access Token 不稳定。会经常出现 `An error occurred: Error request login url.`。

### 功能
1. 指定邮箱只能在指定的代理IP申请。
2. 每个代理IP申请生产access_token后休息5-10秒
3. 加载代理池.

### 打包

```shell
make snapshots
```

### 镜像版本

```shell
docker pull askaigo/chatgpt-accesstoken:latest 
```