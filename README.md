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

## Clone

```
mkdir $GOPATH/github.com/tmsbjp
cd $GOPATH/github.com/tmsbjp
git clone git@github.com:tmsbjp/favicon.git
```

## Installing dependency

```
make install
```

## Running the local development server

```
make serve
```

## Deploying a Go App

```
make deploy
```

## Local Unit Testing for Go

https://cloud.google.com/appengine/docs/standard/go/tools/localunittesting/
