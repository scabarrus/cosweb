package cos

import (
    //"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "log"
    //"fmt"
)

func Provider() *schema.Provider {
    return &schema.Provider{
      Schema: map[string]*schema.Schema{
        "apiKey": &schema.Schema{
          Type:        schema.TypeString,
          //Optional:    true,
          Required: true,
          DefaultFunc: schema.EnvDefaultFunc("COS_APIKEY", nil),
        },
        "bearer": &schema.Schema{
          Type:        schema.TypeString,
          Required:    true,
          Sensitive:   true,
          DefaultFunc: schema.EnvDefaultFunc("COS_BEARER", nil),
        },
        "endpoint": &schema.Schema{
            Type:     schema.TypeString,
            Required: true,
        },
        "protocol": &schema.Schema{
            Type:     schema.TypeString,
            Required: true,
        },
        "sslcheck": &schema.Schema{
            Type:     schema.TypeBool,
            Optional: true,
        },
      },
      ResourcesMap: map[string]*schema.Resource{},
      DataSourcesMap: map[string]*schema.Resource{},
      ConfigureFunc: configureProvider,
    }
  }
/*
func Provider() *schema.Provider {
        return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apiKey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
            "bearer": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
            "endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
            "protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
            "sslcheck": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
            
		},
        ResourcesMap: map[string]*schema.Resource{
                        //"tower_organization": resourceOrganization(),
                        
                        //"tower_credential":resourceCredential(),
                },   
        DataSourcesMap: map[string]*schema.Resource{
                    
                },
        
		ConfigureFunc: configureProvider,
        }
}
*/
func configureProvider(d *schema.ResourceData) (interface{}, error) {
    apiKey:=d.Get("apiKey").(string)
    bearer:=d.Get("bearer").(string)
    endpoint:=d.Get("endpoint").(string)
    protocol:=d.Get("protocol").(string)
    sslCheck:=d.Get("sslcheck").(bool)
    client :=CosWeb{apiKey,bearer,endpoint,protocol,sslCheck}
    log.Printf("[INFO] Initializing Tower Client")
    return client,nil
}
