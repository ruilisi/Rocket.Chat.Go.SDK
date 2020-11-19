# Rocket.Chat Go SDK
[中文文档](https://github.com/ruilisi/Rocket.Chat.Go.SDK/blob/master/READMECN.md)
forked from [RocketChat/Rocket.Chat.Go.SDK](https://github.com/RocketChat/Rocket.Chat.Go.SDK)<br>
this version contains mostly apis that rocket-chat support. 

# struct
models: request struct <br>
conf: get config from application.yml<br>
rest: api function and response struct <br>
test: test 

# how to use 
## example
init client<br> 
host is your rocket_chat server address
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

This function is used to login rocket_chat and get userid and token.<br>
Token in rocket_chat valid in 90 days.You can store it in redis.When token used in other apis which needs token,you can take out token directly from redis,it will be quickly.<br>
Request apis from rocket_chat may timeout,use goroutine can set timeout return.Goroutine can also be used in your api server handlers when it need to use this project sdk. 
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

Unlike origin sdk,this project expose client's token and userid.You can set token and userid directly.
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

set params
```go
params := map[string][]string{
  "count": {"2"},
  "sort":  {"{\"value\":-1}"},
}
res, err := client.GroupList(params)
```

# License
Code and documentation copyright 2020 the [Rocket.Chat.Go.SDK Authors](https://github.com/ruilisi/Rocket.Chat.Go.SDK/graphs/contributors) and [ruilisi Network](https://ruilisi.co/) Code released under the [MIT License](https://github.com/ruilisi/Rocket.Chat.Go.SDK/blob/master/LICENSE).
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
