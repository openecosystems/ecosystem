# dns

==================

# Manually Deploy Docker Image

```bash
cpln profile update jeannotcompany --login --org jeannotcompany
export CPLN_PROFILE="jeannotcompany"
cpln image docker-login
```

```bash
nx container workloads-public-dns-v2alpha-dynamic-dns
docker tag workloads/public/dns/v2alpha/dynamic-dns:latest jeannotcompany.registry.cpln.io/dns-v2alpha-dynamic-dns:JC-123
docker push jeannotcompany.registry.cpln.io/dns-v2alpha-dynamic-dns:JC-123
cpln --org jeannotcompany --gvc local-gvc workload force-redeployment dns-v2alpha-dynamic-dns
```
