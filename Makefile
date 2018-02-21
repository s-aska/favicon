
atom:
	GOPATH=$(abspath appengine/gopath) atom .

install:
	./devtools/dep-ensure.sh

update:
	./devtools/dep-ensure.sh -update

serve:
	cd appengine; GOPATH=$(abspath appengine/gopath) dev_appserver.py --datastore_path=var/.datastore app/app.yaml

deploy:
	cd appengine; GOPATH=$(abspath appengine/gopath) gcloud app deploy --project tmsb-favicon --version 1 app/app.yaml
