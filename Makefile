
install:
	cd appengine; dep ensure -update

serve:
	cd appengine; GOPATH=$(abspath appengine/gopath) dev_appserver.py --datastore_path=var/.datastore app/app.yaml

deploy:
	cd appengine; GOPATH=$(abspath appengine/gopath) gcloud app deploy --project tmsb-favicon --version 1 app/app.yaml
