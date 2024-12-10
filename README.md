# cve2json

A golang application that demonstrates how to convert [TuxCare CVE data](https://cve.tuxcare.com/els/cve) from CSV to JSON.

## Running from go

You can run the source script directly from go, for example:

```bash
go run main.go
```

## Building binaries

The included Makefile can build various binaries, subject to your go environment setup. The easiest invocation is:

```bash
make clean all
```

Alternatively you can compile and install your platform's binary (there are no dependencies) into `$GOPATH/bin/` using:

```bash
go install github.com/sej7278/cve2json@latest
```

Or simply download a binary from the [releases](https://github.com/sej7278/cve2json/releases) page.

## Usage

By default the program only returns `Released` status CVE's, however if you call it like `cve2json --all` it will return all statuses including `Needs Triage`, `Ignored`, `Not Vulnerable` etc.

## Example output

The ESU and FIPS repository data is merged, so you'll sometimes see two fixes for the same CVE, hence why the CVE number can't be used as a key:

```json
[
  {
    "CVE": "CVE-2024-50073",
    "Last updated": "2024-11-18 16:33:07.130913",
    "OS name": "AlmaLinux 9.2 FIPS",
    "Project name": "kernel",
    "Score": "7.8",
    "Severity": "HIGH",
    "Status": "Released",
    "Version": "5.14.0"
  },
  {
    "CVE": "CVE-2024-50073",
    "Last updated": "2024-11-18 16:33:06.14445",
    "OS name": "AlmaLinux 9.2 ESU",
    "Project name": "kernel",
    "Score": "7.8",
    "Severity": "HIGH",
    "Status": "Released",
    "Version": "5.14.0"
  },
]
```
