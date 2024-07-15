resource "mist_org_wxtag" "wtag_one" {
  org_id = mist_org.terraform_test.id
  name   = "wtag_one"
  values = [
    "10.3.0.0/16"
  ]
  op    = "in"
  type  = "match"
  match = "ip_range_subnet"
}