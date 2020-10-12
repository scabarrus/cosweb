package cos

import(
  "context"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  //"net/http"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "log"
  //"fmt"
  //"encoding/json"
  //"time"
  "strconv"
)

func dataSourceCosInstance() *schema.Resource {
	log.Printf("[INFO] data read cos instances")
	return &schema.Resource{
		ReadContext: dataSourceCosInstanceRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			 },
			 "description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			 },
			 "resourcegroup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			 },
			 "resourceplanid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			 },
		},
	}
}


  func dataSourceCosInstanceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] dataSourceCosInstanceRead -start")
	var diags diag.Diagnostics
	client := m.(CosWeb)
	var list []CosInstance
	var err error
	name := d.Get("name").(string)
	
	log.Println("[DEBUG] name: ",name)
	list,err=client.GetAllCosInstances()
	if err != nil {
		return diag.FromErr(err)
	  }
	for k,v := range list{
		log.Println("[DEBUG] loop ",k,v)
		if name == v.Name {
			log.Printf("Name : %s Description : %s ResourceGroup : %s ResourcePlanId : %s ",v.Name,v.Description,v.ResourceGroup,v.ResourcePlanId)
			id:=strconv.Itoa(v.ID)
			d.SetId(id)
			d.Set("name",name)
			d.Set("description",v.Description)
			d.Set("resourcegroup",v.ResourceGroup)
			d.Set("resourceplanid",v.ResourcePlanId)
			
		}
	}
	return diags
}