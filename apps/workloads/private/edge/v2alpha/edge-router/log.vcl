set req.http.log-timing:log = time.elapsed.usec;

declare local var.origin_ttfb FLOAT;
declare local var.origin_ttlb FLOAT;

# Performance metrics for origin fetch if the request was a miss or pass
if (fastly_info.state ~ "^(MISS|PASS)") {
  # Time to first byte = fetch - misspass
  set var.origin_ttfb = std.atof(req.http.log-timing:fetch);
  set var.origin_ttfb -= std.atof(req.http.log-timing:misspass);

  # Time to last byte
  if (req.http.log-timing:do_stream == "1") {
    # log - misspass (and some clustering)
    set var.origin_ttlb = std.atof(req.http.log-timing:log);
    set var.origin_ttlb -= std.atof(req.http.log-timing:misspass);
  } else {
    # deliver - misspass (and some clustering)
    set var.origin_ttlb = std.atof(req.http.log-timing:deliver);
    set var.origin_ttlb -= std.atof(req.http.log-timing:misspass);
  }
  set var.origin_ttfb /= 1000;
  set var.origin_ttlb /= 1000;
  set req.http.log-origin:ttfb = var.origin_ttfb;
  set req.http.log-origin:ttlb = var.origin_ttlb;
}

# Client time to first byte (note this is just Fastly time, not network time to client)
declare local var.response_ttfb FLOAT;
set var.response_ttfb = time.to_first_byte;
set var.response_ttfb *= 1000;
set req.http.log-response:ttfb = var.response_ttfb;

# Client time to last byte
declare local var.response_ttlb FLOAT;
set var.response_ttlb = std.atof(req.http.log-timing:log);
set var.response_ttlb /= 1000;
set req.http.log-response:ttlb = var.response_ttlb;

# Estimated RTT
declare local var.client_tcpi_rtt INTEGER;
set var.client_tcpi_rtt = client.socket.tcpi_rtt;
set var.client_tcpi_rtt /= 1000;
set req.http.log-client:tcpi_rtt = var.client_tcpi_rtt;

# Only emit log origin/shield info if an origin/shield request was made
if (fastly_info.state !~ "^(MISS|PASS|ERROR)") {
  unset req.http.log-origin:host;
  unset req.http.log-origin:ip;
  unset req.http.log-origin:method;
  unset req.http.log-origin:name;
  unset req.http.log-origin:port;
  unset req.http.log-origin:reason;
  unset req.http.log-origin:shield;
  unset req.http.log-origin:status;
  unset req.http.log-origin:url;
}

log "syslog " req.service_id " frontsight :: {'client_ip':" client.ip;

log {"syslog "} req.service_id {" frontsight :: " {
  {""timestamp":""} strftime({"%Y-%m-%dT%H:%M:%S"}, time.start) "." time.start.usec_frac {"Z", "}
  # {""client_as_number":"} if(client.as.number!=0, client.as.number, "null") {", "}
  # {""client_city":"} if(client.geo.city, "%22"+json.escape(client.geo.city)+"%22", "null") {", "}
  # {""client_congestion_algorithm":""} json.escape(client.socket.congestion_algorithm) {"", "}
  # {""client_country_code":"} if(client.geo.country_code3, "%22"+json.escape(client.geo.country_code3)+"%22", "null") {", "}
  # {""client_cwnd":"} client.socket.cwnd {", "}
  # {""client_delivery_rate":"} client.socket.tcpi_delivery_rate {", "}
  # {""client_ip":""} json.escape(req.http.fastly-client-ip) {"", "}
  # {""client_latitude":"} if(client.geo.latitude!=999.9, client.geo.latitude, "null") {", "}
  # {""client_longitude":"} if(client.geo.longitude!=999.9, client.geo.longitude, "null") {", "}
  # {""client_ploss":"} client.socket.ploss {", "}
  # {""client_requests":"} client.requests {", "}
  # {""client_retrans":"} client.socket.tcpi_delta_retrans {", "}
  # {""client_rtt":"} req.http.log-client:tcpi_rtt {", "}
  # {""fastly_is_edge":"} if(fastly.ff.visits_this_service == 0, "true", "false") {", "}
  # {""fastly_is_shield":"} if(req.http.log-origin:shield == server.datacenter, "true", "false") {", "}
  # {""fastly_pop":""} json.escape(server.datacenter) {"", "}
  # {""fastly_server":""} json.escape(server.hostname) {"", "}
  # {""fastly_shield_used":"} if(req.http.log-origin:shield, "%22"+json.escape(req.http.log-origin:shield)+"%22", "null") {", "}
  # {""origin_host":"} if(req.http.log-origin:host, "%22"+json.escape(req.http.log-origin:host)+"%22", "null") {", "}
  # {""origin_ip":"} if(req.http.log-origin:ip, "%22"+json.escape(req.http.log-origin:ip)+"%22", "null") {", "}
  # {""origin_method":"} if(req.http.log-origin:method, "%22"+json.escape(req.http.log-origin:method)+"%22", "null") {", "}
  # {""origin_name":"} if(req.http.log-origin:name, "%22"+json.escape(req.http.log-origin:name)+"%22", "null") {", "}
  # {""origin_port":"} if(req.http.log-origin:port, req.http.log-origin:port, "null") {", "}
  # {""origin_reason":"} if(req.http.log-origin:reason, "%22"+json.escape(req.http.log-origin:reason)+"%22", "null") {", "}
  # {""origin_status":"} if(req.http.log-origin:status, req.http.log-origin:status, "null") {", "}
  # {""origin_ttfb":"} if(req.http.log-origin:ttfb ~ "^[0-9.]+$" , req.http.log-origin:ttfb, "null") {", "}
  # {""origin_ttlb":"} if(req.http.log-origin:ttlb ~ "^[0-9.]+$", req.http.log-origin:ttlb, "null") {", "}
  # {""origin_url":"} if(req.http.log-origin:url, "%22"+json.escape(req.http.log-origin:url)+"%22", "null") {", "}
  # {""request_host":"} if(req.http.log-request:host, "%22"+json.escape(req.http.log-request:host)+"%22", "null") {", "}
  # {""request_is_h2":"} if(fastly_info.is_h2, "true", "false") {", "}
  # {""request_is_ipv6":"} if(req.is_ipv6, "true", "false") {", "}
  # {""request_method":""} json.escape(req.http.log-request:method) {"", "}
  # {""request_referer":"} if(req.http.referer, "%22"+json.escape(req.http.referer)+"%22", "null") {", "}
  # {""request_tls_version":"} if(tls.client.protocol, "%22"+json.escape(tls.client.protocol)+"%22", "null") {", "}
  # {""request_url":""} json.escape(req.http.log-request:url) {"", "}
  # {""request_user_agent":"} if(req.http.user-agent, "%22"+json.escape(req.http.user-agent)+"%22", "null") {", "}
  # {""response_age":"} time.runits("s", obj.age) {", "}
  # {""response_bytes":"} resp.bytes_written {", "}
  # {""response_bytes_body":"} resp.body_bytes_written {", "}
  # {""response_bytes_header":"} resp.header_bytes_written {", "}
  # {""response_cache_control":"} if(resp.http.cache-control, "%22"+json.escape(resp.http.cache-control)+"%22", "null") {", "}
  # {""response_completed":"} if(resp.completed, "true", "false") {", "}
  # {""response_content_length":"} if(resp.http.content-length, resp.http.content-length, "null") {", "}
  # {""response_content_type":"} if(resp.http.content-type, "%22"+json.escape(resp.http.content-type)+"%22", "null") {", "}
  # {""response_reason":"} if(resp.response, "%22"+json.escape(resp.response)+"%22", "null") {", "}
  # {""response_state":""} json.escape(fastly_info.state) {"", "response_status":"} resp.status {", "}
  # {""response_ttfb":"} if(req.http.log-response:ttfb ~ "^[0-9.]+$", req.http.log-response:ttfb, "null") {", "}
  # {""response_ttl":"} obj.ttl {", "}
  # {""response_ttlb":"} if(req.http.log-response:ttlb ~ "^[0-9.]+$", req.http.log-response:ttlb, "null") {", "}
  # {""response_x_cache":"} if(resp.http.x-cache, "%22"+json.escape(resp.http.x-cache)+"%22", "null")
" }";