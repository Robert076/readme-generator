# 🌎 readme-generator 

README file generator for writing your documentation in a faster way.


> ❗️ This project is still in a beta version and currently doesn't have many functionalities.


### 🚀 Running the project: (locally)
1. Run the main package:
```bash
go run cmd/main.go
```

2. **(Postman)** Make a GET request to [localhost:8080/readme](localhost:8080/readme) 

3. **(Postman)** Make a POST request to [localhost:8080/readme](localhost:8080/readme) with a body just like the one you received from the GET.

---

### ⚙️ Project structure 
```
.
├── 📋 LICENSE
├── 🏁 README.md                       # Project documentation
├── cmd
│   └── 💻 main.go                     # Entrypoint to the application
├── ⚙️ go.mod                          # Go module definition
├── ⚙️ go.sum                          # Go module dependency checksums
└── internal
    ├── api
    │   ├── 🌐 endpoints.go            # HTTP handlers for GET/POST /readme
    │   └── 🧩 format.go               # Sample format returned by GET /readme
    ├── blocks
    │   ├── bash_command
    │   │   ├── 🧱 bash_command.go     # Renders bash command markdown blocks
    │   │   └── 📂 templates           # Bash command block templates
    │   ├── 🧱 blocks.go               # Shared interface for all markdown blocks
    │   ├── description
    │   │   ├── 🧱 description.go      # Renders project description block
    │   │   └── 📂 templates           # Description block templates
    │   └── title
    │       ├── 📂 templates           # Title block templates
    │       └── 🧱 title.go            # Renders project title block
    ├── generator
    │   └── 🏗️ generator.go            # Generates full README from blocks just like legos
    └── prompts
        └── 🧠 prompts.go              # Nothing here yet
```
