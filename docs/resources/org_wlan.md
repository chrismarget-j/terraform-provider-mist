---
page_title: "mist_org_wlan Resource - terraform-provider-mist"
subcategory: "WLAN"
description: |-
  This resource manages the Org Wlans.The WLAN object contains all the required configuration to broadcast an SSID (Authentication, VLAN, ...)
---

# mist_org_wlan (Resource)

This resource manages the Org Wlans.The WLAN object contains all the required configuration to broadcast an SSID (Authentication, VLAN, ...)


## Example Usage

```terraform
resource "mist_org_wlan" "wlan_one" {
  ssid              = "wlan_one"
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
  bands             = ["5", "6"]
  vlan_id           = 143
  wlan_limit_up     = 10000
  wlan_limit_down   = 20000
  client_limit_up   = 512
  client_limit_down = 1000
  auth = {
    type = "psk"
    psk  = "secretpsk"
  }
  interface   = "all"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)
- `ssid` (String) the name of the SSID

### Optional

- `acct_immediate_update` (Boolean) enable coa-immediate-update and address-change-immediate-update on the access profile.
- `acct_interim_interval` (Number) how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled
- `acct_servers` (Attributes List) list of RADIUS accounting servers, optional, order matters where the first one is treated as primary (see [below for nested schema](#nestedatt--acct_servers))
- `airwatch` (Attributes) airwatch wlan settings (see [below for nested schema](#nestedatt--airwatch))
- `allow_ipv6_ndp` (Boolean) only applicable when limit_bcast==true, which allows or disallows ipv6 Neighbor Discovery packets to go through
- `allow_mdns` (Boolean) only applicable when limit_bcast==true, which allows mDNS / Bonjour packets to go through
- `allow_ssdp` (Boolean) only applicable when `limit_bcast`==`tru`e, which allows SSDP
- `ap_ids` (List of String) list of device ids
- `app_limit` (Attributes) bandwidth limiting for apps (applies to up/down) (see [below for nested schema](#nestedatt--app_limit))
- `app_qos` (Attributes) app qos wlan settings (see [below for nested schema](#nestedatt--app_qos))
- `apply_to` (String)
- `arp_filter` (Boolean) whether to enable smart arp filter
- `auth` (Attributes) authentication wlan settings (see [below for nested schema](#nestedatt--auth))
- `auth_server_selection` (String) When ordered, AP will prefer and go back to the first server if possible
- `auth_servers` (Attributes List) list of RADIUS authentication servers, at least one is needed if `auth type`==`eap`, order matters where the first one is treated as primary (see [below for nested schema](#nestedatt--auth_servers))
- `auth_servers_nas_id` (String) optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers
- `auth_servers_nas_ip` (String) optional, NAS-IP-ADDRESS to use
- `auth_servers_retries` (Number) radius auth session retries. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘retries’ are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3.
- `auth_servers_timeout` (Number) radius auth session timeout. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘quite-period’ and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10.
- `band_steer` (Boolean) whether to enable band_steering, this works only when band==both
- `band_steer_force_band5` (Boolean) force dual_band capable client to connect to 5G
- `bands` (List of String) list of radios that the wlan should apply to
- `block_blacklist_clients` (Boolean) whether to block the clients in the blacklist (up to first 256 macs)
- `bonjour` (Attributes) bonjour gateway wlan settings (see [below for nested schema](#nestedatt--bonjour))
- `cisco_cwa` (Attributes) Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html (see [below for nested schema](#nestedatt--cisco_cwa))
- `client_limit_down` (Number) kbps
- `client_limit_down_enabled` (Boolean) if downlink limiting per-client is enabled
- `client_limit_up` (Number) kbps
- `client_limit_up_enabled` (Boolean) if uplink limiting per-client is enabled
- `coa_servers` (Attributes List) list of COA (change of authorization) servers, optional (see [below for nested schema](#nestedatt--coa_servers))
- `disable_11ax` (Boolean) some old WLAN drivers may not be compatible
- `disable_ht_vht_rates` (Boolean) to disable ht or vht rates
- `disable_uapsd` (Boolean) whether to disable U-APSD
- `disable_v1_roam_notify` (Boolean) disable sending v2 roam notification messages
- `disable_v2_roam_notify` (Boolean) disable sending v2 roam notification messages
- `disable_wmm` (Boolean) whether to disable WMM
- `dns_server_rewrite` (Attributes) for radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns) (see [below for nested schema](#nestedatt--dns_server_rewrite))
- `dtim` (Number)
- `dynamic_psk` (Attributes) for dynamic PSK where we get per_user PSK from Radius
dynamic_psk allows PSK to be selected at runtime depending on context (wlan/site/user/...) thus following configurations are assumed (currently)
- PSK will come from RADIUS server
- AP sends client MAC as username ans password (i.e. `enable_mac_auth` is assumed)
- AP sends BSSID:SSID as Caller-Station-ID
- `auth_servers` is required
- PSK will come from cloud WLC if source is cloud_psks
- default_psk will be used if cloud WLC is not available
- `multi_psk_only` and `psk` is ignored
- `pairwise` can only be wpa2-ccmp (for now, wpa3 support on the roadmap) (see [below for nested schema](#nestedatt--dynamic_psk))
- `dynamic_vlan` (Attributes) for 802.1x (see [below for nested schema](#nestedatt--dynamic_vlan))
- `enable_local_keycaching` (Boolean) enable AP-AP keycaching via multicast
- `enable_wireless_bridging` (Boolean) by default, we'd inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be orwarded), wireless_bridging can be enabled
- `enable_wireless_bridging_dhcp_tracking` (Boolean) if the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless
- `enabled` (Boolean) if this wlan is enabled
- `fast_dot1x_timers` (Boolean) if set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’.
- `hide_ssid` (Boolean) whether to hide SSID in beacon
- `hostname_ie` (Boolean) include hostname inside IE in AP beacons / probe responses
- `hotspot20` (Attributes) hostspot 2.0 wlan settings (see [below for nested schema](#nestedatt--hotspot20))
- `inject_dhcp_option_82` (Attributes) (see [below for nested schema](#nestedatt--inject_dhcp_option_82))
- `interface` (String) where this WLAN will be connected to
- `isolation` (Boolean) whether to stop clients to talk to each other
- `l2_isolation` (Boolean) if isolation is enabled, whether to deny clients to talk to L2 on the LAN
- `legacy_overds` (Boolean) legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices.
- `limit_bcast` (Boolean) whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through)
- `limit_probe_response` (Boolean) limit probe response base on some heuristic rules
- `max_idletime` (Number) max idle time in seconds
- `mist_nac` (Attributes) (see [below for nested schema](#nestedatt--mist_nac))
- `mxtunnel_ids` (List of String) when `interface`=`mxtunnel`, id of the Mist Tunnel
- `mxtunnel_name` (List of String) when `interface`=`site_medge`, name of the mxtunnel that in mxtunnels under Site Setting
- `no_static_dns` (Boolean) whether to only allow client to use DNS that we’ve learned from DHCP response
- `no_static_ip` (Boolean) whether to only allow client that we’ve learned from DHCP exchange to talk
- `portal` (Attributes) portal wlan settings (see [below for nested schema](#nestedatt--portal))
- `portal_allowed_hostnames` (List of String) list of hostnames without http(s):// (matched by substring)
- `portal_allowed_subnets` (List of String) list of CIDRs
- `portal_denied_hostnames` (List of String) list of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames
- `portal_template_url` (String) N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template.
- `qos` (Attributes) (see [below for nested schema](#nestedatt--qos))
- `radsec` (Attributes) Radsec settings (see [below for nested schema](#nestedatt--radsec))
- `roam_mode` (String)
- `schedule` (Attributes) WLAN operating schedule, default is disabled (see [below for nested schema](#nestedatt--schedule))
- `sle_excluded` (Boolean) whether to exclude this WLAN from SLE metrics
- `template_id` (String)
- `thumbnail` (String) Url of portal background image thumbnail
- `use_eapol_v1` (Boolean) if `auth.type`==’eap’ or ‘psk’, should only be set for legacy client, such as pre-2004, 802.11b devices
- `vlan_enabled` (Boolean) if vlan tagging is enabled
- `vlan_id` (Number)
- `vlan_ids` (List of Number) vlan_ids to use when there’s no match from RA
- `vlan_pooling` (Boolean) vlan pooling allows AP to place client on different VLAN using a deterministic algorithm
- `wlan_limit_down` (Number) kbps
- `wlan_limit_down_enabled` (Boolean) if downlink limiting for whole wlan is enabled
- `wlan_limit_up` (Number) kbps
- `wlan_limit_up_enabled` (Boolean) if uplink limiting for whole wlan is enabled
- `wxtag_ids` (List of String) list of wxtag_ids
- `wxtunnel_id` (String) when `interface`=`wxtunnel`, id of the WXLAN Tunnel
- `wxtunnel_remote_id` (String) when `interface`=`wxtunnel`, remote tunnel identifier

### Read-Only

- `id` (String) The ID of this resource.
- `msp_id` (String)
- `portal_api_secret` (String) api secret (auto-generated) that can be used to sign guest authorization requests
- `portal_image` (String) Url of portal background image
- `portal_sso_url` (String)
- `site_id` (String)

<a id="nestedatt--acct_servers"></a>
### Nested Schema for `acct_servers`

Required:

- `host` (String) ip / hostname of RADIUS server
- `secret` (String, Sensitive) secret of RADIUS server

Optional:

- `keywrap_enabled` (Boolean)
- `keywrap_format` (String)
- `keywrap_kek` (String)
- `keywrap_mack` (String)
- `port` (Number) Acct port of RADIUS server


<a id="nestedatt--airwatch"></a>
### Nested Schema for `airwatch`

Optional:

- `api_key` (String) API Key
- `console_url` (String) console URL
- `enabled` (Boolean)
- `password` (String, Sensitive) password
- `username` (String) username


<a id="nestedatt--app_limit"></a>
### Nested Schema for `app_limit`

Optional:

- `apps` (Map of Number) Map from app key to bandwidth in kbps. 
Property key is the app key, defined in Get Application List
- `enabled` (Boolean)
- `wxtag_ids` (Map of Number) Map from wxtag_id of Hostname Wxlan Tags to bandwidth in kbps
Property key is the wxtag id


<a id="nestedatt--app_qos"></a>
### Nested Schema for `app_qos`

Optional:

- `apps` (Attributes Map) (see [below for nested schema](#nestedatt--app_qos--apps))
- `enabled` (Boolean)
- `others` (Attributes List) (see [below for nested schema](#nestedatt--app_qos--others))

<a id="nestedatt--app_qos--apps"></a>
### Nested Schema for `app_qos.apps`

Optional:

- `dscp` (Number)
- `dst_subnet` (String) subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
- `src_subnet` (String) subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)


<a id="nestedatt--app_qos--others"></a>
### Nested Schema for `app_qos.others`

Optional:

- `dscp` (Number)
- `dst_subnet` (String)
- `port_ranges` (String)
- `protocol` (String)
- `src_subnet` (String)



<a id="nestedatt--auth"></a>
### Nested Schema for `auth`

Optional:

- `anticlog_threshold` (Number) SAE anti-clogging token threshold
- `eap_reauth` (Boolean) whether to trigger EAP reauth when the session ends
- `enable_mac_auth` (Boolean) whether to enable MAC Auth, uses the same auth_servers
- `key_idx` (Number) when type=wep
- `keys` (List of String) when type=wep, four 10-character or 26-character hex string, null can be used. All keys, if provided, have to be in the same length
- `multi_psk_only` (Boolean) whether to only use multi_psk
- `owe` (String) `enabled` means transition mode
- `pairwise` (List of String) when type=psk / eap, one or more of wpa2-ccmp / wpa1-tkip / wpa1-ccmp / wpa2-tkip
- `private_wlan` (Boolean) whether private wlan is enabled. only applicable to multi_psk mode
- `psk` (String, Sensitive) when type=psk, 8-64 characters, or 64 hex characters
- `type` (String)
- `wep_as_secondary_auth` (Boolean) enable WEP as secondary auth


<a id="nestedatt--auth_servers"></a>
### Nested Schema for `auth_servers`

Required:

- `host` (String) ip / hostname of RADIUS server
- `secret` (String, Sensitive) secret of RADIUS server

Optional:

- `keywrap_enabled` (Boolean)
- `keywrap_format` (String)
- `keywrap_kek` (String)
- `keywrap_mack` (String)
- `port` (Number) Auth port of RADIUS server


<a id="nestedatt--bonjour"></a>
### Nested Schema for `bonjour`

Required:

- `additional_vlan_ids` (List of Number) additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
- `services` (Attributes Map) what services are allowed. 
Property key is the service name (see [below for nested schema](#nestedatt--bonjour--services))

Optional:

- `enabled` (Boolean) whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false

<a id="nestedatt--bonjour--services"></a>
### Nested Schema for `bonjour.services`

Optional:

- `disable_local` (Boolean) whether to prevent wireless clients to discover bonjour devices on the same WLAN
- `radius_groups` (List of String) optional, if the service is further restricted for certain RADIUS groups
- `scope` (String) how bonjour services should be discovered for the same WLAN



<a id="nestedatt--cisco_cwa"></a>
### Nested Schema for `cisco_cwa`

Optional:

- `allowed_hostnames` (List of String) list of hostnames without http(s):// (matched by substring)
- `allowed_subnets` (List of String) list of CIDRs
- `blocked_subnets` (List of String) list of blocked CIDRs
- `enabled` (Boolean)


<a id="nestedatt--coa_servers"></a>
### Nested Schema for `coa_servers`

Required:

- `ip` (String)
- `secret` (String, Sensitive)

Optional:

- `disable_event_timestamp_check` (Boolean) whether to disable Event-Timestamp Check
- `enabled` (Boolean)
- `port` (Number)


<a id="nestedatt--dns_server_rewrite"></a>
### Nested Schema for `dns_server_rewrite`

Optional:

- `enabled` (Boolean)
- `radius_groups` (Map of String) map between radius_group and the desired DNS server (IPv4 only)
Property key is the RADIUS group, property value is the desired DNS Server


<a id="nestedatt--dynamic_psk"></a>
### Nested Schema for `dynamic_psk`

Optional:

- `default_psk` (String, Sensitive) default PSK to use if cloud WLC is not available, 8-63 characters
- `default_vlan_id` (Number)
- `enabled` (Boolean)
- `force_lookup` (Boolean) when 11r is enabled, we'll try to use the cached PMK, this can be disabled
`false` means auto
- `source` (String)
- `vlan_ids` (List of Number)


<a id="nestedatt--dynamic_vlan"></a>
### Nested Schema for `dynamic_vlan`

Optional:

- `default_vlan_id` (Number) vlan_id to use when there’s no match from RADIUS
- `enabled` (Boolean) whether to enable dynamic vlan
- `local_vlan_ids` (List of Number) vlan_ids to be locally bridged
- `type` (String) standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco)
- `vlans` (Map of String) map between vlan_id (as string) to airespace interface names (comma-separated) or null for stndard mapping
  * if `dynamic_vlan.type`==`standard`, property key is the Vlan ID and property value is ""
  * if `dynamic_vlan.type`==`airespace-interface-name`, property key is the Vlan ID and property value is the Airespace Interface Name


<a id="nestedatt--hotspot20"></a>
### Nested Schema for `hotspot20`

Optional:

- `domain_name` (List of String)
- `enabled` (Boolean) whether to enable hotspot 2.0 config
- `nai_realms` (List of String)
- `operators` (List of String) list of operators to support
- `rcoi` (List of String)
- `venue_name` (String) venue name, default is site name


<a id="nestedatt--inject_dhcp_option_82"></a>
### Nested Schema for `inject_dhcp_option_82`

Optional:

- `circuit_id` (String)
- `enabled` (Boolean) whether to inject option 82 when forwarding DHCP packets


<a id="nestedatt--mist_nac"></a>
### Nested Schema for `mist_nac`

Optional:

- `enabled` (Boolean) when enabled:
* `auth_servers` is ignored
* `acct_servers` is ignored
* `auth_servers_*` are ignored
* `coa_servers` is ignored
* `radsec` is ignored
* `coa_enabled` is assumed


<a id="nestedatt--portal"></a>
### Nested Schema for `portal`

Optional:

- `amazon_client_id` (String) amazon OAuth2 client id. This is optional. If not provided, it will use a default one.
- `amazon_client_secret` (String) amazon OAuth2 client secret. If amazon_client_id was provided, provide a correspoinding value. Else leave blank.
- `amazon_email_domains` (List of String) Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
- `amazon_enabled` (Boolean) whether amazon is enabled as a login method
- `amazon_expire` (Number) interval for which guest remains authorized using amazon auth (in minutes), if not provided, uses expire`
- `auth` (String) authentication scheme
- `azure_client_id` (String) Required if `azure_enabled`==`true`.
Azure active directory app client id
- `azure_client_secret` (String) Required if `azure_enabled`==`true`.
Azure active directory app client secret
- `azure_enabled` (Boolean) whether Azure Active Directory is enabled as a login method
- `azure_expire` (Number) interval for which guest remains authorized using azure auth (in minutes), if not provided, uses expire`
- `azure_tenant_id` (String) Required if `azure_enabled`==`true`.
Azure active directory tenant id.
- `broadnet_password` (String, Sensitive) when `sms_provider`==`broadnet`
- `broadnet_sid` (String) when `sms_provider`==`broadnet`
- `broadnet_user_id` (String) when `sms_provider`==`broadnet`
- `bypass_when_cloud_down` (Boolean) whether to bypass the guest portal when cloud not reachable (and apply the default policies)
- `clickatell_api_key` (String) when `sms_provider`==`clickatell`
- `cross_site` (Boolean) whether to allow guest to roam between WLANs (with same `WLAN.ssid`, regardless of variables) of different sites of same org without reauthentication (disable random_mac for seamless roaming)
- `email_enabled` (Boolean) whether email (access code verification) is enabled as a login method
- `enabled` (Boolean) whether guest portal is enabled
- `expire` (Number) how long to remain authorized, in minutes
- `external_portal_url` (String) external portal URL (e.g. https://host/url) where we can append our query parameters to
- `facebook_client_id` (String) Required if `facebook_enabled`==`true`.
Facebook OAuth2 app id. This is optional. If not provided, it will use a default one.
- `facebook_client_secret` (String) Required if `facebook_enabled`==`true`.
Facebook OAuth2 app secret. If facebook_client_id was provided, provide a correspoinding value. Else leave blank.
- `facebook_email_domains` (List of String) Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
- `facebook_enabled` (Boolean) whether facebook is enabled as a login method
- `facebook_expire` (Number) interval for which guest remains authorized using facebook auth (in minutes), if not provided, uses expire`
- `forward` (Boolean) whether to forward the user to another URL after authorized
- `forward_url` (String) the URL to forward the user to
- `google_client_id` (String) Google OAuth2 app id. This is optional. If not provided, it will use a default one.
- `google_client_secret` (String) Google OAuth2 app secret. If google_client_id was provided, provide a correspoinding value. Else leave blank.
- `google_email_domains` (List of String) Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
- `google_enabled` (Boolean) whether google is enabled as login method
- `google_expire` (Number) interval for which guest remains authorized using google auth (in minutes), if not provided, uses expire`
- `gupshup_password` (String, Sensitive) when `sms_provider`==`gupshup`
- `gupshup_userid` (String) when `sms_provider`==`gupshup`
- `microsoft_client_id` (String) microsoft 365 OAuth2 client id. This is optional. If not provided, it will use a default one.
- `microsoft_client_secret` (String) microsoft 365 OAuth2 client secret. If microsoft_client_id was provided, provide a correspoinding value. Else leave blank.
- `microsoft_email_domains` (List of String) Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
- `microsoft_enabled` (Boolean) whether microsoft 365 is enabled as a login method
- `microsoft_expire` (Number) interval for which guest remains authorized using microsoft auth (in minutes), if not provided, uses expire`
- `passphrase_enabled` (Boolean) whether password is enabled
- `passphrase_expire` (Number) interval for which guest remains authorized using passphrase auth (in minutes), if not provided, uses `expire`
- `password` (String, Sensitive) passphrase
- `portal_api_secret` (String) api secret (auto-generated) that can be used to sign guest authorization requests
- `portal_image` (String) Url of portal background image
- `portal_sso_url` (String) for SAML, this is used as the ACS URL
- `predefined_sponsors_enabled` (Boolean) whether to show list of sponsor emails mentioned in `sponsors` object as a dropdown. If both `sponsor_notify_all` and `predefined_sponsors_enabled` are false, behaviour is acc to `sponsor_email_domains`
- `privacy` (Boolean)
- `puzzel_password` (String, Sensitive) when `sms_provider`==`puzzel`
- `puzzel_service_id` (String) when `sms_provider`==`puzzel`
- `puzzel_username` (String) when `sms_provider`==`puzzel`
- `sms_enabled` (Boolean) whether sms is enabled as a login method
- `sms_expire` (Number) interval for which guest remains authorized using sms auth (in minutes), if not provided, uses expire`
- `sms_message_format` (String)
- `sms_provider` (String)
- `sponsor_auto_approve` (Boolean) whether to automatically approve guest and allow sponsor to revoke guest access, needs predefined_sponsors_enabled enabled and sponsor_notify_all disabled
- `sponsor_email_domains` (List of String) list of domain allowed for sponsor email. Required if `sponsor_enabled` is `true` and `sponsors` is empty.
- `sponsor_enabled` (Boolean) whether sponsor is enabled
- `sponsor_expire` (Number) interval for which guest remains authorized using sponsor auth (in minutes), if not provided, uses expire`
- `sponsor_link_validity_duration` (Number) how long to remain valid sponsored guest request approve/deny link received in email, in minutes.
- `sponsor_notify_all` (Boolean) whether to notify all sponsors that are mentioned in `sponsors` object. Both `sponsor_notify_all` and `predefined_sponsors_enabled` should be true in order to notify sponsors. If true, email sent to 10 sponsors in no particular order.
- `sponsor_status_notify` (Boolean) if enabled, guest will get email about sponsor's action (approve/deny)
- `sponsors` (Map of String) object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty.
Property key is the sponsor email, Property value is the sponsor name
- `sso_default_role` (String) default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched
- `sso_forced_role` (String)
- `sso_idp_cert` (String) IDP Cert (used to verify the signed response)
- `sso_idp_sign_algo` (String) signing algorithm for SAML Assertion
- `sso_idp_sso_url` (String) IDP Single-Sign-On URL
- `sso_issuer` (String) IDP issuer URL
- `sso_nameid_format` (String)
- `telstra_client_id` (String) when `sms_provider`==`telstra`, Client ID provided by Telstra
- `telstra_client_secret` (String) when `sms_provider`==`telstra`, Client secret provided by Telstra
- `twilio_auth_token` (String) when `sms_provider`==`twilio`, Auth token account with twilio account
- `twilio_phone_number` (String) when `sms_provider`==`twilio`, Twilio phone number associated with the account. See example for accepted format.
- `twilio_sid` (String) when `sms_provider`==`twilio`, Account SID provided by Twilio


<a id="nestedatt--qos"></a>
### Nested Schema for `qos`

Optional:

- `class` (String)
- `overwrite` (Boolean) whether to overwrite QoS


<a id="nestedatt--radsec"></a>
### Nested Schema for `radsec`

Optional:

- `coa_enabled` (Boolean)
- `enabled` (Boolean)
- `idle_timeout` (Number)
- `mxcluster_ids` (List of String) To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids.
Org mxedge(s) identified by mxcluster_ids
- `proxy_hosts` (List of String) default is site.mxedge.radsec.proxy_hosts which must be a superset of all wlans[*].radsec.proxy_hosts
when radsec.proxy_hosts are not used, tunnel peers (org or site mxedges) are used irrespective of use_site_mxedge
- `server_name` (String) name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge.
- `servers` (Attributes List) List of Radsec Servers. Only if not Mist Edge. (see [below for nested schema](#nestedatt--radsec--servers))
- `use_mxedge` (Boolean) use mxedge(s) as radsecproxy
- `use_site_mxedge` (Boolean) To use Site mxedges when this WLAN does not use mxtunnel

<a id="nestedatt--radsec--servers"></a>
### Nested Schema for `radsec.servers`

Optional:

- `host` (String)
- `port` (Number)



<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Optional:

- `enabled` (Boolean)
- `hours` (Attributes) hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun). 

**Note**: If the dow is not defined then it’s treated as 00:00-23:59. (see [below for nested schema](#nestedatt--schedule--hours))

<a id="nestedatt--schedule--hours"></a>
### Nested Schema for `schedule.hours`

Optional:

- `fri` (String)
- `mon` (String)
- `sat` (String)
- `sun` (String)
- `thu` (String)
- `tue` (String)
- `wed` (String)


