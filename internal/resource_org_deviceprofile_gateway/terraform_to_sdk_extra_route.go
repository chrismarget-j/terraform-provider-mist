package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func extraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayExtraRoute {
	tflog.Debug(ctx, "extraRouteTerraformToSdk")
	data_map := make(map[string]models.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ExtraRoutesValue)

		data := models.GatewayExtraRoute{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		data_map[k] = data
	}
	return data_map
}
