container := $(shell docker create checkin-client)

docker: curl.sh
	docker build -t checkin-client .

bundle: docker config.json
	rm -rf bundle
	mkdir -p bundle/rootfs
	docker export $(container) | tar -C bundle/rootfs -xvf -
	cp config.json bundle

run: bundle parallel-run.sh
	./parallel-run.sh 10

clean:
	rm -rf bundle bundle-* *~
