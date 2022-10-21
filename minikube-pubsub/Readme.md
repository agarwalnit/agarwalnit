# Overview
This document provide an example of calling a GCP service from minikube pod. 

# Set credentials and enable gcp auth add on

```
# set gcp project id

gcloud config set project $YOUR_PROJECT_ID # For ex: agarwalnitin-test

# Get GCP creds
gcloud auth application-default login

# Enable GCP auth add on in minikube
minikube addons enable gcp-auth --force
```

# Run 
Run using `slaffold run` 