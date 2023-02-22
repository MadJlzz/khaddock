# kubeconf

Kubeconf is focused on providing declarative APIs and tooling to simplify configuration VMs (_or agents_) by only declaring states, providing a more secure and maintanable fleet of machines.

## Why?

This idea emerged after I've been confronted with a very simple problem. 

> How do you deploy hundreds or thousands of VMs on a lot of different providers while keeping a configuration consistency between all of them and shipping
> patches very fast?

A lot of tools exists to perform such tasks but they all have their own specific downsides conducting you to give up on either maintainability, ease of use or even speed. Or maybe I am realy a bad Googler 🙃

### What about Terraform?

Terraform is great for _provisionning_ any kind of ressources such as VMs for _providers_ that build their own **terraform provider abstraction**. Sadly, not all of the providers I am using developed their own abstraction for me to simply create new machines a configure them through an `init script` or even a `cloud-init` configuration file.

I could develop the abstraction for each missing provider but:
  - I do not have control on API changes meaning that if the provider is unoficial, most likely they won't care my code breaks
  - If I have to handle a new IaS provider, I have to start the work again if they don't do it
  - None are part of the OSS ecosystem so why bother?

### What about Ansible?

Ansible is great for _configuration management_ which feels like one of the perfect tool for this task. This is particularly true when you run `playbooks` on an _okay-ish_ inventory size. By default, Ansible's strategy run each tasks one by one ; having one error (e.g. a network error) be responsible of the `playbook` failure for all the other machines.

I know that Ansible provides also a `free` strategy that tries to execute as fast as it can (up to the maximum number of **forks** you define) tasks but that don't really solve the issue of having to configure a thousand machines (unless you want to put $$$ for configuring VMs which seems _meh_ to me)

To finish, Ansible do not handle a **state** meaning that even with no change, it will still go through everything again. Not very practical if you only wanted to patch OpenSSL e.g.

### What about Puppet or Chef?

They represent pretty much what I had in mind when thinking of a `state` or a `recipe`. So what's the problem?

I think we all agree saying that both of these tools are HARD to master and to deploy and I am not talking of PuppetDSL and RubyDSL which to me, are just a nightmare to work with.
