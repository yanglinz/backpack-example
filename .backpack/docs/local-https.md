# Local Development with HTTPS

The closer we can replicate the actual production runtime in development, the
better. This document describes the setup we need to run local development
server with `https`.

## Setup DNSMasq

Dnsmasq is a small DNS server that we can run locally to support wildcard domain
resolution.

```sh
brew install dnsmasq
```

Once installed we can create a configuration file.

```sh
mkdir -pv "$(brew --prefix)/etc/"
touch "$(brew --prefix)/etc/dnsmasq.conf"
```

Put the following content inside the newly created `dnsmasq.conf`.

```
address=/.localhost/127.0.0.1
port=53

# Don't read /etc/resolv.conf or any other configuration files.
no-resolv
# Never forward plain names (without a dot or domain part)
domain-needed
# Never forward addresses in the non-routed address spaces.
bogus-priv
```

Once the configuration file is in place, we can start the DNS service.

```sh
brew services start dnsmasq
```

With the local DNS server in place, the last step is to integrate the resolver.

```sh
sudo mkdir -pv "/etc/resolver"
sudo bash -c 'echo "nameserver 127.0.0.1" > /etc/resolver/localhost'
```

This will allow us to reach have wildcard localhost domains.

## Setup Certificates

Now the we have the wildcard DNS setup, the next and final piece is provisioning
the certificates using `mkcert`.

```sh
brew install mkcert
brew install nss # For firefox
```

Once `mkcert` is installed, we can generate wildcard certificates for local
development.

```sh
mkcert -install
```

> NOTE: Backpack will take care of generating the application specific
> certificate for you. We just need to install the root CA once.

And now we should have wildcard `https://app-name.localhost` domain setup for
development.
