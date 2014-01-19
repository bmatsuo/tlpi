##The Linux Programming Interface (No Starch Press)

Code from the book as well as excercise implementations and, probably,
code from other languages like Go and Rust.

###Contents

- `tlpi-dist` Fully annotated source code of The Linux Programming Interface.

###Setup

Get virtual box and a disk image for your desired linux distribution.
This setup was performed with Ubuntu 13.04 (Raring Ringtail).

Create an Linux (Ubuntu) VM with all the default configurations. Once created
open the virtual machine, select the disk image containing the linux
distribution, and start the vitual machine. Guided installation should be
adequate.

Install `git` and `sshd`

	sudo apt-get install git ssh

In order to SSH into the VM open the its Network settings and change Adapter 1
into a Bridged Adapter. The restart the VM.

	sudo shutdown -r now

You will now be able to get a locally-visible IP address for the vm

	ifconfig eth0

Finally ssh into the VM, forward the user-agent from your host machine, clone
the repository, and build things...

	ssh XXX.XXX.XXX.XXX
	git clone git@github.com:bmatsuo/tpli.git
	cd tpli
	./boom
