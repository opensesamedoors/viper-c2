# Viper | JavaScript API Library 
JavaScript API library can be used to interact with a viper web server,
**this libary only supports operator features**

## Setup
To use the viperjs library, link it to your HTML before your actual JS:
```html
<script type="text/javascript" src="https://github.com/ngn13/viper/raw/main/lib/javascript/viper.js"></script>
```

## Examples
### Example Operator Client 
```html
<textarea id="area"></textarea>
<br>
<button id="version">Get Server Version</button>

<script type="text/javascript" src="https://github.com/ngn13/viper/raw/main/lib/javascript/viper.js"></script>
<script>
  const version = document.getElementById("version")
  const area = document.getElementById("area")

  op = new Operator("http://localhost:8080", "ngn", "changeme")

  version.addEventListener("click", async ()=>{
    await op.login()
    area.value = await op.get_version()  
    await op.logout()
  })
</script>
```
