.PHONY: help
help:
	@echo "Please use any of the following targets:"
	@echo "    build            Build the website, updating dist/apps/web/public/docs/v2alpha."
	@echo "    clean            Remove the contents from dist/apps/web/public/docs/v2alpha."
	@echo "    scan             Scan for broken links."
	@echo "    serve            Run a webserver in dist/apps/web/public/docs/v2alpha."

.PHONY: build
build: clean
	hugo

.PHONY: clean
clean:
	rm -rf ../../../../../dist/apps/web/public/docs/v2alpha/*

.PHONY: scan
scan: build
	linkinator http://localhost:1313 --recurse --format CSV --verbosity debug --skip bills.parliament.uk/bills/2806 --skip app.oeco.cloud 2>&1 | tee scan.csv

.PHONY: serve
serve:
	cd ../../../../../dist/apps/web/public/docs/v2alpha && python -m http.server 1313
