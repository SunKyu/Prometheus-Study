CONFIG_PATH:= $(shell pwd)

run_vc:
	docker run -it --rm -p 8428:8428 -v $(CONFIG_PATH)/config.yaml:/etc/config.yaml \
		victoriametrics/victoria-metrics -promscrape.config=/etc/config.yaml

test:
	cat $(CONFIG_PATH)
