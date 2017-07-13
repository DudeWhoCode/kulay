# Kulay
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=plastic)](https://raw.githubusercontent.com/DudeWhoCode/kulay/master/LICENSE) [![GitHub issues](https://img.shields.io/github/issues/DudeWhoCode/kulay.svg?style=plastic)](https://github.com/DudeWhoCode/kulay/issues) 

An high speed message passing system between various queues and services.

## Getting started
* Download the [latest binary](https://github.com/DudeWhoCode/kulay/releases/tag/v0.1.1) and add its location to your `PATH`
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

# Built with
* [Cobra](https://github.com/spf13/cobra) - command line framework
* [Viper](https://github.com/spf13/viper) - configuration handler
* [Logrus](https://github.com/sirupsen/logrus) - logging

# Versioning
This project uses [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/DudeWhoCode/kulay/tags)

