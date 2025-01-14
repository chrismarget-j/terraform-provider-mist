---
page_title: "mist_org_wxtags Data Source - terraform-provider-mist"
subcategory: "WLAN"
description: |-
  This resource provides the list of Org Wlan tags (labels).The tags can be used   * within the WxRules to create filtering rules, or assign specific VLAN  * in the WLANs configuration to assign a WLAN to specific APs  * to identify unknown application used by Wi-Fi clients
---

# mist_org_wxtags (Data Source)

This resource provides the list of Org Wlan tags (labels).The tags can be used   * within the WxRules to create filtering rules, or assign specific VLAN  * in the WLANs configuration to assign a WLAN to specific APs  * to identify unknown application used by Wi-Fi clients


## Example Usage

```terraform
data "mist_device_switch_stats" "switch_stats" {
  org_id  = "15fca2ac-b1a6-47cc-9953-cc6906281550"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Read-Only

- `org_wxtags` (Attributes Set) (see [below for nested schema](#nestedatt--org_wxtags))

<a id="nestedatt--org_wxtags"></a>
### Nested Schema for `org_wxtags`

Read-Only:

- `created_time` (Number)
- `id` (String)
- `modified_time` (Number)
- `name` (String)
- `org_id` (String)