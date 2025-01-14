package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func stpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchStpConfig) StpConfigValue {

	var type_stp basetypes.StringValue

	if d.Type != nil {
		type_stp = types.StringValue(string(*d.Type))
	}

	data_map_attr_type := StpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"type": type_stp,
	}
	data, e := NewStpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
