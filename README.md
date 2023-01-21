# CurioNoah
CurioNoah is a set of applications that arouse curiosity and help you explore through intellectual pursuits.

***Cloning and building the environment from scratch is too much work, so it will be improved.***

## Setting up

### First, place the three .env files.
The variables are example. So you can configure it.

/.env

```
MYSQL_ROOT_PASSWORD="testthisisatest"
MYSQL_DATABASE="curio_noah"
MYSQL_USER="userhogehoge"
MYSQL_PASSWORD="testthisisatesttestestest"
```


<br/>

/InfoIsland/InfoIsland/.env

```
VITE_INSIGHT_STREAM="http://curionoah.com:9000/api/v1"
# VITE_ORIGIN="http://curionoah.com:5173"
VITE_ORIGIN="http://curionoah.com:4173"
```

<br/>

/InsightStream/InsightStream/.env


```
MYSQL_USER="userhogehoge"
MYSQL_PASSWORD="testthisisatest"
MYSQL_DATABASE="curio_noah"
MYSQL_ADDR="192.168.100.10:3306"
NET_TYPE="tcp"
```


<br/>

### Next, enter the InsightStream container and execute the following command

```
cd /usr/src/app
go run main.go
```

## Why am I working on this project?

It is solely to improve my overall development capabilities.
This includes everything from bit infrastructures teritory such as Docker, DB, backend, and frontend development.
In addition, this is an ambitious project that aims to reach out to QA and performance tuning. (Yeah, I get it. I'm not a genius, so these things will be very challenging.)

<br>
