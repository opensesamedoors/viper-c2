class Operator {
  constructor(url, username, password){
    this.server = url+"/operator"
    this.username = username
    this.password = password
  }

  async login(){
    let res = await fetch(this.server+"/login", { 
      method: "POST", 
      credentials: "include",
      body: JSON.stringify({
        "username": this.username,
        "password": this.password
      })
    })
    res = await res.json()

    if (res["error"] != "")
      throw res["error"] 
  }

  async get_operators() {
    let res = await fetch(this.server+"/operators", {
      method: "GET",
      credentials: "include", 
    })
    res = await res.json()

    if (res["error"] != "")
      throw res["error"]

    return res["result"]
  }

  async get_result(uid) {
    let res = await fetch(this.server+"/results?uid="+uid, {
      method: "GET", 
      credentials: "include", 
    })
    res = await res.json()
   
    if (res["error"] != "")
      throw res["error"]

    return res["result"] 
  }

  async get_tokens() {
    let res = await fetch(this.server+"/tokens", {
      method: "GET", 
      credentials: "include", 
    })
    res = await res.json()
   
    if (res["error"] != "")
      throw res["error"]

    return res["result"]
  }

  async get_agents() {
    let res = await fetch(this.server+"/agents", {
      method: "GET", 
      credentials: "include", 
    })
    res = await res.json()
   
    if (res["error"] != "")
      throw res["error"]

    return res["result"]
  }

  async get_jobs() {
    let res = await fetch(this.server+"/jobs", {
      method: "GET", 
      credentials: "include", 
    })
    res = await res.json()
   
    if (res["error"] != "")
      throw res["error"]

    return res["result"]
  }

  async get_version(){
    let res = await fetch(this.server+"/version", {
      method: "GET", 
      credentials: "include", 
    })
    res = await res.json()
   
    if (res["error"] != "")
      throw res["error"]

    return res["result"] 
  }

  async del_token(token){
    let res = await fetch(this.server+"/tokens?token="+token, {
      method: "DELETE", 
      credentials: "include", 
    })
    res = await res.json()

    if (res["error"] != "")
      throw res["error"]

    return res["result"] 
  }

  async create_token(){
    let res = await fetch(this.server+"/tokens", {
      method: "POST", 
      credentials: "include", 
    })
    res = await res.json()

    if (res["error"] != "")
      throw res["error"]

    return res["result"] 
  }

  async add_job(action, data, target){
    let res = await fetch(this.server+"/job", {
      method: "POST", 
      credentials: "include", 
      body: JSON.stringify({
        "action": action,
        "data": data,
        "target": target
      })
    })
    res = await res.json()

    if (res["error"] != "")
      throw res["error"]

    return res["result"] 
  }

  async logout(){
    let res = await fetch(this.server+"/logout", {
      method: "GET",
      credentials: "include",
    })
    res = await res.json()

    if (res["error"] != "")
      throw res["error"]

    this.token = ""
  }
}
