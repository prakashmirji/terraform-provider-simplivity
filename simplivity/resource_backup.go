package simplivity

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBakcup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceBackupRead,
		Create: resourceBackupCreate,
		Update: resourceBackupUpdate,
		Delete: resourceBackupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"consistency_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datastore_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_deletion_tm": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vm_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_consistent": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"comp_clust_parent_hyp_obj": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"omni_stack_clust_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sent_completion_tm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cluster_grp_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:      schema.HashString,
				Computed: true,
			},
			"unique_size_tm_stamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiration_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"omnistack_cluster_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sent": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vm_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datastore_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comp_clust_parent_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hyp_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sent_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceBackupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	backups, err := config.ovcClient.Backups.GetByName(id)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(id)
	d.Set("name", backups.Name)
	d.Set("id", backups.Id)
	d.Set("vm_name", backups.VirtualMachineName)
	d.Set("created_at", (backups.CreatedAt).String())
	d.Set("consistency_type", backups.ConsistencyType)
	d.Set("type", backups.Type)
	d.Set("datastore_name", backups.DatastoreName)
	d.Set("vm_deletion_tm", string((backups.VirtualMachineDeletionTime)))
	d.Set("vm_id", backups.VirtualMachineID)
	d.Set("app_consistent", backups.ApplicationConsistent)
	d.Set("comp_clust_parent_hyp_obj", backups.ComputeClusterParentHypervisorObjectID)
	d.Set("state", backups.State)
	d.Set("omni_stack_clust_id", backups.OmnistackClusterID)
	d.Set("vm_type", backups.VirtualMachineType)
	d.Set("sent_completion_tm", backups.SentCompletionTime)
	d.Set("unique_size_bytes", backups.UniqueSizeBytes)
	d.Set("cluster_grp_ids", backups.ClusterGroupIds)
	d.Set("unique_size_tm_stamp", backups.UniqueSizeTimestamp)
	d.Set("expiration_time", backups.ExpirationTime)
	d.Set("omnistack_cluster_name", backups.OmnistackClusterName)
	d.Set("sent", backups.Sent)
	d.Set("size", backups.Size)
	d.Set("vm_state", backups.VirtualMachineState)
	d.Set("datastore_id", backups.DatastoreID)
	d.Set("comp_clust_parent_name", backups.ComputeClusterParentName)
	d.Set("hyp_type", backups.HypervisorType)
	d.Set("sent_duration", backups.SentDuration)
	return nil
}

func resourceBackupCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceBackupRead(d, meta)
}

func resourceBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceBackupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
