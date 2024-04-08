#A backend system for review service using Golang.

##QuickView For Codes

**review-service**:providing Remote Process Calls for users, stores and audits.
**service for users**: not inplemented serperately, http apis and grpc.
**service for shops**: review-b.
**service for audits**: review-o.
##Necessary Environments:
######Kratos:
  ```
  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  ```
######MySQL(Local):v8.1.0
######Redis(Local):v.3.2.100
######Docker: find a suitable version in https://www.docker.com/, you need create a account first.
######Consul(In Docker):
######Canal(In Docker): 
```
  docker pull canal/canal-server:latest
  docker run -d --name canal-server -p 11111:11111 canal/canal-server
```
######Kafka(In docker):
```
  go get github.com/segmentio/kafka-go
```
######ElasticSearch(In Docker):
```
  go get github.com/elastic/go-elasticsearch/v8@latest
```
