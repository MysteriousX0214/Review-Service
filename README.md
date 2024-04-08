<h1>A backend system for review service using Golang.<h1>

<h2>QuickView For Codes</h2>
- **review-service**:providing Remote Process Calls for users, stores and audits.<br>
- **service for users**: not inplemented serperately, http apis and grpc.<br>
- **service for shops**: review-b.<br>
- **service for audits**: review-o.<br>
<h2>Necessary Environments:</h2>
<h6>Kratos:</h6><br>
  ```
  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  ```
<h6>MySQL(Local)</h6>:v8.1.0<br>
<h6>Redis(Local)</h6>:v.3.2.100<br>
<h6>Docker<h6>: find a suitable version in https://www.docker.com/, you need create a account first.<br>
<h6>Consul(In Docker)</h6>:<br>
<h6>Canal(In Docker)</h6>: <br>
```
  docker pull canal/canal-server:latest
  docker run -d --name canal-server -p 11111:11111 canal/canal-server
```
 
<h6>Kafka(In docker)</h6>:<br>
```
  go get github.com/segmentio/kafka-go
```
<h6>ElasticSearch(In Docker)</h6>:<br>
```
  go get github.com/elastic/go-elasticsearch/v8@latest
```
