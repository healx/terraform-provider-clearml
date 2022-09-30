package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/healx/terraform-provider-clearml/internal/client"
)

func resourceQueue() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A queue in ClearML.",

		CreateContext: resourceQueueCreate,
		ReadContext:   resourceQueueRead,
		UpdateContext: resourceQueueUpdate,
		DeleteContext: resourceQueueDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the queue.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tags": {
				Description: "Tags to set on the queue.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceQueueCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*client.ClearMLClient)

	tags := []string{}

	if v, ok := d.GetOk("tags"); ok {
		for _, s := range v.([]interface{}) {
			tags = append(tags, s.(string))
		}
	}

	queueId, err := c.CreateQueue(
		ctx,
		d.Get("name").(string),
		tags,
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(queueId)

	return resourceQueueRead(ctx, d, meta)
}

func resourceQueueRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*client.ClearMLClient)

	queue, err := c.GetQueue(ctx, d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", queue["name"].(string))

	if v, ok := queue["tags"].([]interface{}); ok {
		d.Set("tags", v)
	}

	return nil
}

func resourceQueueUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*client.ClearMLClient)

	tags := []string{}

	if v, ok := d.GetOk("tags"); ok {
		for _, s := range v.([]interface{}) {
			tags = append(tags, s.(string))
		}
	}

	err := c.UpdateQueue(ctx,
		d.Id(),
		d.Get("name").(string),
		tags)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceQueueRead(ctx, d, meta)
}

func resourceQueueDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*client.ClearMLClient)

	err := c.DeleteQueue(ctx, d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
