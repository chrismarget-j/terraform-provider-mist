package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func eslStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApStatsEslStat) basetypes.ObjectValue {
	var channel basetypes.Int64Value
	var connected basetypes.BoolValue
	var type_esl basetypes.StringValue
	var up basetypes.BoolValue

	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.Connected.Value() != nil {
		connected = types.BoolValue(*d.Connected.Value())
	}
	if d.Type.Value() != nil {
		type_esl = types.StringValue(*d.Type.Value())
	}
	if d.Up.Value() != nil {
		up = types.BoolValue(*d.Up.Value())
	}

	data_map_attr_type := EslStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"channel":   channel,
		"connected": connected,
		"type":      type_esl,
		"up":        up,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
