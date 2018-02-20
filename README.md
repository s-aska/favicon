# Tomoshibi favicon generator

## Installing Cloud SDK

[Installing Cloud SDK](https://cloud.google.com/sdk/downloads#interactive)

```
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init

```

## Installing App Engine extension for Go

```
gcloud components install app-engine-go
```

## Installing direnv

https://github.com/direnv/direnv

## Installing dependency

```
cd appengine/gopath/src/lib
go get ./...
cd -
```

## Running the local development server

```
cd appengine/
dev_appserver.py app/app.yaml
```

## Deploying a Go App

```
cd appengine/
gcloud app deploy --project tmsb-favicon --version 1 app/app.yaml
```

## Local Unit Testing for Go

https://cloud.google.com/appengine/docs/standard/go/tools/localunittesting/
