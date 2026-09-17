[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](README_ru.md)

# Go bindings for the CS2 Custom Hud API Plugify plugin

This repository contains Go bindings for the [customhud](https://github.com/fr0nch/customhud) plugin. The bindings are automatically generated and kept in sync with the original plugin.

## API Documentation

Full API documentation can be found here on the [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/fr0nch/customhud/refs/heads/main/customhud.pplugin) website.

## Installation

```bash
go get github.com/fr0nch/go-plugify-customhud
```

## Updates

This repository uses GitHub Actions to automatically check for updates in [customhud](https://github.com/fr0nch/customhud) every day at 2:00 UTC. When a new version is detected, the bindings are updated, and a new release is created.

## License

This Go module for Plugify is licensed under the [MIT License](LICENSE).
