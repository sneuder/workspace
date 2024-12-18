# Workspace

Workspace is a CLI tool that streamlines container management for developers working within Docker environments. With Workspace, you can set a default container for quick access, add or remove containers from your workspace, and execute any command on the default container directly, all from a single interface. This tool is designed to simplify interactions with containers, making it easier to configure and manage your development environment efficiently.

## Steps for installation

**0. Install dependencies**

This tool uses docker, it is important to have it in your local machine.

**1. Download the executable file**
```
wget https://github.com/sneuder/workspace/releases/download/v1.0.0/install.sh
```

**2. Make the Script Executable**
```
sudo chmod +x install.sh
```

**3. Run the Script**
```
sudo ./install.sh
```

## Commands

- Add container
```
wkspace add <container_name>
```

- List tig organization
```
wkspace show
```

- Set default container
```
wkspace default <organization_name>
```

- Remove container
```
wkspace remove <organization_name>
```

- Any comand within container's terminal
```
wkspace ls
```

```
wkspace ls -l
```

```
wkspace git status
```
