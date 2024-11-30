package main

const CloudInitStartupScript = `
#cloud-config
write_files:
  - path: /etc/motd
    content: |
      Welcome to your Pulumi-managed EC2 instance!

# Install and start an Apache server
#packages:
#  - httpd
#runcmd:
#  - systemctl start httpd

//setcap cap_net_admin=+pe /nebula
`
