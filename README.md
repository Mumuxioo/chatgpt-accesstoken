# ChatGPT-AccessToken

简介：主要用于批量动态生成 ChatGPT的Access Token，支持集群大批量生成过期时间校验。


### 项目产生的原因

由于潘多拉项目生产的Access Token 不稳定。会经常出现 `An error occurred: Error request login url.`。

### 功能
1. 指定邮箱只能在指定的代理IP申请。[已完成]
2. 每个代理IP申请生产access_token后休息5-10秒 [未完成]
3. 加载代理池 [已完成]
4. 添加调度策略 [目前仅支持随机算法]
5. 每个IP申请完access_token后，10秒后才能申请。[未完成]
6. 目前仅实现本地版本。[缺失分布式版本]
7. IP代理可用统计 [未实现]
8. 对IP的增删改查 [实现]

### 如何使用

- 环境变量解释：
    - UseLocalDB: 是否使用本地加载代理的形式. 
    - PROXY_FILENAME：代理文件路径 [批量代理以文件读取的形式加载]
    - LogLevel: 日志级别
    - HttpBindAddress: 监听端口
    - REDIS_ADDRESS: redis地址配置
    - REDIS_PASSWORD: redis密码配置

- 示例演示：
   - /chatgpt-accesstoken/docker/local-docker-compose.yaml [本地演示]
   - /chatgpt-accesstoken/test/local-unuse-proxy.txt [代理示例文件]
   - /chatgpt-accesstoken/mux/api.http [接口测试示例]


### 打包

```shell
make snapshots
```

### 镜像版本

```shell
docker pull askaigo/chatgpt-accesstoken:latest 
```