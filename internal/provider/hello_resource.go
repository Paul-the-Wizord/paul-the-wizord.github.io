package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NewHelloWorldResource returns a factory for hello_world resources.
func NewHelloWorldResource() resource.Resource {
	return &HelloWorldResource{}
}

// HelloWorldResource implements the hello_world resource.
type HelloWorldResource struct{}

// HelloWorldResourceModel is the state model for hello_world.
type HelloWorldResourceModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

// Metadata returns the resource type name.
func (r *HelloWorldResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "hello_world"
}

// Schema defines the hello_world resource schema.
func (r *HelloWorldResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Prints a greeting on create, update, and delete.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("world"),
				Description: "Name to greet. Defaults to \"world\".",
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Static identifier for this resource.",
			},
		},
	}
}

// Configure is a no-op.
func (r *HelloWorldResource) Configure(_ context.Context, _ resource.ConfigureRequest, _ *resource.ConfigureResponse) {
}

// Create prints the hello message and stores state.
func (r *HelloWorldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HelloWorldResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := plan.Name.ValueString()
	msg := fmt.Sprintf("hello, %s!", name)
	tflog.Info(ctx, msg)
	fmt.Println(msg)

	plan.Id = types.StringValue("hello-world")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read is a no-op for this stateless resource.
func (r *HelloWorldResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

// Update prints the changed message and updates state.
func (r *HelloWorldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan HelloWorldResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := plan.Name.ValueString()
	msg := fmt.Sprintf("%s changed!", name)
	tflog.Info(ctx, msg)
	fmt.Println(msg)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete prints the bye message and clears state.
func (r *HelloWorldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HelloWorldResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := state.Name.ValueString()
	msg := fmt.Sprintf("bye, %s!", name)
	tflog.Info(ctx, msg)
	fmt.Println(msg)
}
