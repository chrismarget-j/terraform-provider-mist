// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_vpn

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgVpnResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
			"paths": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bfd_profile": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"",
									"broadband",
									"lte",
								),
							},
							Default: stringdefault.StaticString("broadband"),
						},
						"ip": schema.StringAttribute{
							Optional:            true,
							Description:         "if different from the wan port",
							MarkdownDescription: "if different from the wan port",
						},
						"pod": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{
								int64validator.Between(1, 128),
							},
							Default: int64default.StaticInt64(1),
						},
					},
					CustomType: PathsType{
						ObjectType: types.ObjectType{
							AttrTypes: PathsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Required: true,
			},
		},
	}
}

type OrgVpnModel struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	OrgId types.String `tfsdk:"org_id"`
	Paths types.Map    `tfsdk:"paths"`
}

var _ basetypes.ObjectTypable = PathsType{}

type PathsType struct {
	basetypes.ObjectType
}

func (t PathsType) Equal(o attr.Type) bool {
	other, ok := o.(PathsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t PathsType) String() string {
	return "PathsType"
}

func (t PathsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	bfdProfileAttribute, ok := attributes["bfd_profile"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`bfd_profile is missing from object`)

		return nil, diags
	}

	bfdProfileVal, ok := bfdProfileAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`bfd_profile expected to be basetypes.StringValue, was: %T`, bfdProfileAttribute))
	}

	ipAttribute, ok := attributes["ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip is missing from object`)

		return nil, diags
	}

	ipVal, ok := ipAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip expected to be basetypes.StringValue, was: %T`, ipAttribute))
	}

	podAttribute, ok := attributes["pod"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`pod is missing from object`)

		return nil, diags
	}

	podVal, ok := podAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`pod expected to be basetypes.Int64Value, was: %T`, podAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return PathsValue{
		BfdProfile: bfdProfileVal,
		Ip:         ipVal,
		Pod:        podVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewPathsValueNull() PathsValue {
	return PathsValue{
		state: attr.ValueStateNull,
	}
}

func NewPathsValueUnknown() PathsValue {
	return PathsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewPathsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (PathsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing PathsValue Attribute Value",
				"While creating a PathsValue value, a missing attribute value was detected. "+
					"A PathsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("PathsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid PathsValue Attribute Type",
				"While creating a PathsValue value, an invalid attribute value was detected. "+
					"A PathsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("PathsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("PathsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra PathsValue Attribute Value",
				"While creating a PathsValue value, an extra attribute value was detected. "+
					"A PathsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra PathsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewPathsValueUnknown(), diags
	}

	bfdProfileAttribute, ok := attributes["bfd_profile"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`bfd_profile is missing from object`)

		return NewPathsValueUnknown(), diags
	}

	bfdProfileVal, ok := bfdProfileAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`bfd_profile expected to be basetypes.StringValue, was: %T`, bfdProfileAttribute))
	}

	ipAttribute, ok := attributes["ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip is missing from object`)

		return NewPathsValueUnknown(), diags
	}

	ipVal, ok := ipAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip expected to be basetypes.StringValue, was: %T`, ipAttribute))
	}

	podAttribute, ok := attributes["pod"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`pod is missing from object`)

		return NewPathsValueUnknown(), diags
	}

	podVal, ok := podAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`pod expected to be basetypes.Int64Value, was: %T`, podAttribute))
	}

	if diags.HasError() {
		return NewPathsValueUnknown(), diags
	}

	return PathsValue{
		BfdProfile: bfdProfileVal,
		Ip:         ipVal,
		Pod:        podVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewPathsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) PathsValue {
	object, diags := NewPathsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewPathsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t PathsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewPathsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewPathsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewPathsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewPathsValueMust(PathsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t PathsType) ValueType(ctx context.Context) attr.Value {
	return PathsValue{}
}

var _ basetypes.ObjectValuable = PathsValue{}

type PathsValue struct {
	BfdProfile basetypes.StringValue `tfsdk:"bfd_profile"`
	Ip         basetypes.StringValue `tfsdk:"ip"`
	Pod        basetypes.Int64Value  `tfsdk:"pod"`
	state      attr.ValueState
}

func (v PathsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["bfd_profile"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["ip"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["pod"] = basetypes.Int64Type{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.BfdProfile.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["bfd_profile"] = val

		val, err = v.Ip.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ip"] = val

		val, err = v.Pod.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["pod"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v PathsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v PathsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v PathsValue) String() string {
	return "PathsValue"
}

func (v PathsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"bfd_profile": basetypes.StringType{},
		"ip":          basetypes.StringType{},
		"pod":         basetypes.Int64Type{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"bfd_profile": v.BfdProfile,
			"ip":          v.Ip,
			"pod":         v.Pod,
		})

	return objVal, diags
}

func (v PathsValue) Equal(o attr.Value) bool {
	other, ok := o.(PathsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.BfdProfile.Equal(other.BfdProfile) {
		return false
	}

	if !v.Ip.Equal(other.Ip) {
		return false
	}

	if !v.Pod.Equal(other.Pod) {
		return false
	}

	return true
}

func (v PathsValue) Type(ctx context.Context) attr.Type {
	return PathsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v PathsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bfd_profile": basetypes.StringType{},
		"ip":          basetypes.StringType{},
		"pod":         basetypes.Int64Type{},
	}
}
