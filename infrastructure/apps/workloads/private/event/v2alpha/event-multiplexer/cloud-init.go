package main

import (
	"fmt"

	infrastructurev2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
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
      Welcome to the Event Multiplexer!
  - path: /etc/sysctl.d/99-disable-ipv6.conf
    content: |
      # Disable IPv6 on all interfaces
      net.ipv6.conf.all.disable_ipv6 = 1
      net.ipv6.conf.default.disable_ipv6 = 1
      net.ipv6.conf.lo.disable_ipv6 = 1
  - content: |
      [Unit]
      Description=Event Multiplexer
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
    path: /lib/systemd/system/app.service
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
  - path: /opt/spec.yaml
    content: |
      app:
        name: event-v2alpha-event-multiplexer
        environmentName: 'dev-1'
        environmentType: 'development'
      http:
        port: '6477'
      zap:
        level: 'debug'
        development: true
        encoding: 'console'
      opentelemetry:
        traceProviderEnabled: true
      natsd:
        enabled: true
        options:
          serverName: "platform-leaf-node-local"
          host: "0.0.0.0"
          port: 4222
          debug: true
          leafNode:
            remotes:
              - urls:
                  scheme: "tls"
                  host:   "connect.ngs.global"
                credentials: "/etc/ngs/ngs.creds"
                tLSConfig:
                  insecureSkipVerify: true

      eventStreamRegistry:
        streams:
          - name: "audit"
            subjects:
              - "audit.>"
          - name: "configuration"
            subjects:
              - "configuration.>"
          - name: "certificateAuthority"
            subjects:
              - "certificateAuthority.>"
          - name: "certificate"
            subjects:
              - "certificate.>"
          - name: "decision"
            subjects:
              - "decision.>"
      nebula:
        tun:
          user: true
        punchy:
          punch: true
          respond: true
          delay: 1s
          respond_delay: 5s
        static_host_map:
          - '192.168.100.1'
        host:
          - '192.168.100.1': ['45.63.49.173:4242']
        lighthouse:
          am_lighthouse: false
          interval: 60
          hosts:
            - '192.168.100.1'
        firewall:
          outbound:
            - port: any
              proto: any
              host: any
          inbound:
            - port: any
              proto: icmp
              host: any
            - port: any
              proto: any
              host: any
        pki:
          ca: /etc/nebula/ca.crt
          cert: /etc/nebula/host.crt
          key: /etc/nebula/host.key

packages:
  - polkitd

runcmd:
  - sed -i '/PermitRootLogin/d' /etc/ssh/sshd_config
  - echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
  - echo "PasswordAuthentication no" >> /etc/ssh/sshd_config
  - echo "ChallengeResponseAuthentication no" >> /etc/ssh/sshd_config
  - systemctl restart sshd
  - sysctl --system
  - curl -L https://github.com/openecosystems/ecosystem/releases/download/apps-workloads-private-event-v2alpha-event-multiplexer/devel/apps-workloads-private-event-v2alpha-event-multiplexer_%s_Linux_x86_64.tar.gz | tar zx --strip-components=1 --directory /opt
  - chmod +x /opt/app
  - setcap cap_net_admin=+pe /opt/app
  - sudo systemctl enable app.service
  - sudo systemctl start app.service
  - sudo /opt/app &
  - sudo /opt/app &
  - ufw allow 6477/tcp
  - ufw allow 4222/tcp

`, key, _caCrt, _hostCrt, _hostKey, version)
}
