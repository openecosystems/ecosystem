package main

import (
	"fmt"
	infrastructurev2alphalib "libs/private/go/infrastructure/v2alpha"
)

func cloudinit(key, caCrt, hostCrt, hostKey, version string) string {
	_caCrt := infrastructurev2alphalib.WriteIndentedMultilineText(caCrt)
	_hostCrt := infrastructurev2alphalib.WriteIndentedMultilineText(hostCrt)
	_hostKey := infrastructurev2alphalib.WriteIndentedMultilineText(hostKey)

	return fmt.Sprintf(
		`#cloud-config
package_update: true
package_upgrade: true
users:
  - default
  - name: notroot
    groups: sudo
    sudo:
      - ALL=(ALL) NOPASSWD:ALL
    shell: /bin/bash
    lock_passwd: true
    ssh_authorized_keys:
      - %s
write_files:
  - path: /etc/motd
    content: |
      Welcome to your Lighthouse!
  - content: |
      [Unit]
      Description=Lighthouse Service
      ConditionPathExists=/opt/app
      After=network.target
       
      [Service]
      Type=simple
      User=notroot
      Group=notroot
      LimitNOFILE=1024
      
      Restart=on-failure
      RestartSec=10
      startLimitIntervalSec=60
      
      WorkingDirectory=/opt
      ExecStart=/opt/app
      
      # make sure log directory exists and owned by syslog
      #PermissionsStartOnly=true
      #ExecStartPre=/bin/mkdir -p /var/log/app
      #ExecStartPre=/bin/chown syslog:adm /var/log/app
      #ExecStartPre=/bin/chmod 755 /var/log/app
      #SyslogIdentifier=app
       
      [Install]
      WantedBy=multi-user.target
    path: /lib/systemd/system/lighthouse.service
    permissions: '0755'
    defer: true
  - path: /etc/nebula/ca.crt
    content: |
%s
  - path: /etc/nebula/host.crt
    content: |
%s
  - path: /etc/nebula/host.key
    content: |
%s

packages:
  - polkitd

runcmd:
  - sed -i '/PermitRootLogin/d' /etc/ssh/sshd_config
  - echo "PermitRootLogin no" >> /etc/ssh/sshd_config
  - echo "PasswordAuthentication no" >> /etc/ssh/sshd_config
  - echo "ChallengeResponseAuthentication no" >> /etc/ssh/sshd_config
  - systemctl restart sshd
  - curl -L https://github.com/openecosystems/ecosystem/releases/download/apps-workloads-public-mesh-v2alpha-lighthouse/devel/apps-workloads-public-mesh-v2alpha-lighthouse_%s_Linux_x86_64.tar.gz | tar zx --strip-components=1 --directory /opt
  - chmod +x /opt/app
  - sudo systemctl enable lighthouse.service
  - sudo systemctl start lighthouse.service

`, key, _caCrt, _hostCrt, _hostKey, version)
}