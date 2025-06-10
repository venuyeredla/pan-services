#!/bin/bash
MODE=$1
USER="root"
PASS="ecompwd#24"

case $MODE in
    "server")
        echo "Bootstrapping server instance"
	docker run --name local-redis -p 6379:6379 -d redis:7.4.0-alpine
     ;;
    "start")
        ## starting existing conetainer
        docker start local-redis
      ;;

    "stop")
        ## stop existing conetainer
        docker stop local-redis
      ;;

    "obcli")
         echo "Post getting into container hit  command => "
         docker exec -it local-redis redis-cli
    ;;

    "excli")
         echo "Entering MySql instance"
    ;;

    *)
         echo "Invalid mode. It should be = server | start | stop | obcli | excli"
    ;;
esac

<<com
Multi line comment
docker run --name local-mysql -e MYSQL_ROOT_PASSWORD=ecompwd -d mysql:latest

#once inside the container 
mysql -uroot -p
#put/paste the password, and once inside MySQL CLI run
show databases;

$quit


<<com !"
