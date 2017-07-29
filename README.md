# Kulay
[![Build Status](https://travis-ci.org/DudeWhoCode/kulay.svg?branch=master)](https://travis-ci.org/DudeWhoCode/kulay)
[![codecov](https://codecov.io/gh/DudeWhoCode/kulay/branch/master/graph/badge.svg)](https://codecov.io/gh/DudeWhoCode/kulay)
[![Go Report Card](https://goreportcard.com/badge/github.com/dudewhocode/kulay)](https://goreportcard.com/report/github.com/dudewhocode/kulay)
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/gokulay)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/DudeWhoCode/kulay/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/DudeWhoCode/kulay.svg)](https://github.com/DudeWhoCode/kulay/issues)

An high speed message passing system between various queues and services.

<a href="https://asciinema.org/a/IrbTcz6eO0IoBhZ196t4rdz6Y" target="_blank"><img src="https://asciinema.org/a/IrbTcz6eO0IoBhZ196t4rdz6Y.png" /></a>

## Getting started
* Download the [latest binary in binaries section](https://github.com/DudeWhoCode/kulay/releases) and add its location to your `PATH`
* Create a config file, kulay.toml and copy the following contents to it.
```
[sqs]
    [sqs.test_sg]
    queue_url = "https://sqs.ap-southeast-1.amazonaws.com/12345678/test_sg"
    region = "ap-southeast-1"
    delete_msg = false
    
    [sqs.test_us]
    queue_url = "https://sqs.us-east-1.amazonaws.com/12345678/test_us"
    region = "us-east-1"
    delete_msg = false
    
[jsonl]
    [jsonl.local_backup]
    path = "/tmp/backup.jsonl"
    rotate = true
    batch = 100
    
[redisq]
    [redisq.logbuffer]
    host = "localhost"
    port = "6379"
    password = ""
    database = 0
    queue = "test"
```
* Point the environment variable `KULAY_CONF` to the above file's location
```
$ export KULAY_CONF=/path/to/kulay.toml
```
* Change the sqs url and redis host to your endpoints. With reference to the config file, you can route messages between various services
```
$ kulay -f sqs.test_sg -t redisq.logbuffer

$ kulay -f sqs.test_us -t sqs.test_sg

$ kulay -f sqs.test_sg -t jsonl.local_backup

$ kulay -f jsonl.local_backup -t redisq.logbuffer
```

# Config file structure
Each section in the config file is a service which we can route messages to, It can be redis queue, redis pubsub, (AWS)SQS, rabbitmq, zeromq, csv or jsonl ([json delimited](http://jsonlines.org/)). Section names are already predefined for a service.
Each sub section can be any of your favourite name strictly preceeded by the section name. The sub section will contain various options with respect to the service you are using. An example redis queue section
```
[redisq]
    [redisq.logbuffer]
    host = "localhost"
    port = "6379"
    password = ""
    database = 0
    queue = "test"
```
Here the logbuffer is a custom name you give to your subsection, it can be production_buffer or prodqueue or which ever makes more sense to you. The `host`, `port`, `password`, `database` are the options we need to initilize the redis client and `queue` is the key name of the list in redis which will be used as queue.

Currently kulay supports passing messages between :
1. Redis queue
2. Redis pubsub
3. AWS SQS
4. Jsonl file read/write

### Config structures for supported services : 
#### Redis pubsub
```
[redispubsub]
       [redispubsub.yourCustomName]
       host = "localhost"
       port = "6379"
       password = "topsecret"
       database = 0
       channel = "mychannel"
```
host - Your redis hostname or IP address   
port - Port in which redis runs, default is 6379   
password - Password for your redis server, leave it as "" for no password     
channel - The pubsub channel which you need to send or receive messages 
 

### Redis queue
```
[redisq]
       [redisq.yourCustomName]
       host = "localhost"
       port = "6379"
       password = "topsecret"
       database = 0
       queue = "mychannel"
```
host - Your redis hostname or IP address   
port - Port in which redis runs, default is 6379   
password - Password for your redis server, leave it as "" for no password    
database - Default value is between 0-15, refer [redis documentation](https://redis.io/commands/SELECT)
queue - The queue to which you will send or receive messages   

### SQS
```
[sqs]
    [sqs.test_singapore]
    queue_url = "https://sqs.ap-southeast-1.amazonaws.com/12345678/test_queue"
    region = "ap-southeast-1"
    delete_msg = true
```
queue_url - URL of the queue found in AWS console    
region - The region where given queue was created   
database - Default value is between 0-15, refer [redis documentation](https://redis.io/commands/SELECT)
delete_msg - Delete flag, should be true if you want to delete the message from sqs after reading      

### Jsonl
```
[jsonl]
    [jsonl.local_backup]
    path = "/tmp/backup.jsonl"
    rotate = true
    batch = 1000
```
path   - Location where the files has to be created  
rotate - If rotate flag is enabled, kulay will create a new file everytime line count reaches number specified in `batch`    
batch  - The line count for a single file if rotate=true   
  
# Built with
* [Cobra](https://github.com/spf13/cobra) - command line framework
* [Viper](https://github.com/spf13/viper) - configuration handler
* [Logrus](https://github.com/sirupsen/logrus) - logging

# Versioning
This project uses [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/DudeWhoCode/kulay/tags)
