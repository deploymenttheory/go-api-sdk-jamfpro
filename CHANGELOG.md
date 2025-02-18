# Changelog

## [1.24.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.23.2...v1.24.0) (2025-02-18)


### Features

* added locales to the SDK ([9b7d3b5](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/9b7d3b5937df48363a5816ddabc539bd2d402cf0))
* added time zones to sdk ([986d85f](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/986d85f900ba751c22f352c86151f0c117468dbb))

## [1.23.2](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.23.1...v1.23.2) (2025-02-18)


### Bug Fixes

* changed account-preferences to correct uri for PATCH method ([f64a2a2](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/f64a2a2fc5962134985b9689dc5fcec5d09be801))

## [1.23.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.23.0...v1.23.1) (2025-02-17)


### Bug Fixes

* Updated SMTP server settings examples ([104fc6c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/104fc6c1064b8b2b37ad92c4386fac97e9a9cc59))

## [1.23.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.22.1...v1.23.0) (2025-02-17)


### Features

* added examples for device communication settings ([2b09dcd](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/2b09dcddd439d1138b7a460bf3dd8af05df2024e))
* updated smtp server to use v2 endpoint ([735a7d4](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/735a7d4b33d0504480dd863cc13fe0d0bb3d725d))

## [1.22.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.21.0...v1.22.0) (2025-02-17)


### Features

* added policy properties to the sdk with examples ([9e0fceb](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/9e0fceb8860a3bbc31b0d28906538221b110c17a))


### Bug Fixes

* added UseFido2 field to self service settings ([981b31c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/981b31ceffcc2c1150701c787fa5e12806ba2f4f))

## [1.21.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.20.0...v1.21.0) (2025-01-22)


### Features

* add examples for device enrollment management in Jamf Pro SDK ([b32d204](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/b32d2046c4f2f56714c3bdb6bf3d62e02da7e292))


### Bug Fixes

* for DoPackageUpload utility + recipe for patch policies ([4dc6336](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/4dc6336036ecdad7b5d14809781735c1e4fa44e2))

## [1.20.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.19.0...v1.20.0) (2025-01-21)


### Features

* add examples for managing patch policies in Jamf Pro SDK using jamf pro api. deprecated old endpoint ([4b01255](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/4b012551662223fe9d0245b4b602533a67ad0251))
* add functions to retrieve and update Jamf Connect config profiles by UUID, ID, and name ([7663601](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/76636011676093ef0a3c62e8a90699147701d686))
* refactored patch policies, patch_software_title_configurations and updated examples ([947a087](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/947a08718055bfb51584e5b2d64dfb66b9aa4600))

## [1.19.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.18.1...v1.19.0) (2025-01-20)


### Features

* add AcceptPatchManagementDisclaimer example for Jamf Pro client ([a9785e5](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/a9785e57b2fed29795409f934cdc428b7c43640e))
* add examples for managing external patch sources in Jamf Pro SDK ([4c06f01](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/4c06f01db9eadd02803ea029f509138b7f5dad3c))

## [1.18.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.17.1...v1.18.0) (2025-01-06)


### Features

* add examples for log flushing tasks and settings in Jamf Pro API ([132dc2c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/132dc2ce7405ba220fd2e0e7a51d90a4817d089d))

## [1.17.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.16.1...v1.17.0) (2024-12-20)


### Features

* add GetJamfAppCatalogAppInstallerByName function and example usage ([92fea22](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/92fea22d0323c7f1aaf622c07b77df25f1630de3))

## [1.16.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.16.0...v1.16.1) (2024-12-20)


### Features

* adding outstanding GET by name funcs with examples ([6dc9bf1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/6dc9bf1a38c12a3acd6b69429e11d23185723863))
* implement Jamf Pro system initialization and startup status monitoring ([1d54302](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/1d54302b079c1c40f057538a79e8ac723a277168))
* update macOS configuration profile examples and add export functionality ([bca68b5](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/bca68b5a8af11d7c82bcfad0cd47c961f0c15f2a))


### Bug Fixes

* update profile name comment for clarity in macOS configuration example ([113197f](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/113197f7c03ff44fcc9f2574a89ae9d66f1d4ffc))
