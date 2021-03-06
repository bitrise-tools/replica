# replica

Create a CI environment (with Docker / virtual machine)

_Based on the awesome [https://github.com/timsutton/osx-vm-templates](https://github.com/timsutton/osx-vm-templates) project,
and [https://spin.atomicobject.com/2015/11/17/vagrant-osx/](https://spin.atomicobject.com/2015/11/17/vagrant-osx/) blog post._


## Install

Check the [releases](https://github.com/bitrise-io/replica/releases) page for install instructions.

The `replica` binary is a stand alone binary, which includes every resource
it uses, so the only tools you have to install are:

- [VirtualBox](https://www.virtualbox.org)
- [vagrant](https://www.vagrantup.com)
- [packer](https://www.packer.io)

_Once these are installed, you just have to download the replica binary on macOS
and run it. That's it, there's no installation process for `replica`,
just download the binary and run it._


## Supported features

### `replica create ...`

Creates a ready-to-use `vagrant` box automatically. The only required input
is a macOS Installer (app), which you can download for free from the Mac App Store.

Once you have that, just run:

```
replica create '/path/to/macOS Installer.app'
```

For example, if you downloaded macOS Sierra from the Mac App Store, into the `/Applications` directory:

```
replica create '/Applications/Install macOS Sierra.app'
```

This process consists of two steps:

1. creating the base, auto installing `dmg` mac os installer, which can be used by `packer` / VirtualBox
2. using the installer, creating the `vagrant` box, using `packer`

If all you want is the creation of the auto installer `dmg`, you'll be able to
stop before step 2 - `replica` will ask you whether you want to proceed to actually create the
`vagrant` box.

__Step 1 takes about 3 mins and requires about 20 GB free disk space in total__,
from which the created DMG will take ~5.5 GB,
and an additional ~15 GB free disk space will be used during the creation
for temporary files, which are deleted when the step successfully completes.
After the `create` command finishes feel free to move the created
DMG file to an external hard drive.

__Step 2 takes about 35-40 mins and requires about 25 GB free disk space in total__,
from which the created `box` file will take ~9 GB,
and an additional ~17 GB free disk space will be used during the creation
for temporary files, which are deleted when the step completes.
After the `create` command finishes feel free to move the created
`vagrant` `box` file to an external hard drive.


### `replica create vagrant`

Creates and boots a `vagrant` VM, from a `vagrant` box,
in the directory you specify.

__This step takes about 4 mins and requires about 20 GB free disk space in total without `Xcode.app` sync__,
if you allow `Xcode.app` to be synced into the VM that will take
an additional ~15 mins and ~10 GB disk space (the size of `Xcode.app`).


## Links

* [Developing on OS X Inside Vagrant - automated MacOS vagrant box creation](https://spin.atomicobject.com/2015/11/17/vagrant-osx/)
    * Using: https://github.com/timsutton/osx-vm-templates
* [New Adventures in Automating OS X Installs with startosinstall](https://macops.ca/new-adventures-in-automating-os-x-installs-with-startosinstall)

## InstallDMG location:

```
"Install OS X..." / "Install macOS ..." app -> Show package content -> Contents -> SharedSupport -> InstallESD.dmg
```


## Tested tool versions

You can find the tool versions we tested with in the `TESTED_TOOL_VERSIONS.md`
file. __Feel free to add your tool versions report to the list!__


## Development

### Embedded resources / the `resources` directory

__If you change any resource in the `resources` directory, you should run
the resource embed process!__

You can do that by running: `bitrise run embed-resources`


## TODO

- auto add the created box into vagrant
- tool versions: auto save into file if create is successful
- save vagrant box into _out, and maybe expose commands to only do parts (prep | packer)
- delete tmp dir, unless error or flag passed
- password: try to write it into file from string

- elimintate `cd`s - generate the files right where it have to be
- annotate the code, based on the original
- remove `veewee` from everywhere

- add support for `docker` / the Linux environment


