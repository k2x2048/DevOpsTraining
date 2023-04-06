
# Ansible

Type of challenge: **learning** </br>
Duration: **5 days** </br>
Team challenge : **solo**

## The mission
Deploying and managing servers reliably and efficiently has always been a challenge. Previously, system administrators managed servers by hand, installing software, changing configurations, and administering services on individual servers. IT operations were spending more time configuring the systems. That’s why server provisioning and configuration management tools came to flourish. One of these tools is Ansible. <br /><br />
Ansible is a radically simple IT automation engine that automates cloud provisioning, configuration management, application deployment, intra-service orchestration, and many other IT needs.

In order to learn the fundamental concepts behind Ansible:

1. Start with the basics: Before diving into Ansible, make sure you have a good understanding of the basics of system administration, such as managing servers, networking, and security.

2. Learn the core concepts: Ansible is based on a set of core concepts, including inventory, playbooks, modules, and roles. Make sure you understand what these are and how they work.

3. Practice, practice, practice: The best way to learn Ansible is by using it. Set up a test environment and start writing playbooks for simple tasks.

4. Join the community: Ansible has a large and active community of users who are happy to help and answer questions. Join forums, mailing lists, and social media groups to learn from others and share your experiences.

Here are some learning objectives that you may want to consider as you begin your journey with Ansible:

1. Understand the core concepts and terminology of Ansible
2. Install and configure Ansible on your machine
3. Create and manage an inventory of hosts
4. Write and run simple Ansible playbooks
5. Use Ansible modules to manage files, users, packages, and services
6. Use variables and templates to make your playbooks more flexible
7. Create and use Ansible roles to organize your playbooks
8. Use Ansible to manage configuration files and services on multiple servers
9. Automate deployment and scaling of applications using Ansible
10. Understand best practices for using Ansible in production environments

Your mission will be to write Ansible playbooks that will allow you to deploy a software solution and it's configuration (Ex: Gitlab, Apache Kafka, Eclipse Mosquitto, ...). By doing so, you will have to learn to do things like :

- Manage user accounts (Create)
- Install packages
- Provision config files
- Handle security (ssh login via Private Key only) & configure firewall
- Install docker and deploy what ever container(s) was previously deployed on the original VM.
- Restore your dotfiles (.bashrc, .bash_aliases, .vimrc, etc...)

And don't forget, you can use the command "ansible-doc -l" to have a complete list of the default modules that comes with ansible.

By the end of the challenge, you must be able to delete and re-deploy you new VM by running one single Ansible Playbook.


### IMPORTANT NOTE:


- At some point you will have to work with sensitive data, such as password, ssh keys, etc... Please make sure that theses are stored in a secure manner (encrypted) using Ansile Vault. 

## RESOURCES

- [https://docs.ansible.com/ansible/2.9/](https://docs.ansible.com/ansible/2.9/)