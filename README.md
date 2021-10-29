# ggen

Project to provide a alternative for [Cookiecutter](https://github.com/cookiecutter/cookiecutter) for define project generation.

## Install

```bash
curl -LO https://github.com/sergsoares/ggen/releases/download/v0.3.0/ggen
```

## Usage

Create a new folder and use:
```
./ggen -init

# It will create ggen.yml and a template folder with simple example.
```

Then you can use Go templates [Tutorial](https://blog.gopheracademy.com/advent-2017/using-go-templates/) for create scaffolds for your project.

---

## Roadmap next steps

- Simple cli with go templates for create one file.
- Add ggen init to generate basic configuration (air inspired)
- Implement safe manipulation of paths.
- Add parameters or config file.
- Add tests
- Pre-Build
- Strict mode to avoid "<no value>" appear inside templates
- Validate output based on extension (JSON/YAML/HCL)
- Verify if folder is empty in init
- Open templates from URL