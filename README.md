# letstry templates

> THIS PROJECT IS UNDER ACTIVE DEVELOPMENT.

This repository contains publicly available templates for use with [LetsTry](https://github.com/letstrygo/letstry).

# Adding your template

To add a template to this repository:

1. Fork this repository.
2. Clone the your fork of the repository.
3. Add your template using the `add` command.
   ```powershell
   cd dist
   ./templates-{platform} add <name> <author> <author_url> <clone_url> <description>
   ```
4. Commit the updated sqlite database file.
5. Open a pull request with the updated database file.