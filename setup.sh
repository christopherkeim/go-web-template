#! /bin/bash
# 
# This script sets up a local development environment on an Ubuntu 20.04/22.04 machine
# to work with go1.22.0 projects. 
# 
# Targets:
#   - go1.22.0
#   - Docker 24.0.6
#
# Requirements:
#   - Ubuntu 20.04/22.04
#

# -----------------------------------------------------------------------------------------------------------
# 0) Config: here we'll set default variables and handle options for controlling 
#    the target versions of Golang.
# -----------------------------------------------------------------------------------------------------------

# Pull the current machine's distro for GPG key targeting
readonly DISTRO="$(lsb_release -d | awk -F ' ' '{print tolower($2)}')"

# Pull the machine's chip architecture
if [[ "$(uname -m)" == "x86_64" ]]; then
  CHIP_ARCH="amd64"
else
  CHIP_ARCH="arm64"
fi

readonly CHIP_ARCH

# -----------------------------------------------------------------------------------------------------------
# 1) Base Requirements: this will ensure that you base requirements along with some data science
#    CLI tools installed.
# -----------------------------------------------------------------------------------------------------------

# Check if ca-certificates is in the apt-cache
if ( apt-cache show ca-certificates > /dev/null ); then
  echo 'ca-certificates is already cached 🟢'
else
  sudo apt update
fi

# Ensure ca-certificates package is installed on the machine
if ( which update-ca-certificates > /dev/null ); then
  echo 'ca-certificates is already installed 🟢'
else
  echo 'Installing ca-certificates 📜'
  sudo apt-get install -y ca-certificates
fi

# Check if curl is in the apt-cache
if ( apt-cache show curl > /dev/null ); then
  echo 'curl is already cached 🟢'
else
  sudo apt update
fi

# Ensure curl is installed on the machine
if ( which curl > /dev/null ); then
  echo 'curl is already installed 🟢'
else
  echo 'Installing curl 🌀'
  sudo apt install -y curl
fi

# Check if make is in the apt-cache
if ( apt-cache show make > /dev/null ); then
  echo 'make is already cached 🟢'
else
  sudo apt update
fi

# Ensure make is installed on the machine
if ( which make > /dev/null ); then
  echo 'make is already installed 🟢'
else
  echo 'Installing make 🔧'
  sudo apt install -y make
fi

# Check if gnupg is in the apt-cache
if ( apt-cache show gpg > /dev/null ); then
  echo 'gnupg is already cached 🟢'
else
  sudo apt update
fi

# Ensure gnupg is installed on the machine
if ( which gpg > /dev/null ); then
  echo 'make is already installed 🟢'
else
  echo 'Installing gnugp 🔧'
  sudo apt install -y gnupg
fi

# Check if bat is in the apt-cache
if ( apt-cache show bat > /dev/null ); then
  echo 'batcat is already cached 🟢'
else
  sudo apt update
fi

# Ensure bat is installed on the machine
if ( which batcat > /dev/null ); then
  echo 'batcat is already installed 🟢'
else
  echo 'Installing batcat 🔧'
  sudo apt install -y bat
fi

# Check if jq is in the apt-cache
if ( apt-cache show jq > /dev/null ); then
  echo 'jq is already cached 🟢'
else
  sudo apt update
fi

# Ensure jq is installed on the machine
if ( which jq > /dev/null ); then
  echo 'jq is already installed 🟢'
else
  echo 'Installing jq 🔧'
  sudo apt install -y jq
fi


# -----------------------------------------------------------------------------------------------------------
# 2) Go Install: here we'll install and configure Go
# -----------------------------------------------------------------------------------------------------------

# Install Go
if ( which go > /dev/null ); then
  echo "Go is already installed 🟢"
else
  echo "Installing Go 🦫"
  wget -O https://go.dev/dl/go1.22.0.linux-${CHIP_ARCH}.tar.gz
  sudo tar -C /usr/local -xzf go1.22.0.linux-${CHIP_ARCH}.tar.gz
  rm go1.22.0.linux-${CHIP_ARCH}.tar.gz
fi

# Add Go to PATH
if ( cat ~/.bashrc | grep "go/bin" > /dev/null ); then
  echo 'Go is already in path 🟢'
else
  echo 'Adding go to PATH 🔧'
  echo -e "# Add Go to PATH\nexport PATH="/usr/local/go/bin:$PATH"" >> ~/.bashrc
  source ~/.bashrc
fi

# Verify installation of Go
if ( go version > /dev/null ); then
  echo "$(go version) 🦫 🚀"
else
  echo "Go was not installed successfully 🔴"
fi


# -----------------------------------------------------------------------------------------------------------
# 3) Docker Install: here we'll install Docker
# -----------------------------------------------------------------------------------------------------------

# -----------------------------------------------------------------------------------------------------------
# 3.1) Set up the repository: Before you install Docker Engine for the first time on a new host machine, 
# you need to set up the Docker repository. Afterward, you can install and update Docker from the repository.
# -----------------------------------------------------------------------------------------------------------

# Add Docker’s official GPG key
if [[ -f /etc/apt/keyrings/docker.gpg ]]; then
  echo 'Docker GPG Key already installed at /etc/apt/keyrings/docker.gpg 🟢'
else
  echo 'Installing Docker GPG Key at /etc/apt/keyrings/docker.gpg 🔧'
  
  # Create the /etc/apt/keyrings directory with appropriate permissions
  sudo install -m 0755 -d /etc/apt/keyrings
  
  # Download the GPG key from Docker
  curl -fsSL https://download.docker.com/linux/${DISTRO}/gpg \
    | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

  sudo chmod a+r /etc/apt/keyrings/docker.gpg
fi

# Set up the repository
if [[ -f /etc/apt/sources.list.d/docker.list ]]; then
  echo 'docker.list repository already exists at /etc/apt/sources.list.d/docker.list 🟢'
else
  echo 'Installing docker.list repository at /etc/apt/sources.list.d/docker.list 🔧'
  echo \
    "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] \
    https://download.docker.com/linux/$DISTRO \
    "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" \
    | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
fi

# -----------------------------------------------------------------------------------------------------------
# 3.2) Install Docker Engine
# -----------------------------------------------------------------------------------------------------------

# Check if docker-ce is in the apt-cache
if ( apt-cache show docker-ce > /dev/null ); then
  echo 'docker-ce is already cached 🟢'
else
  sudo apt update
fi

# Install Docker Engine, containerd, and Docker Compose
if ( docker --version > /dev/null ); then
  echo 'Docker is already installed 🟢'
  echo "Using $(docker --version)"
else
  echo 'Installing Docker 🐳'

  # Installs
  sudo apt-get install -y docker-ce \
    docker-ce-cli \
    containerd.io \
    docker-buildx-plugin \
    docker-compose-plugin
  
  # Verify that the Docker Engine installation is successful by running the hello-world image
  sudo docker run --rm hello-world
fi
