# A backend system for review service using Golang.
The project is the modified version of Qimi's online courses. If there is any inferingement, please contact me for deletion.
Course Address： [course address](https://study.163.com/course/courseMain.htm?courseId=1212937804)
## QuickView For Codes
**review-service**:providing Remote Process Calls for users, stores and audits.

supported methods:
- Detailed logics of all methods in **service for users/shops/audits**, http apis and remote process calls are provided
- select reviews from elasticsearch by storeID.
- select reviews from elasticsearch with not null comments.
 
**service for users**: not inplemented serperately, http apis and grpc methods are written in **review-service**.

supported methods:
- create review
- update review
- delete review
- see review details
- see all reviews created by someone

**service for shops**: review-b.

supported methods:(remote calls in **review-service**)
- create for reply
- update for reply
- appeal for reply

**service for audits**: review-o.
supported methods:(remote calls in **review-service**)
- audit for reviews
- audit for appeals

**read messages from kafka into elasticsearch**:review-job.

## Necessary Environments:
###### Kratos:
  ```
  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  ```
###### MySQL(Local):v8.1.0
Suggest to setup data tables in your MySQL first (review_info,review_reply_info and review_appeal_info), see [.sql file](https://github.com/MysteriousX0214/Review-Service/blob/master/review-service/review.sql) for details.  
###### Redis(Local):v.3.2.100
(**unimplemented**) Add cache to redis when querying for reviews.
###### Docker(Local): 
Find a suitable version in https://www.docker.com/, you need to create a account first.
###### Consul(In Docker):
```
git clone https://github.com/hashicorp/learn-consul-docker.git
cd datacenter-deploy-service-discovery
docker-compose up -d
```
Or view [link](https://www.liwenzhou.com/posts/Go/consul/) for details
###### Canal(In Docker): 
```
  docker pull canal/canal-server:latest
  docker run -d --name canal-server -p 11111:11111 canal/canal-server
```
###### Kafka(In docker):
```
  go get github.com/segmentio/kafka-go
```
See tutorial in [link](https://www.liwenzhou.com/posts/Go/kafka-go/) to setup kafka,zookeeper and kafka-ui in Docker
###### ElasticSearch(In Docker):
```
  go get github.com/elastic/go-elasticsearch/v8@latest
```
See tutorial in [link](https://www.liwenzhou.com/posts/Go/elasticsearch/) to setup elasticsearch and Kibana in Docker
###### Postman(Local,Optional):
To test if http apis and grpcs works well.
Find a suitable version in https://www.postman.com/, you may need to create an account for convinent use (like storing a certain http/grpc request route for multiplexing).
