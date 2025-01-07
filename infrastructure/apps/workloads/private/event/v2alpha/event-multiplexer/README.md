# Allow port 6477 and 4222 on this server as it actually listens on ports

- Install Nebula client using TUN device
- Route all traffic through TUN device except for what is on port 6477
- Bind port 4222 to TUN device
- Bind port 7999 to TUN device
- Bind port 6477 to eth0 device

run as nonroot

- Download binary
- Untar
- chmod +x
copy spec.yaml
- Mount block storagedrive and backup drive

# Allow Fastly to connect via mTLS

