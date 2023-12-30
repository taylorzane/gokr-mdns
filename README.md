# gokr-mdns

Advertise `gokrazy.local` through mDNS from rasp pi set up through gokrazy.
Useful since for some reason I cannot reach gokrazy pi's just through their hostname.

## Setup
```
gokr-packer github.com/taylorzane/gokr-mdns
```
populate `GOKRAZY_UPDATE` to update existing rasp pi on the network.

Use local modifications (fish shell expansion)
```
cd builddir/github.com/taylorzane/gokr-mdns/
go mod edit -replace github.com/taylorzane/gokr-mdns=(realpath ../../../..)
```

## Customizing

The host, domain, and interfaces can be configured by setting command line flags in your gokrazy config.

```json
{
  "Hostname": "my-gokrazy-service.local",
  "PackageConfig": {
    "github.com/taylorzane/gokr-mdns": {
      "CommandLineFlags": [
        "-host",
        "my-gokrazy-service",
        "-domain",
        "local",
        "-ifaces",
        "eth0"
      ]
    }
  }
}
```