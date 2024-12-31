package main

const CloudInitStartupScript = `
#cloud-config
write_files:
  - path: /etc/motd
    content: |
      Welcome to your the Event Multiplexer

# Install and start an Apache server
#packages:
#  - httpd
#runcmd:
#  - systemctl start httpd
`
