# Enmity-store

Script used to generate a website front-end for your plugins.  
Usage:

```shell
go run cmds/store/main.go
  -dir string
        Your repository's folder.
        
  -name string
        Your repo's name. (default "Your plugin repo")
```

The script will generate an index.html and info.json file according to the content of the `plugins/` folder found in your repository's folder.  
The layout of your store should look like this:

```none
  repo:
    plugins:
      TestPlugin.js
      TestPlugin.json
    index.html
    info.json
```
