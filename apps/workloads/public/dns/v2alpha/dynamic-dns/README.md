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

# Manually Deploy Docker Image

```bash
cpln profile update jeannotcompany --login --org jeannotcompany
export CPLN_PROFILE="jeannotcompany"
cpln image docker-login
```

```bash
nx build workloads-private-mesh-v2alpha-cryptographic-mesh
nx container workloads-private-mesh-v2alpha-cryptographic-mesh
docker tag workloads/public/mesh/v2alpha/cryptographic-mesh:latest jeannotcompany.registry.cpln.io/mesh-v2alpha-cryptographic-mesh:JC-123
docker push jeannotcompany.registry.cpln.io/mesh-v2alpha-cryptographic-mesh:JC-123
cpln --org jeannotcompany --gvc local-gvc workload force-redeployment mesh-v2alpha-cryptographic-mesh
```

./nebula-cert sign -name "lighthouse1" -ip "192.168.100.1/24"
./nebula-cert sign -name "multiplexer1" -ip "192.168.100.5/24" -groups "multiplexers,ssh"
./nebula-cert sign -name "configuration-v2alpha-configuration" -ip "192.168.100.9/24" -groups "services"

### DEBUG ================

# On host

cd /root/.ssh
ssh-keygen -t ed25519 -f ssh_host_ed25519_key
chmod 400 ssh_host_ed25519_key

## Generate user key

ssh-keygen -t ed25519 -C "your_email@example.com"

eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

### DEBUG ================

# Enable CAP_NET for TUN device

setcap cap_net_admin=+pe /nebula
setcap cap_net_admin=+pe /lighthouse

## Disable IpV6

vim /etc/sysctl.d/70-disable-ipv6.conf

net.ipv6.conf.all.disable_ipv6 = 1

sysctl -p -f /etc/sysctl.d/70-disable-ipv6.conf

# Allow UDP on firewall

ufw allow 4242/udp

cd /
chmod +x nebula
./nebula -config /config.yml > /dev/null 2>&1 &
