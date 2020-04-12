provider "simplivity" {
    ovc_username = ""
    ovc_password = ""
    ovc_host_address = ""
    ovc_cert_path = ""
}

data "simplivity_backup_data"  "backup" {
    name = "abc"
}
resource "simplivity_backup_resource"  "backup" {
    name = "abc"
}
