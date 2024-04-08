<h1>A backend system for review service using Golang.<h1>
<h2>QuickView For Codes</h2>
- review-service:providing Remote Process Calls for users, stores and audits.
- service for users: not inplemented serperately, http apis and grpc
- service for shops: review-b
- service for audits: review-o
<h2>Necessary Environments:</h2>h1
1.Kratos:```go install github.com/go-kratos/kratos/cmd/kratos/v2@latest```<br>
2.MySQL(Local):v8.1.0<br>
3.Redis(Local):v.3.2.100<br>
4.Docker: find a suitable version in https://www.docker.com/, you need create a account first.<br>
5:Consul(In Docker):``` ```
6.Canal(In Docker):```
  docker pull canal/canal-server:latest
  docker run -d --name canal-server -p 11111:11111 canal/canal-server```<br>
7.Kafka(In docker):```
  go get github.com/segmentio/kafka-go
```<br>
8.ElasticSearch(In Docker):```
  go get github.com/elastic/go-elasticsearch/v8@latest
  ```<br>
