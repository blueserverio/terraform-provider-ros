package ros

import (
	"context"
	"fmt"

	rosclient "github.com/blueserverio/ros/client"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIdentity() *schema.Resource {
	return &schema.Resource{
		Description: "Setting the System's Identity provides a unique identifying name for when the system identifies itself to other routers in the network and when accessing services such as DHCP, Neighbour Discovery, and default wireless SSID. The default system Identity is set to 'MikroTik'.",

		CreateContext: resourceSystemIdentityCreate,
		ReadContext:   resourceSystemIdentityRead,
		UpdateContext: resourceSystemIdentityUpdate,
		DeleteContext: resourceSystemIdentityDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The display name for the device.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func fillModel(d *schema.ResourceData) rosclient.SystemIdentity {

	// Create the model from the resource data.
	system_identity := new(rosclient.SystemIdentity)
	// system_identity.Id = d.Id()
	system_identity.Name = d.Get("name").(string)

	return *system_identity
}

func resourceSystemIdentityCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*rosclient.Client)
	var res interface{}
	err := client.Create("POST", "rest/system/identity/set", fillModel(d), &res)
	if err != nil {
		return diag.Errorf("Error creating system_identity", err)
	}
	return resourceSystemIdentityRead(ctx, d, meta)
}

func resourceSystemIdentityRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*rosclient.Client)
	res := rosclient.SystemIdentity{}
	err := client.Get("GET", "rest/system/identity", &res)

	if err != nil {
		return diag.Errorf("Error reading system_identity", err)
	}

	tflog.Info(ctx, fmt.Sprintf("Output of read (name): %s", res.Name))
	tflog.Info(ctx, fmt.Sprintf("Output of read (id): %s", res.Id))

	d.SetId(res.Name)
	d.Set("name", res.Name)
	return nil
}

func resourceSystemIdentityUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*rosclient.Client)
	var res interface{}
	err := client.Create("POST", "rest/system/identity/set", fillModel(d), &res)
	if err != nil {
		return diag.Errorf("Error updating system_identity", err)
	}
	return resourceSystemIdentityRead(ctx, d, meta)
}

func resourceSystemIdentityDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*rosclient.Client)
	input := fillModel(d)
	input.Name = "MikroTik"
	var res interface{}
	err := client.Create("POST", "rest/system/identity/set", input, &res)
	if err != nil {
		return diag.Errorf("Error updating system_identity", err)
	}
	return resourceSystemIdentityRead(ctx, d, meta)
}
