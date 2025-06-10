# Pan-services
Practical implementations of data storage systems like RDBMS,NOSQL,Information Retrieval as POC. In addition wiritng golan based micro services systems.

# Data managment
    Data integrity (CRC, checksums), Security(Encryption/decryption), encoding/decoding(XML,JSON,AVRO,Thrift,Protocol) and compresssion are imiportant aspects to be consider either storing to disk or transferring over network.
# Storage Systems for scalbilty
    Scalability, Fault tolerance, Latency.
# Scalability
    1. Partioning/Sharding 
    2. Replication
    3. Transactions
    4. Consistancy and consensus. 

# Data structures for creating indexes.

    Any storage system cosists of main data record and meta data called index for fast retrival.
    Storage system = Main data record + Index.

    1. Key value storage - In memory and on Disk storage. Map<K,Offset value main file> , SortedMap
    2. RDBMS - B+ Trees for indexing.
    3. Append only writes.
        a. SSTables - Sorted string tables.
        b. LSMTree - Log structured Merge Tree 

# RDBMS
    Using B+Tree based index.

# NoSQL 
  Append only for high write throghputs

## Developing Micro Services using. Golang .

# Web frameworks : 
    Gorilla =>
    Gin =>

# Authentication & authorization
(Oauth2 is used to access Microsoft Graph API)

JWT token implementation:
https://github.com/golang-jwt/jwt

 
# RDBMD Accessing databases 
Oracle : https://github.com/sijms/go-ora
Mysql : go get -u github.com/go-sql-driver/mysql

# HTTP client with certs to access HTTPS service.


Redis Integration:
https://github.com/redis/go-redis


Kafka :
https://github.com/confluentinc/confluent-kafka-go


Mail library:

Logging 
splunk.
Data dog.

# commands useful 
  to clean $> go mod tidy
  $ > go get github.com/gorilla/websocket


  ## Minikube :
  https://minikube.sigs.k8s.io/docs/drivers/docker/
  $> minikube start --driver=docker
  $> minikube config set driver docker
  $> minikube  service hello-minikube

## Kubernetes
    Services : kubectl get services | pods | deployments

    create $ kubectl create deployment hello-node --image=registry.k8s.io/e2e-test-images/agnhost:2.39 -- /agnhost netexec --http-port=8080

    expose $ kubectl expose deployment hello-minikube --type=LoadBalancer --port=8080


kubectl expose deployment/goapp-deployment --type=LoadBalancer --port=2024

minikube  service goapp-deployment


kubectl describe pod goapp-deployment-6995c466fd-5jkbk



# Docker compose commands 
  $> docker compose <command> <option>

   *Option*
        up -d   Create and start containers.
        down    Stop and remove containers. 
        stop    Stop services
        start   Start services 
 *docker compose start <service_name>

docker compose stop*