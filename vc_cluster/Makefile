NAMESPACE = vm-cluster


helm-repo-update:
	helm repo add vm https://victoriametrics.github.io/helm-charts/
	helm repo update

helm-chart-install:
	helm install victoria-metrics vm/victoria-metrics-cluster -f values.yaml -n $(NAMESPACE)

helm-chart-install-test: make-helm-chart make-namespace
	helm install victoria-metrics vm/victoria-metrics-cluster -f values.yaml -n $(NAMESPACE) --debug --dry-run

helm-chart-uninstall:
	helm uninstall victoria-metrics -n $(NAMESPACE)

make-namespace:
	kubectl create -f ./my-namespace.yaml

make-helm-chart:
	helm show values vm/victoria-metrics-cluster > values.yaml
