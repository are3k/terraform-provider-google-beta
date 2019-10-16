// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeRegionTargetHttpProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionTargetHttpProxyCreate,
		Read:   resourceComputeRegionTargetHttpProxyRead,
		Update: resourceComputeRegionTargetHttpProxyUpdate,
		Delete: resourceComputeRegionTargetHttpProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionTargetHttpProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url_map": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeRegionTargetHttpProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionTargetHttpProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionTargetHttpProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	urlMapProp, err := expandComputeRegionTargetHttpProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	regionProp, err := expandComputeRegionTargetHttpProxyRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionTargetHttpProxy: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionTargetHttpProxy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating RegionTargetHttpProxy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionTargetHttpProxy: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating RegionTargetHttpProxy %q: %#v", d.Id(), res)

	return resourceComputeRegionTargetHttpProxyRead(d, meta)
}

func resourceComputeRegionTargetHttpProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionTargetHttpProxy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionTargetHttpProxyCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionTargetHttpProxyDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}
	if err := d.Set("proxy_id", flattenComputeRegionTargetHttpProxyProxyId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionTargetHttpProxyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}
	if err := d.Set("url_map", flattenComputeRegionTargetHttpProxyUrlMap(res["urlMap"], d)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionTargetHttpProxyRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpProxy: %s", err)
	}

	return nil
}

func resourceComputeRegionTargetHttpProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("url_map") {
		obj := make(map[string]interface{})

		urlMapProp, err := expandComputeRegionTargetHttpProxyUrlMap(d.Get("url_map"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
			obj["urlMap"] = urlMapProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}/setUrlMap")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating RegionTargetHttpProxy %q: %s", d.Id(), err)
		}

		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating RegionTargetHttpProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("url_map")
	}

	d.Partial(false)

	return resourceComputeRegionTargetHttpProxyRead(d, meta)
}

func resourceComputeRegionTargetHttpProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionTargetHttpProxy %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionTargetHttpProxy")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting RegionTargetHttpProxy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionTargetHttpProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionTargetHttpProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/targetHttpProxies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionTargetHttpProxyCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionTargetHttpProxyDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionTargetHttpProxyProxyId(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionTargetHttpProxyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionTargetHttpProxyUrlMap(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionTargetHttpProxyRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeRegionTargetHttpProxyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpProxyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpProxyUrlMap(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("urlMaps", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionTargetHttpProxyRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
