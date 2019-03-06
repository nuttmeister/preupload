# preupload

Small program for generating pre-signed URLs for S3 PutObject.

## Pre-requisities

You will need go installed fo build and install this program.  
Get it at [https://golang.org/dl/](https://golang.org/dl/).

Your `$GOPATH/bin` folder must be in your $PATH.

## Installation

```bash
go install
```

## Usage

```bash
preupload -bucket <MYBUCKET> -key <MYKEY> -expire <EXPIRE>
```

Where `<MYBUCKET>` is the bucket where to upload to and `<MYKEY>` the file
and `<EXPIRE>` the expirey time of the pre-signed url in minutes.

## Example

The example below will generate a url for uploding a file to my-cool-bucket/test.zip that
is valid for 60 minutes.

```bash
preupload -bucket my-cool-bucket -key test.zip -expire 60
curl https://url-from-preupload-program --upload-file /path/to/myFile.zip
```