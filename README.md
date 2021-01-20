# checkov2vim

Translate [Checkov](https://github.com/bridgecrewio/checkov) error messages
to a Vim friendly format.

## Usage

Pipe the `checkov` results through.

```
$ checkov -f main.tf| checkov2vim
main.tf:95: CKV_AWS_79 Ensure Instance Metadata Service Version 1 is not enabled https://docs.bridgecrew.io/docs/bc_aws_general_31
```

## Using it with Ale

If you are an [Ale](https://github.com/dense-analysis/ale) person, you can
generate the required configuration via this binary, for example

```
$ checov2vim generate
```

This will create the required vimscript in your path (see
`checov2vim generate -h` for more info)
