package lighthouse

import (
	"fmt"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

func cloudinit(key, caCrt, hostCrt, hostKey, version string) string {
	_caCrt := sdkv2betalib.WriteIndentedMultilineText(caCrt)
	_hostCrt := sdkv2betalib.WriteIndentedMultilineText(hostCrt)
	_hostKey := sdkv2betalib.WriteIndentedMultilineText(hostKey)

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
  - path: /etc/sysctl.d/99-disable-ipv6.conf
    content: |
      # Disable IPv6 on all interfaces
      net.ipv6.conf.all.disable_ipv6 = 1
      net.ipv6.conf.default.disable_ipv6 = 1
      net.ipv6.conf.lo.disable_ipv6 = 1
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
  - echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
  - echo "PasswordAuthentication no" >> /etc/ssh/sshd_config
  - echo "ChallengeResponseAuthentication no" >> /etc/ssh/sshd_config
  - systemctl restart sshd
  - sysctl --system
  - curl -L https://github.com/openecosystems/ecosystem/releases/download/apps-workloads-public-mesh-v2alpha-lighthouse/devel/apps-workloads-public-mesh-v2alpha-lighthouse_%s_Linux_x86_64.tar.gz | tar zx --strip-components=1 --directory /opt
  - chmod +x /opt/app
  - setcap cap_net_admin=+pe /opt/app
  - sudo systemctl enable lighthouse.service
  - sudo systemctl start lighthouse.service
  - ufw allow 4242/udp

`, key, _caCrt, _hostCrt, _hostKey, version)
}
