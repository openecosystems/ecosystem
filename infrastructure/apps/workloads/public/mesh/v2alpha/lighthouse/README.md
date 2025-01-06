# 

Copy over
- ca.crt
- server.crt
- server.key

create new nonroot user

run everything as nonroot

mv ca.crt /etc/nebula/ca.crt
mv server.crt /etc/nebula/host.crt
mv server.key /etc/nebula/host.key

download and untar lighthouse.tar.gz
chmod +x lighthouse
add script to run on startup


Open firewall 4242/udp

pulumi config set version
cat /path/to/id.pub | pulumi config set publicKey
cat /path/to/ca.crt | pulumi config set caCrt


./nebula-cert ca -name "Open Ecosystems, Inc"
./nebula-cert sign -name "local-1-mesh-v2alpha-lighthouse" -ip "192.168.100.1/24"
./nebula-cert sign -name "local-1-event-v2alpha-event-multiplexer" -ip "192.168.100.5/24" -groups "multiplexers,ssh"

cat ca.crt | pulumi config set caCrt
cat local-1-mesh-v2alpha-lighthouse.crt | pulumi config set hostCrt
cat local-1-mesh-v2alpha-lighthouse.key | pulumi config set hostKey

./nebula-cert sign -name "configuration-v2alpha-configuration" -ip "192.168.100.9/24" -groups "connectors"
