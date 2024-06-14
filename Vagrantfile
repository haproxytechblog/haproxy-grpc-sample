Vagrant.configure("2") do |config|

    config.vm.define "grpc" do |server|
      server.vm.box = "debian/bookworm64"
      server.vm.hostname = "grpc"
      server.vm.network "private_network", ip: "192.168.56.6"
      server.vm.provider "virtualbox" do |v|
        v.memory = 2048
        v.cpus = 2
        v.linked_clone = true
      end
    end

  
  end
  