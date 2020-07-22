package template

var Conf =`[server]
ProjectName = {{.Name}}
#debug or prod
RunMode = prod
HttpPort = 8084
HttpHost = 127.0.0.1
ReadTimeout = 60
WriteTimeout = 60

[log]
ServiceName = {{.Name}}
Gin = ./logs/gin.log
App = ./logs/app.log

[database]
Type = mysql
Host = 127.0.0.1
User = developer
Password = 
Name = 
TablePrefix=

[redis]
Host = 127.0.0.1:6379
Password =
DB = 0
#redis集群
Cluster = false
Hosts =
`
