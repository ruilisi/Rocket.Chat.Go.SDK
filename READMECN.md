# Rocket.Chat Go SDK
从 [RocketChat/Rocket.Chat.Go.SDK](https://github.com/RocketChat/Rocket.Chat.Go.SDK) fork来<br>
这个版本完成了几乎所有的rocket_chat的接口的sdk化工作

# 结构说明
models: 请求所需的结构  <br>
conf: 从application.yml获取配置的函数<br>
rest: sdk和返回结构 <br>
test: 测试

# 如何使用
## 样例
初始化client<br>
host 是rocket_chat服务地址
```go
var Client *rest.Client

func InitRcServer(b bool) {
	var _url url.URL
	var host string
	if b {
		host = conf.GetEnv("ROCKETCHAT_URL")
	} else {
		host = conf.GetEnv("ROCKETCHAT_LOCALHOST")
	}

	if host == "localhost:3000" {
		_url = url.URL{
			Host: host,
			//Scheme: "https",
		}
	} else {
		_url = url.URL{
			Host:   host,
			Scheme: "https",
		}
	}
	Client = rest.NewClient(&_url, false)
}
```

这个函数用来登录rocket_chat获取user_token和user_id<br>
token在rocket_chat里能有效90天，把token存储在redis里，当请求后续需要token的接口时，可以直接使用保存在redis里的token，这样请求效率比较高，避免重复登录。<br>
请求rocket_chat的api时，可能会遇见超时情况，使用并发可以设置超时时间，提升体验，如果你在使用此sdk时，同一handler下多次请求rocket_chat时，可以考虑使用并发，能够显著提升效率。
```go
import	"github.com/ruilisi/Rocket.Chat.Go.SDK/rest"

func LoginRc(user, password string) (userid, token string, err error) {
	redisToken := redis.HGetAll("rc:" + user)
	if len(redisToken) == 0 {
		finish := make(chan int)
		ch := make(chan string)
		res := make(chan string, 2)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			login := rest.AuthLoginRequest{
				User:     user,
				Password: password,
			}
			loginres, err := Client.AuthLogin(login)
			if err != nil {
				ch <- err.Error()
				return
			}
			if loginres.Status.Status != "success" {
				ch <- "login fail"
				return
			}
			redis.HSet("rc:"+user, loginres.Data.UserID, loginres.Data.AuthToken)
			redis.Expire("rc:"+user, time.Duration(2000)*time.Hour)
			res <- loginres.Data.UserID
			res <- loginres.Data.AuthToken
			defer wg.Done()
			return
		}()
		go func() {
			wg.Wait()
			finish <- 1
		}()

		select {
		case errmsg := <-ch:
			return "", "", fmt.Errorf("%s", errmsg)
		case <-finish:
			return <-res, <-res, nil
		case <-time.After(3 * time.Second):
			return "", "", fmt.Errorf("%s", "time out")
		}
	}

	for i, j := range redisToken {
		userid = i
		token = j
	}
	return userid, token, nil
}
```
不同于原项目，本项目将client的token和userid暴露了出来，这样就可以随时进行设置。
```go
import (
    immodels "github.com/ruilisi/Rocket.Chat.Go.SDK/models"
	"github.com/ruilisi/Rocket.Chat.Go.SDK/rest"
)

userid, token, err := LoginRc(user.Email, rocket_chat.ROCKETCHAT_PASS)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err})
	return
}
rocket_chat.Client.Auth = &rest.AuthInfo{
	Token: token,
	ID:    userid,
}

createGroup := immodels.CreateGroupRequest{
	Name:    fmt.Sprintf("%s_%s", department.Name, user.Username),
	Members: []string{user.Username, agentName},
}
createGroupRes, _ := Client.CreateGroup(&createGroup)
```

在某些方法里设置params参数
```go
params := map[string][]string{
  "count": {"2"},
  "sort":  {"{\"value\":-1}"},
}
res, err := client.GroupList(params)
```

# 执照
  代码和文档版权归2020年[MIT许可](https://github.com/ruilisi/Rocket.Chat.Go.SDK/blob/master/LICENSE)下发布的[Golang-pangu Authors](https://github.com/ruilisi/Rocket.Chat.Go.SDK/graphs/contributors) 和 [Ruilisi Network](https://ruilisi.co/)所拥有。
<table frame=void>
<tr>
<td >
<img src="logo.png" width="100px;" alt="hophacker"/>
</td>
</tr>
</table>

# Contributors
<table>
  <tr>
        <td align="center"><a href="https://github.com/ExcitingFrog"><img src="https://avatars2.githubusercontent.com/u/25655802?s=460&u=23017079e78e3c3bfa57a14bc369607b1b23c470&v=4" width="100px;" alt="ExcitingFrog"/><br /><sub><b>ExcitingFrog</b></sub></a><br /><a href="https://github.com/ruilisi/Rocket.Chat.Go.SDK/commits?author=ExcitingFrog" title="Code">💻</a> <a href="https://github.com/ruilisi/Rocket.Chat.Go.SDK/commits?author=ExcitingFrog" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Soryu23"><img src="https://avatars0.githubusercontent.com/u/67567977?s=460&u=fea632ad315bcdcfeff4de7ac5e2482b249929ac&v=4" width="100px;" alt="Soryu23"/><br /><sub><b>Soryu23</b></sub></a><br /><a href="https://github.com/ruilisi/Rocket.Chat.Go.SDK/commits?author=Soryu23" title="Code">💻</a> <a href="https://github.com/ruilisi/Rocket.Chat.Go.SDK/commits?author=Soryu23" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/chioazhao"><img src="https://avatars2.githubusercontent.com/u/59110803?s=460&u=1ac5a6b9811de1a89c38a6368c96ee3d552f62bf&v=4" width="100px;" alt="chioazhao"/><br /><sub><b>chioazhao</b></sub></a><br /><a href="https://github.com/ruilisi/Rocket.Chat.Go.SDK/commits?author=chioazhao" title="Code">💻</a> <a href="https://github.com/ruilisi/Rocket.Chat.Go.SDK/commits?author=chioazhao" title="Documentation">📖</a></td>
  </tr>
</table>

