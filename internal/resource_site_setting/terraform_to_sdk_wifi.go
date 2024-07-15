package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wifiTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WifiValue) *models.SiteWifi {
	tflog.Debug(ctx, "wifiTerraformToSdk")
	data := models.SiteWifi{}

	data.CiscoEnabled = d.CiscoEnabled.ValueBoolPointer()
	data.Disable11k = d.Disable11k.ValueBoolPointer()
	data.DisableRadiosWhenPowerConstrained = d.DisableRadiosWhenPowerConstrained.ValueBoolPointer()
	data.EnableArpSpoofCheck = d.EnableArpSpoofCheck.ValueBoolPointer()
	data.EnableSharedRadioScanning = d.EnableSharedRadioScanning.ValueBoolPointer()
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.LocateConnected = d.LocateConnected.ValueBoolPointer()
	data.LocateUnconnected = d.LocateUnconnected.ValueBoolPointer()
	data.MeshAllowDfs = d.MeshAllowDfs.ValueBoolPointer()
	data.MeshEnableCrm = d.MeshEnableCrm.ValueBoolPointer()
	data.MeshEnabled = d.MeshEnabled.ValueBoolPointer()
	data.MeshPsk = models.NewOptional(d.MeshPsk.ValueStringPointer())
	data.MeshSsid = models.NewOptional(d.MeshSsid.ValueStringPointer())
	data.ProxyArp = models.NewOptional(models.ToPointer(models.SiteWifiProxyArpEnum(d.ProxyArp.ValueString())))

	return &data
}