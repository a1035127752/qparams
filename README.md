# qparams

#### 介绍
golang  url参数转struct



#### 安装教程



#### 测试


```

type Demo struct {
	Name     string   `json:"name"`
	Age      int 	  `json:"age"`
	Vip      bool	  `json:"vip"`
	Height   float32  `json:"height"`
}



func main(){
	u, err := url.Parse("https://example.com/users?name=Tom&age=23&vip=false&height=165.8")
	if err != nil {
		panic(err)
	}
	var demo Demo
	qparams.Unmarshal(u,&demo)
	fmt.Println(demo)
}
```


