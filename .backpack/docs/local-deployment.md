# Local Deployment

Running ansible locally.

```sh
ansible-playbook .backpack/digitalocean/playbooks/setup.yml -i etc/ansible/hosts -u root --key-file ~/path/to/ssh/private_key
```
