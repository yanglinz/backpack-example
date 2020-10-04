# Getting Started

To get started with adopting `backpack`, run the following commands in your repository.

First, add the backpack subtree:

```sh
git remote add -f backpack git@github.com:yanglinz/backpack.git
git subtree pull --prefix=.backpack backpack master --squash
```

Then, we can create relevant files by running:

```
./backpack setup
./backpack setup --resources
```

Next, to setup Github Actions properly, we'll need to populate the following secrets:

- `DIGITALOCEAN_PRIVATE_KEY` - SSH private key pre-authorized for hosted VMs.
- `DIGITALOCEAN_TOKEN` - Authentication token for [https://www.digitalocean.com/](https://www.digitalocean.com/).
- `GCP_SERVICE_ACCOUNT_KEY` - `base64`'d GCP service account JSON.
- `TERRAFORM_CLOUD_TOKEN` - Authentication token for [https://terraform.io](https://terraform.io).
