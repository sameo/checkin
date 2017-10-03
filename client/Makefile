instances = ${INSTANCES}
runtime=${RUNTIME}

root = $(PWD)/checkin-clear-containers/

docker: curl.sh rm.sh
	docker build -t checkin-client .

bundle: docker config.json
	rm -rf bundle
	mkdir -p $(root)/bundle/rootfs
	docker export $(shell docker create checkin-client) | tar -C $(root)/bundle/rootfs -xvf -
	cp config.json $(root)/bundle/

prepare: bundle prepare-bundles.sh
	./prepare-bundles.sh $(instances) $(root)

run: parallel-run.sh
	./parallel-run.sh $(runtime) $(root)

clean:
	rm -rf bundle $(root) *~
