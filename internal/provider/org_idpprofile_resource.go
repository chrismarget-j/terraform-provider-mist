package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_idpprofile"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgOrgIdpprofileResource{}
	_ resource.ResourceWithConfigure = &orgOrgIdpprofileResource{}
)

func NewOrgIdpprofileResource() resource.Resource {
	return &orgOrgIdpprofileResource{}
}

type orgOrgIdpprofileResource struct {
	client mistapi.ClientInterface
}

func (r *orgOrgIdpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgIdpprofile client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *orgOrgIdpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_idpprofile"
}
func (r *orgOrgIdpprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages WAN Assurance Idp Profiles." +
			"An IDP Profile is a configuration setting that defines the behavior and actions of an intrusion detection and prevention (IDP) system." +
			"It specifies how the idp system should detect and respond to potential security threats or attacks on a network." +
			"The profile includes rules and policies that determine which types of traffic or attacks should be monitored," +
			"what actions should be taken when a threat is detected, and any exceptions or exclusions for specific destinations or attack types.",
		Attributes: resource_org_idpprofile.OrgIdpprofileResourceSchema(ctx).Attributes,
	}
}

func (r *orgOrgIdpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgIdpprofile Create")
	var plan, state resource_org_idpprofile.OrgIdpprofileModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idpprofile, diags := resource_org_idpprofile.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId := uuid.MustParse(plan.OrgId.ValueString())
	tflog.Info(ctx, "Starting OrgIdpprofile Create for Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsIDPProfiles().CreateOrgIdpProfile(ctx, orgId, &idpprofile)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating idpprofile",
			"Could not create idpprofile, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_idpprofile.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgOrgIdpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_idpprofile.OrgIdpprofileModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgIdpprofile Read: idpprofile_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	idpprofileId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsIDPProfiles().GetOrgIdpProfile(ctx, orgId, idpprofileId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting idpprofile",
			"Could not get idpprofile, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_idpprofile.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgOrgIdpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_idpprofile.OrgIdpprofileModel
	tflog.Info(ctx, "Starting OrgIdpprofile Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idpprofile, diags := resource_org_idpprofile.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgIdpprofile Update for OrgIdpprofile "+plan.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	idpprofileId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsIDPProfiles().UpdateOrgIdpProfile(ctx, orgId, idpprofileId, &idpprofile)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating idpprofile",
			"Could not update idpprofile, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_idpprofile.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgOrgIdpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_idpprofile.OrgIdpprofileModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgIdpprofile Delete: idpprofile_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	idpprofileId := uuid.MustParse(state.Id.ValueString())
	_, err := r.client.OrgsIDPProfiles().DeleteOrgIdpProfile(ctx, orgId, idpprofileId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting idpprofile",
			"Could not delete idpprofile, unexpected error: "+err.Error(),
		)
		return
	}
}
