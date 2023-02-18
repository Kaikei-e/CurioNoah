# CurioNoah
CurioNoah is a set of applications that arouse curiosity and help you explore through intellectual pursuits.

## About

CurioNoah is composed of several applications that are based on RSS feeds. Some of the applications' roles have yet to be determined, and I am working on their implementation and ideas.

<br/>

- Curiosity: This is the core application of CurioNoah. It will help curate and expand knowledge in a way that stimulates curiosity and encourages exploration. It may use some kind of external API to aid in curation.

- InfoIsland: This is a web-based application that serves as a user interface and front end. It will aggregate and list all services. It is intended to be minimalistic with an integrated, unified user interface.
- InsightStream: This is the main backend and will be responsible for registering, reading, and updating feeds. The update function may be implemented in Rust for greater speed and efficiency, but currently this app is written in Go, my main development language.
- CoreManager: This will be the core service. It will be written in Rust as mentioned above (to study Rust) and will be the behind-the-scenes help, doing batch processing, etc.
- FeedFlare: This will be an RSS feed search aid, and will focus on discovering new RSS feeds.
- SlateFlex: This is a desire, but it is envisioned as a native application, and we would like to make it note-taking capable. I have absolutely no experience in native app development, so this is the lowest priority.


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

### Then, edit the hosts.

```
127.0.0.1 curionoah.com
```

<br/>

### Once these preparations are complete, run the following command at project root


```
docker compose up -d
```

<br/>

### Next, enter the InsightStream container and execute the following command

```
cd /usr/src/app
go generate ./ent
go run main.go
```


<br/>

## Why am I working on this project?

It is solely to improve my overall development capabilities.
This includes everything from bit infrastructures teritory such as Docker, DB, backend, and frontend development.
In addition, this is an ambitious project that aims to reach out to QA and performance tuning. (Yeah, I get it. I'm not a genius, so these things will be very challenging.)

<br>
