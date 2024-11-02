GedeBox is an open source project to make CLI easier to use including document management tools. Inspired by busybox which uses 1 program with many functions.

This project was created for the development of the project in the [dotfiles repo](https://github.com/dotcreep/dotfiles.git) to be used on Linux, Windows and MacOS.

The goal is to make it easier for users to have their own tools in their local system to do things like document conversion, size compression, network tools like ssh access but with account management locally or in the cloud with E2EE.

# Shortcut Example
**Package Manager**

 Shortcut | Function
 --- | ---
`install` | Install Package
`installnc` | Install Package without confirm
`update` | Update Repository
`upgrade` | Upgrade Packages
`updateupgrade` | Update Repo and Upgrade Package
`remove` | Uninstall Package
`reinstall` | Reinstall package
`detail` | Detail Package
`orphan` | Remove Non-used Packages
`checkpackage` | Check Package
`listpackage` | List Package
`holdpackage` | Hold Package

**Specially AUR**

Shortcut | Function
--- | ---
`auri` | AUR Install
`auru` | AUR Update
`auruu` | AUR Upgrade
`aurr` | AUR Remove


# Tracking Feature
- [x] Package Manager
- [ ] Document Tools
- [ ] System Init / Loader on boot shortcut

# Target Support

System | Support
--- | ---
Linux | ✔️
Mac | ✔️
Windows (PowerShell) | ✔️