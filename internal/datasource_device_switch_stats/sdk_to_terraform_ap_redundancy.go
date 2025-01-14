package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func apRedundancyModuleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchStatsApRedundancyModule) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var num_aps basetypes.Int64Value
		var num_aps_with_switch_redundancy basetypes.Int64Value

		if d.NumAps != nil {
			num_aps = types.Int64Value(int64(*d.NumAps))
		}
		if d.NumApsWithSwitchRedundancy != nil {
			num_aps_with_switch_redundancy = types.Int64Value(int64(*d.NumApsWithSwitchRedundancy))
		}

		data_map_attr_type := ModulesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"num_aps":                        num_aps,
			"num_aps_with_switch_redundancy": num_aps_with_switch_redundancy,
		}
		data, e := NewModulesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, ModulesValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}

func apRedundancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchStatsApRedundancy) basetypes.ObjectValue {

	var modules basetypes.MapValue = types.MapNull(ModulesValue{}.Type(ctx))
	var num_aps basetypes.Int64Value
	var num_aps_with_switch_redundancy basetypes.Int64Value

	if d.Modules != nil {
		modules = apRedundancyModuleSdkToTerraform(ctx, diags, d.Modules)
	}
	if d.NumAps != nil {
		num_aps = types.Int64Value(int64(*d.NumAps))
	}
	if d.NumApsWithSwitchRedundancy != nil {
		num_aps_with_switch_redundancy = types.Int64Value(int64(*d.NumApsWithSwitchRedundancy))
	}

	data_map_attr_type := ApRedundancyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"modules":                        modules,
		"num_aps":                        num_aps,
		"num_aps_with_switch_redundancy": num_aps_with_switch_redundancy,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
