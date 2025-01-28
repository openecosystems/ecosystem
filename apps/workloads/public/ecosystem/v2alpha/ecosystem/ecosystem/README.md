# ecosystem

==================

## Curl/HTTPie or simple network semantics
http localhost:6477/v2alpha/ecosystem \
Content-Type:application/json \
x-spec-workspace-slug:workspace123 \
x-spec-organization-slug:organization123 \
slug=123 \
name="Public Ecosystem" \
short_description="Public Ecosystem" \
description="Public Ecosystem" \
cidr="192.168.0.0/16"
