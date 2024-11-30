
# Manually Deploy Docker Image

```bash
cpln profile update jeannotcompany --login --org jeannotcompany
export CPLN_PROFILE="jeannotcompany"
cpln image docker-login
```

```bash
nx build workloads-public-mesh-v2alpha-lighthouse
nx container workloads-public-mesh-v2alpha-lighthouse
docker tag workloads/public/mesh/v2alpha/lighthouse:latest jeannotcompany.registry.cpln.io/mesh-v2alpha-lighthouse:JC-123
docker push jeannotcompany.registry.cpln.io/mesh-v2alpha-lighthouse:JC-123
cpln --org jeannotcompany --gvc local-gvc workload force-redeployment mesh-v2alpha-lighthouse
```



./nebula-cert sign -name "lighthouse1" -ip "192.168.100.1/24"
./nebula-cert sign -name "multiplexer1" -ip "192.168.100.5/24" -groups "multiplexers,ssh"
./nebula-cert sign -name "configuration-v2alpha-configuration" -ip "192.168.100.9/24" -groups "services"
