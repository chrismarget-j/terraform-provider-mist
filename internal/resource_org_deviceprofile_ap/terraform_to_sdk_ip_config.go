package resource_org_deviceprofile_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IpConfigValue) *models.ApIpConfig {
	tflog.Debug(ctx, "ipConfigTerraformToSdk")
	data := models.ApIpConfig{}

	if !d.Dns.IsNull() && !d.Dns.IsUnknown() {
		data.Dns = mist_transform.ListOfStringTerraformToSdk(ctx, d.Dns)
	}
	if !d.DnsSuffix.IsNull() && !d.DnsSuffix.IsUnknown() {
		data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, d.DnsSuffix)
	}
	if d.Gateway.ValueStringPointer() != nil {
		data.Gateway = d.Gateway.ValueStringPointer()
	}
	if d.Gateway6.ValueStringPointer() != nil {
		data.Gateway6 = d.Gateway6.ValueStringPointer()
	}
	if d.Ip.ValueStringPointer() != nil {
		data.Ip = d.Ip.ValueStringPointer()
	}
	if d.Ip6.ValueStringPointer() != nil {
		data.Ip6 = d.Ip6.ValueStringPointer()
	}
	if d.Mtu.ValueInt64Pointer() != nil {
		data.Mtu = models.ToPointer(int(d.Mtu.ValueInt64()))
	}
	if d.Netmask.ValueStringPointer() != nil {
		data.Netmask = d.Netmask.ValueStringPointer()
	}
	if d.Netmask6.ValueStringPointer() != nil {
		data.Netmask6 = d.Netmask6.ValueStringPointer()
	}
	if d.IpConfigType.ValueStringPointer() != nil {
		data.Type = models.ToPointer(models.IpTypeEnum(d.IpConfigType.ValueString()))
	}
	if d.Type6.ValueStringPointer() != nil {
		data.Type6 = models.ToPointer(models.IpType6Enum(d.Type6.ValueString()))
	}
	if d.VlanId.ValueInt64Pointer() != nil {
		data.VlanId = models.ToPointer(int(d.VlanId.ValueInt64()))
	}

	return &data
}
