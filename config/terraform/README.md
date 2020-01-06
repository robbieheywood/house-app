# Terraform

Terraform is used to setup the cloud infrastructure we use.
The state is stored in Terraform cloud.

## How to use
* Ensure you have a terraform cloud token at `~/.terraformrc`.
  * This is used to allow access to the remote state in Terraform cloud 
  * If not, a token will need to be generated from within Terraform cloud.
* Run `terraform init $REPO/config/terraform` to setup terraform.
* After changes, run `terraform plan $REPO/config/terraform` to check the changes.
* Then run `terraform apply $REPO/config/terraform` to make the changes.