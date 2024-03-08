# GCP Functions with Terraform

GCP function demo with terraform.

## Setup
```bash
gcloud auth application-default login

export GOOGLE_CLOUD_PROJECT=PROJECT_ID

terraform init

terraform apply
```

## Local Run
```bash
export FUNCTION_TARGET=router

cd cmd
go run main.go
```

##

Tutorial: https://cloud.google.com/functions/docs/create-deploy-http-go

