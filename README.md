# mash-slack-mybot

## Getting Started

Edit credentials.sh from credentials.sh.sample
### Test locally

```
$ source credentials.sh
$ go run main.go
```
### Deploy to GKE

```
$ docker build -t mash-slack-mybot:latest .
$ docker tag mash-slack-mybot gcr.io/<your-pj-name>/mash-slack-mybot
$ gcloud docker -- push gcr.io/<your-pj-name>/mash-slack-mybot
```

Before creating this bot's pod, prepare redis.

See  https://github.com/kubernetes/kubernetes/tree/master/examples/storage/redis 

```
$ kubectl create -f k8s-deployment.yml
```



