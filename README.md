<h1>A backend system for review service using Golang.<h1>
<h2>QuickView For Codes</h2>
- **review-service**:providing Remote Process Calls for users, stores and audits.<br>
- **service for users**: not inplemented serperately, http apis and grpc.<br>
- **service for shops**: review-b.<br>
- **service for audits**: review-o.<br>
<h2>Necessary Environments:</h2>
<h3>Kratos:</h3>
  ```
  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  ```<br>
<h3>MySQL(Local)</h3>:v8.1.0<br>
<h3>Redis(Local)</h3>:v.3.2.100<br>
<h3>Docker<h3>: find a suitable version in https://www.docker.com/, you need create a account first.<br>
<h3>Consul(In Docker)</h3>:
<h3>Canal(In Docker)</h3>:
```
  docker pull canal/canal-server:latest
  docker run -d --name canal-server -p 11111:11111 canal/canal-server
```<br>
<h3>Kafka(In docker)</h3>:
```
  go get github.com/segmentio/kafka-go
```<br>
<h3>ElasticSearch(In Docker)</h3>:
```
  go get github.com/elastic/go-elasticsearch/v8@latest
```<br>
