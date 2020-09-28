# Local Deployment

Running a Ansible playbook locally.

```sh
ansible-playbook .backpack/digitalocean/playbooks/setup.yml -i etc/ansible/hosts -u root --key-file ~/path/to/ssh/private_key
```

Deploying an aplication to dokku locally.

```sh
git remote add dokku dokku@11.22.33.44:app-name
GIT_SSH_COMMAND="ssh -i /path/to/ssh/key -l dokku" git push dokku master -f
```
