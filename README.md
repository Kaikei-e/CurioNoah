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

<br/>

## Activation and termination. This is work around for development.

After you've completed the following ***Setting up*** steps, you can launch the program by running ***activation.sh***, and terminate the Rust program running on port 5100 by using ***terminate_all.sh***. Please note: Before proceeding, make sure to first execute the command 
```sh
docker compose up -d
```
and then you should add the permission 
to both shell files.

```bash
sudo chmod +x EACH_FILE
```
<br/>

Now, you can use it by accessing this

```url
http://curionoah.com:4173/
```

<br />

Basic auth is implementing right now.
You can bypass by using below account.

```
user > admin
password > admin
```


## Decision

I have made a certain decision. It is to base this suite of applications on the use of partial cloud services. There will be no changes to the basic structure or functionality. Operation is assumed with Docker Compose. However, I have decided to adopt a policy of allowing high flexibility in external use by constantly running this application on my home machine, authenticated via the internet using Cloudflare Tunnel. This will make the setup more complex, but the needs for use cases that require complete local operation will be met by a desktop application I am developing with Rust and Tauri, named CardinalAura.

Therefore, the following setup, which was never complete, will now become even less relevant. However, I plan to update the setup method when I feel inclined to do so in the future.

<br/>

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
POSTGRES_PASSWORD="testtesttestthisisatestmustmodify"
POSTGRES_USER="user1"
POSTGRES_DB="oath_keepers_db"
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
MYSQL_PASSWORD="testthisisatesttestestest"
MYSQL_DATABASE="curio_noah"
MYSQL_ADDR="192.168.100.10:3306"
NET_TYPE="tcp"
```

<br/>

/FeedHarmony/feed_harmony/.env

```
DATABASE_URL="mysql://user1:testthisisatesttestestest@localhost:3306/curio_noah"
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


## What is the motivation behind undertaking this project? 

The primary objective is to enhance my comprehensive development skills encompassing various areas such as Docker, database management, backend and frontend development. Furthermore, this ambitious project strives to encompass quality assurance and performance optimization, acknowledging that these aspects may pose significant challenges even for someone like me who isn't a prodigy.

<br>
