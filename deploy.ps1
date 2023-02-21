# Deploy API from local repo

# Must be logged in to GCP locally

# TODO Promt for creds if not already logged in

gcloud functions deploy go-http-function `
    --gen2 `
    --runtime=go120 `
    --region=us-west2 `
    --source=. `
    --entry-point=HelloGet `
    --trigger-http `
    --allow-unauthenticated