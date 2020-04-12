provider "simplivity" {
    ovc_username = ""
    ovc_password = ""
    ovc_host_address = ""
    ovc_cert_path = ""
}

data "svt_backup_data"  "backup" {
    name = "abc"
}