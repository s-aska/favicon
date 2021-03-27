# favicon

## Setup

```sh
mkdir -p $GOPATH/src/github.com/s-aska
cd $GOPATH/src/github.com/s-aska
git clone git@github.com:s-aska/favicon.git
```

## Install

### Installing Go

```sh
brew install go
```

### Installing SDK and GAE/Go Components

[Installing Cloud SDK](https://cloud.google.com/sdk/downloads#interactive)

```sh
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init
gcloud components install app-engine-go
```

### Installing dependency

```sh
make installdeps
```

## Running the local development server

```sh
make dev-app
```

## Deployment command

```sh
make deploy-app
```

## Deployment App Engine

git push origin master -> Cloud Build -> https://YOUR_PROJECT_ID.appspot.com/

### How to Setup

See https://cloud.google.com/source-repositories/docs/quickstart-triggering-builds-with-source-repositories

1. `[PROJECT_NUMBER]@cloudbuild.gserviceaccount.com` Add `App Engine Admin` role [IAM](https://console.cloud.google.com/iam-admin/iam?project=YOUR_PROJECT_ID)
2. Enable the `Google App Engine Admin API` [App Engine Admin API](https://console.cloud.google.com/apis/api/appengine.googleapis.com/overview?project=YOUR_PROJECT_ID)
3. Add Trigger [Cloud Build](https://console.cloud.google.com/cloud-build/triggers?project=YOUR_PROJECT_ID)
   1. Select Source GitHub
   2. Choose Repository
   3. Click `Cloud Build 構成ファイル（yaml または json）`
   4. Settings Branches or Tags rule
   5. Create
