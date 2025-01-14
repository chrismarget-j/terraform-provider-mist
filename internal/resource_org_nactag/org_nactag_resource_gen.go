// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_nactag

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgNactagResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allow_usermac_override": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "can be set to true to allow the override by usermac result",
				MarkdownDescription: "can be set to true to allow the override by usermac result",
				Default:             booldefault.StaticBool(false),
			},
			"egress_vlan_names": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `type`==`egress_vlan_names`, list of egress vlans to return",
				MarkdownDescription: "if `type`==`egress_vlan_names`, list of egress vlans to return",
			},
			"gbp_tag": schema.Int64Attribute{
				Optional:            true,
				Description:         "if `type`==`gbp_tag`",
				MarkdownDescription: "if `type`==`gbp_tag`",
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"match": schema.StringAttribute{
				Optional:            true,
				Description:         "if `type`==`match`",
				MarkdownDescription: "if `type`==`match`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"cert_cn",
						"cert_issuer",
						"cert_san",
						"cert_serial",
						"cert_sub",
						"client_mac",
						"idp_role",
						"mdm_status",
						"radius_group",
						"realm",
						"ssid",
						"user_name",
						"usermac_label",
					),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"match_all": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "This field is applicable only when `type`==`match`\n* `false`: means it is sufficient to match any of the values (i.e., match-any behavior)\n* `true`: means all values should be matched (i.e., match-all behavior)\n\n\nCurrently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`",
				MarkdownDescription: "This field is applicable only when `type`==`match`\n* `false`: means it is sufficient to match any of the values (i.e., match-any behavior)\n* `true`: means all values should be matched (i.e., match-all behavior)\n\n\nCurrently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`",
				Default:             booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"radius_attrs": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field \"radius_attrs\". \nIt is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.\nNote that it is allowed to have more than one radius_attrs in the result of a given rule.",
				MarkdownDescription: "if `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field \"radius_attrs\". \nIt is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.\nNote that it is allowed to have more than one radius_attrs in the result of a given rule.",
			},
			"radius_group": schema.StringAttribute{
				Optional:            true,
				Description:         "if `type`==`radius_group`",
				MarkdownDescription: "if `type`==`radius_group`",
			},
			"radius_vendor_attrs": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field \"radius_vendor_attrs\". \nIt is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.\nNote that it is allowed to have more than one radius_vendor_attrs in the result of a given rule.",
				MarkdownDescription: "if `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field \"radius_vendor_attrs\". \nIt is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.\nNote that it is allowed to have more than one radius_vendor_attrs in the result of a given rule.",
			},
			"session_timeout": schema.Int64Attribute{
				Optional:            true,
				Description:         "if `type`==`session_timeout, in seconds",
				MarkdownDescription: "if `type`==`session_timeout, in seconds",
			},
			"type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"egress_vlan_names",
						"match",
						"vlan",
						"gbp_tag",
						"radius_attrs",
						"radius_group",
						"radius_vendor_attrs",
						"session_timeout",
					),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"values": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `type`==`match`",
				MarkdownDescription: "if `type`==`match`",
			},
			"vlan": schema.StringAttribute{
				Optional:            true,
				Description:         "if `type`==`vlan`",
				MarkdownDescription: "if `type`==`vlan`",
			},
		},
	}
}

type OrgNactagModel struct {
	AllowUsermacOverride types.Bool   `tfsdk:"allow_usermac_override"`
	EgressVlanNames      types.List   `tfsdk:"egress_vlan_names"`
	GbpTag               types.Int64  `tfsdk:"gbp_tag"`
	Id                   types.String `tfsdk:"id"`
	Match                types.String `tfsdk:"match"`
	MatchAll             types.Bool   `tfsdk:"match_all"`
	Name                 types.String `tfsdk:"name"`
	OrgId                types.String `tfsdk:"org_id"`
	RadiusAttrs          types.List   `tfsdk:"radius_attrs"`
	RadiusGroup          types.String `tfsdk:"radius_group"`
	RadiusVendorAttrs    types.List   `tfsdk:"radius_vendor_attrs"`
	SessionTimeout       types.Int64  `tfsdk:"session_timeout"`
	Type                 types.String `tfsdk:"type"`
	Values               types.List   `tfsdk:"values"`
	Vlan                 types.String `tfsdk:"vlan"`
}
