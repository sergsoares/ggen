# ggen

Project to provide a alternative for [Cookiecutter](https://github.com/cookiecutter/cookiecutter) for define project generation.

### Usage
```
cat << EOF > g.yml
template_path: "examples"
output_path: "dist"

data:
  internal_port: 3000 
  name: ggen
  Timeout: 2000
EOF
```

---

## Roadmap next steps

x Simple cli with go templates for create one file.
- Add ggen init to generate basic configuration (air inspired)
- Implement safe manipulation of paths.
- Add tests
- Pre-Build