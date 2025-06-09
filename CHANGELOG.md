# Changelog

## [1.33.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.33.0...v1.33.1) (2025-06-09)


### Bug Fixes

* add allow_user_to_delete field to MobileDeviceapplicationSubsetGeneral struct ([#799](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/799)) ([1b9441e](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/1b9441eeb6f117349d5c9de441bad48565339fb5))

## [1.33.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.32.0...v1.33.0) (2025-06-06)


### Features

* Add functions for GetCloudDistributionPoint and GetCloudDistributionPointTestConnection ([#797](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/797)) ([eae7257](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/eae725721f24a066424a44aaccf0359f98531ff4))
* add reclaim functionality for volume purchasing locations ([#791](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/791)) ([7f39971](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/7f39971f4dadc51eb62ffbb9a7b056afbb7c9976))


### Bug Fixes

* add error handling for SyncJamfProtectPlans ([#790](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/790)) ([3840423](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/38404231fa72117866fd22ce218cf24dbcf4efc2))
* mobile device applications - various ([#789](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/789)) ([5791e63](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/5791e631d7293dcaff4c6072ed8758e6859e2647))

## [1.32.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.31.0...v1.32.0) (2025-05-19)


### Features

* add new 11.16 fields for mobile device prestages and webhooks ([#787](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/787)) ([4b66f79](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/4b66f792ff4a60c95cfc20055615d994057579d1))

## [1.31.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.30.1...v1.31.0) (2025-05-15)


### Features

* add support for v2/engage ([#785](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/785)) ([ce74dc3](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/ce74dc3c5c287031e5396986878ed7557f0a9d1d))

## [1.30.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.30.0...v1.30.1) (2025-05-14)


### Bug Fixes

* refactor Jamf Protect functions ([#783](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/783)) ([b9203f4](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/b9203f4d096c32264317c5b27b0a34743e4a9ee4))

## [1.30.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.29.1...v1.30.0) (2025-05-12)


### Features

* add SSO certificate support ([#779](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/779)) ([811ebe8](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/811ebe8cb7b629d1ad9e4ee67879f879d63d2fbe))
* add support for device-enrollments/public-key retrieval ([#776](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/776)) ([230f3fe](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/230f3fed74a748dece91965a1753ba26b56a59a0))
* add volume purchasing location management examples and update SDK methods ([#777](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/777)) ([408ed30](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/408ed303f7b5be1bb39226123163cb1556a2a421))
* fix/complete cloud-ldaps support ([#781](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/781)) ([06dbd6a](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/06dbd6a5621dd0bf74c1dfe05b2a198063f34734))
* update support for /v3/sso (sso-settings)  ([#780](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/780)) ([56e95d2](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/56e95d2066feda8adfbbc034d75696113241a289))


### Bug Fixes

* HTTP response handling in SyncJamfProtectPlans ([#782](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/782)) ([6db7da4](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/6db7da426442232661f97ee46e9a2efbffb64b65))

## [1.29.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.29.0...v1.29.1) (2025-05-02)


### Bug Fixes

* mobile_device_prestages updates ([087f326](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/087f3266f2c223150f4bbb8374fc4e386b619304))
* mobile_device_prestages updates ([1a2e5cd](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/1a2e5cdbf355c25506e61aa5ad86cee38a2deab8))

## [1.29.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.28.0...v1.29.0) (2025-04-29)


### Features

* add IntPtr function to return a pointer to an int value ([dc473b4](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/dc473b4283447804041c4c87ed7cfaacbbc48508))
* add membership scoping optimization field to LDAPServerSubsetMappingUserGroupMemberships ([0ca14bc](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/0ca14bc9b654791e82ff9bce13020e9e68302480))
* add membership scoping optimization field to LDAPServerSubsetMappingUserGroupMemberships ([c311192](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/c311192c8e41085e50f0b9b31c803a5c91cb7e2e))
* added jamf service status check to recipes ([e3a1ecf](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/e3a1ecfd54a17c2710c19cadbc481546eed5c79b))
* implement CRUD examples for mobile device prestages ([32626b6](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/32626b603f37b088d9195d8fc51253d3ea4f67c3))
* Update ResourceMobileDevicePrestage structure to Jamf Pro 11.15.1 ([4e63159](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/4e63159688445e18f90e17bd090846be3e4aee51))


### Bug Fixes

* Add missing fields to ldapservers and correct existing fields ([755dc3c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/755dc3cb4394f8f46ff0624bb5f59704af7ffb10))
* Add missing fields to ldapservers and existing field mappings/data type corrections ([b2991c0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/b2991c0aac695a3572a15f95022d6ae5431ce5e2))
* Pagination size not returning ([3f98119](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/3f98119b0af4d6e721ae3078e31bec245a3fa61a))

## [1.28.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.27.1...v1.28.0) (2025-04-17)


### Features

* new pagination logic ([90fe99c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/90fe99c54de569c62a0be0ecb903dc83005775c1))

## [2.0.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v2.0.0...v2.0.1) (2025-04-17)


### Bug Fixes

* release please to release type go ([af961b5](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/af961b5cb3de7c9df1ec53b1dcf60aecfbae00f9))

## [2.0.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.27.1...v2.0.0) (2025-04-17)


### ⚠ BREAKING CHANGES

* Pagination func filter passing refactored ([#760](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/760))

### Features

* Pagination func filter passing refactored ([#760](https://github.com/deploymenttheory/go-api-sdk-jamfpro/issues/760)) ([5199c4e](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/5199c4e81489c30b8da64849d29fec5919b80a9f))

## [1.27.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.27.0...v1.27.1) (2025-04-14)


### Bug Fixes

* Add FlushSoftwareUpdatePlans to jamfproapi_enrollment_settings.go ([bdc5592](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/bdc5592a359786175cd7544c447d9b251ef59ab5))

## [1.27.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.26.4...v1.27.0) (2025-04-08)


### Features

* added notifications and command flushes with new recipes ([840e8ad](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/840e8adf185ed4cd4bc514a32ab7bcb7a27402de))
* added notifications and command flushes with new recipes ([8123cac](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/8123cac7b6126a57f2686dea9b14da481ed875bc))


### Bug Fixes

* added GetCloudIdentityProviderConfigurationByName ([4796644](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/4796644ce84a463bb5ada1ae7830dd5dc8e92f39))
* added GetCloudIdentityProviderConfigurationByName ([2227aca](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/2227aca14e21a760a986f7c9f936cb24a15f926b))

## [1.26.4](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.26.3...v1.26.4) (2025-04-04)


### Bug Fixes

* for CreateAccountDrivenUserEnrollmentAccessGroup URL construction ([9bd9bfa](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/9bd9bfa76240e3bc5134441525b0df7aa938de04))
* url construction for CreateAccountDrivenUserEnrollmentAccessGroup ([13c9fd4](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/13c9fd461de97ad418a247b0a7a760251966227e))

## [1.26.3](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.26.2...v1.26.3) (2025-04-04)


### Bug Fixes

* added missing func GetEnrollmentMessages to return a list ([5b197b2](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/5b197b21738a97490aac368282f572a207c13e0b))
* added missing func GetEnrollmentMessages to return a list ([75618b9](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/75618b9f89902133b585b85b2e264180fdf45efb))

## [1.26.2](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.26.1...v1.26.2) (2025-04-04)


### Bug Fixes

* added language code validation step to get/update/delete enrollment message by language id functions ([da5b3c9](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/da5b3c9fd27e5d07bb56af6b9f01f5ae8879b2db))
* added language code validation step to get/update/delete enrollment message by language id functions ([e9d9f7c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/e9d9f7cad1550bf52d8732f99932d8d6d481f36d))
* for returning list ([a5aa08c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/a5aa08c5cafd651899409344f071747cdfb03607))

## [1.26.1](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.26.0...v1.26.1) (2025-03-25)


### Bug Fixes

* added update support for global jamf pro app catalog settings ([8dfa37c](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/8dfa37c7f5eb03760a82d1471062d8daf6b72429))

## [1.26.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.25.2...v1.26.0) (2025-03-14)


### Features

* removed mandatory package timeout, made it adjustable via funcs ([5cb0eb3](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/5cb0eb38080672ff912a93c2467196a0a22ea676))

## [1.25.0](https://github.com/deploymenttheory/go-api-sdk-jamfpro/compare/v1.24.0...v1.25.0) (2025-02-21)


### Features

* added extra examples for updates ([1f10854](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/1f10854e2f50f895d631eaefff8216ae867ccd20))
* added undocumented api endpoints for enrollment customization prestage panes with examples ([3952b15](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/3952b15d0b4b3bf5d65605e795b753e73f1cff71))
* numerous fixes and new examples for Enrollment Customization ([8fc84b7](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/8fc84b72a382b0940ce7b03932e95c4b98777f4e))


### Bug Fixes

* removed redundant code ([fe8fc7e](https://github.com/deploymenttheory/go-api-sdk-jamfpro/commit/fe8fc7e417391e154f200717f8dce2120b0e4f9e))

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
