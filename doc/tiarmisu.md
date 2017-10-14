
### 思路

#### Client
定时查询数据库，获取下载任务，然后通过 rpc 调用 aria 执行。


#### Server
提供以下 api。  
login  
add_task  
show_task  


#### 配置
1. 用户名以及密码。
2. 查询频率
3. mysql 配置

#### task 结构
```
url
type
status
```