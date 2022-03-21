# gh-repo

Simple CLI program that gets the open issues and their associated labels from Public GitHub Repositories and returns them in JSON format.

## Installation

### Static Binaries

Download the latest release from the [GitHub Release Page](https://github.com/leb4r/gh-repo/releases/latest) based on your system.

### From Source

Clone this repository and use go to build a binary specific for your system.

```bash
git clone https://github.com/leb4r/gh-repo.git
cd gh-repo
go install
```

## Usage

There is only a single required argument: `-repo`. This must be the canonical name of a GitHub repository in the form of: `owner/repo`.

e.g. for the repository [leb4r/semtag](https://github.com/leb4r/semtag):

```bash
gh-repo -repo leb4r/semtag
```
