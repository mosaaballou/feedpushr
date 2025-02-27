<a name=""></a>
# [](https://github.com/ncarlier/feedpushr/compare/v2.0.0-rc.1...v) (2019-09-01)



<a name="2.0.0-rc.1"></a>
# [2.0.0-rc.1](https://github.com/ncarlier/feedpushr/compare/v1.2.0...v2.0.0-rc.1) (2019-09-01)


### Features

* **api:** WIP add CRUD API for filters and outputs ([fbbbd78](https://github.com/ncarlier/feedpushr/commit/fbbbd78))
* **api:** WIP add CRUD API for filters and outputs ([a421877](https://github.com/ncarlier/feedpushr/commit/a421877))
* **api:** WIP add Spec API for filters and outputs ([7c00d24](https://github.com/ncarlier/feedpushr/commit/7c00d24))
* **store:** add filter repository ([f28665e](https://github.com/ncarlier/feedpushr/commit/f28665e))
* **ui:** add about page ([ce341c8](https://github.com/ncarlier/feedpushr/commit/ce341c8))
* **ui:** add output pages ([f0e1e63](https://github.com/ncarlier/feedpushr/commit/f0e1e63))
* **ui:** configure main theme ([464f049](https://github.com/ncarlier/feedpushr/commit/464f049))
* **ui:** new filter pages ([85eb0f6](https://github.com/ncarlier/feedpushr/commit/85eb0f6))
* **ui:** switch to new UI ([495e410](https://github.com/ncarlier/feedpushr/commit/495e410))
* **ui:** WIP filter pages ([209add6](https://github.com/ncarlier/feedpushr/commit/209add6))
* auto create all buckets ([779e897](https://github.com/ncarlier/feedpushr/commit/779e897))
* new UI foundation ([7307580](https://github.com/ncarlier/feedpushr/commit/7307580))
* persist filter and output configuration ([c7bc4b3](https://github.com/ncarlier/feedpushr/commit/c7bc4b3))
* persist output configuration ([2ff1d20](https://github.com/ncarlier/feedpushr/commit/2ff1d20))



<a name="1.2.0"></a>
# [1.2.0](https://github.com/ncarlier/feedpushr/compare/v1.1.0...v1.2.0) (2019-06-12)


### Bug Fixes

* **aggregator:** reset delay when manual start ([0d66cc9](https://github.com/ncarlier/feedpushr/commit/0d66cc9)), closes [#6](https://github.com/ncarlier/feedpushr/issues/6)
* disable agent for ARM architecture ([4faea98](https://github.com/ncarlier/feedpushr/commit/4faea98))


### Features

* move agent from base code to contrib ([d4c55c3](https://github.com/ncarlier/feedpushr/commit/d4c55c3))
* plugins autoload ([4af51b7](https://github.com/ncarlier/feedpushr/commit/4af51b7))
* use systray for desktop environment ([3826b84](https://github.com/ncarlier/feedpushr/commit/3826b84))



<a name="1.1.0"></a>
# [1.1.0](https://github.com/ncarlier/feedpushr/compare/v1.0.3...v1.1.0) (2019-05-14)


### Bug Fixes

* **opml:** fix import with inline categories ([1aebfef](https://github.com/ncarlier/feedpushr/commit/1aebfef))


### Features

* **aggregator:** store aggregation status ([b1d00db](https://github.com/ncarlier/feedpushr/commit/b1d00db)), closes [#5](https://github.com/ncarlier/feedpushr/issues/5)
* **filter:** use readflow readability function for fetch filter ([7af532d](https://github.com/ncarlier/feedpushr/commit/7af532d))
* **tags:** add negative tag ([1692e5d](https://github.com/ncarlier/feedpushr/commit/1692e5d))
* **ui:** add form micro help ([00b6487](https://github.com/ncarlier/feedpushr/commit/00b6487)), closes [#3](https://github.com/ncarlier/feedpushr/issues/3)
* **ui:** loading screen on OPNML imports ([315d32a](https://github.com/ncarlier/feedpushr/commit/315d32a))



<a name="1.0.3"></a>
## [1.0.3](https://github.com/ncarlier/feedpushr/compare/v1.0.2...v1.0.3) (2019-05-07)


### Bug Fixes

* **aggregator:** limit max delay between checks to h24 ([e0a64d1](https://github.com/ncarlier/feedpushr/commit/e0a64d1))


### Features

* **opml:** add category support ([060effe](https://github.com/ncarlier/feedpushr/commit/060effe)), closes [#2](https://github.com/ncarlier/feedpushr/issues/2) [#4](https://github.com/ncarlier/feedpushr/issues/4)



<a name="1.0.2"></a>
## [1.0.2](https://github.com/ncarlier/feedpushr/compare/v1.0.1...v1.0.2) (2019-04-25)


### Bug Fixes

* fix articles counter ([28850d1](https://github.com/ncarlier/feedpushr/commit/28850d1))
* **aggregator:** reduce error logs level ([9ddce0d](https://github.com/ncarlier/feedpushr/commit/9ddce0d))


### Features

* add maskSecret function for plugins ([e6440f6](https://github.com/ncarlier/feedpushr/commit/e6440f6))
* export configuration variables ([65217fc](https://github.com/ncarlier/feedpushr/commit/65217fc))



<a name="1.0.1"></a>
## [1.0.1](https://github.com/ncarlier/feedpushr/compare/v1.0.0...v1.0.1) (2019-04-10)


### Bug Fixes

* fix tags usage for filters and outputs ([78b92a4](https://github.com/ncarlier/feedpushr/commit/78b92a4))



<a name="1.0.0"></a>
# [1.0.0](https://github.com/ncarlier/feedpushr/compare/4a64395...v1.0.0) (2019-04-09)


### Bug Fixes

* **aggregator:** fix bad cache retention setup ([fe603bc](https://github.com/ncarlier/feedpushr/commit/fe603bc))
* **aggregator:** fix unit test ([9936ac0](https://github.com/ncarlier/feedpushr/commit/9936ac0))
* **builder:** fix nil pointer ([eb74c80](https://github.com/ncarlier/feedpushr/commit/eb74c80))
* **docker:** fix plugins copy ([047693a](https://github.com/ncarlier/feedpushr/commit/047693a))
* **docker:** remove plugins from image ([d859025](https://github.com/ncarlier/feedpushr/commit/d859025))
* **filter:** minify filter should also filtering description ([087c4be](https://github.com/ncarlier/feedpushr/commit/087c4be))
* **output:** fix error messages ([04787d3](https://github.com/ncarlier/feedpushr/commit/04787d3))
* copy article link ([44c7da5](https://github.com/ncarlier/feedpushr/commit/44c7da5))
* **plugin:** init output and filter plugin registry ([692c934](https://github.com/ncarlier/feedpushr/commit/692c934))
* **pshb:** add support for (broken) Wordpress plugin ([e7334b3](https://github.com/ncarlier/feedpushr/commit/e7334b3))
* fix bad error format string ([74ab0a8](https://github.com/ncarlier/feedpushr/commit/74ab0a8))
* **store:** fix db uri ([aa58032](https://github.com/ncarlier/feedpushr/commit/aa58032))
* **ui:** show error messages and pshb status ([fa366e7](https://github.com/ncarlier/feedpushr/commit/fa366e7))


### Features

* configure CORS ([a6e53f1](https://github.com/ncarlier/feedpushr/commit/a6e53f1))
* improve exploitation logs ([8be34ae](https://github.com/ncarlier/feedpushr/commit/8be34ae))
* improve filter/output descriptions ([444cea7](https://github.com/ncarlier/feedpushr/commit/444cea7))
* **tags:** add tags support ([a59998b](https://github.com/ncarlier/feedpushr/commit/a59998b))
* multiple outputs support ([2052116](https://github.com/ncarlier/feedpushr/commit/2052116))
* **ui:** add Web user interface ([aa8d50e](https://github.com/ncarlier/feedpushr/commit/aa8d50e))
* overwriting of the environmental configuration via parameters ([4a64395](https://github.com/ncarlier/feedpushr/commit/4a64395))
* **aggregator:** add timeout configuration ([45a887a](https://github.com/ncarlier/feedpushr/commit/45a887a))
* **api:** add feed title managment ([f837590](https://github.com/ncarlier/feedpushr/commit/f837590))
* **api:** remove feed list limit ([a76d4c2](https://github.com/ncarlier/feedpushr/commit/a76d4c2))
* **docker:** create an image with plugins ([20b9e49](https://github.com/ncarlier/feedpushr/commit/20b9e49))
* **feed:** add processed items counter ([a5d3645](https://github.com/ncarlier/feedpushr/commit/a5d3645))
* **feed:** add status attribute ([9b82c9b](https://github.com/ncarlier/feedpushr/commit/9b82c9b))
* **filter:** add counter to title filter props ([69d2c42](https://github.com/ncarlier/feedpushr/commit/69d2c42))
* **filter:** add fetch filter ([0e72749](https://github.com/ncarlier/feedpushr/commit/0e72749))
* **filter:** add filter system (with plugins) ([8d6fd44](https://github.com/ncarlier/feedpushr/commit/8d6fd44))
* **filter:** add minify filter ([d37d104](https://github.com/ncarlier/feedpushr/commit/d37d104))
* **filter:** use last readability lib for fetch filter ([0530fbc](https://github.com/ncarlier/feedpushr/commit/0530fbc))
* **logging:** add Sentry for error recording ([f674543](https://github.com/ncarlier/feedpushr/commit/f674543))
* **logging:** use async Sentry call ([55715e0](https://github.com/ncarlier/feedpushr/commit/55715e0))
* **opml:** don't import existing feeds ([a804718](https://github.com/ncarlier/feedpushr/commit/a804718))
* **output:** add external plugin support ([f2d4656](https://github.com/ncarlier/feedpushr/commit/f2d4656))
* **plugin:** refactoring of the plugin system ([12c009f](https://github.com/ncarlier/feedpushr/commit/12c009f))
* **pshb:** add max TTL for a PSHB subscription ([cf6b34b](https://github.com/ncarlier/feedpushr/commit/cf6b34b))
* **pshb:** compute subscribtion details URL ([83c1aae](https://github.com/ncarlier/feedpushr/commit/83c1aae))
* **ui:** add feed filter bar ([03314bc](https://github.com/ncarlier/feedpushr/commit/03314bc))
* **ui:** display filter and output descriptions as details ([9046f0a](https://github.com/ncarlier/feedpushr/commit/9046f0a))
* **ui:** make feed table sortable ([cc319fc](https://github.com/ncarlier/feedpushr/commit/cc319fc))
* **ui:** make Output a functional component ([f43d4c5](https://github.com/ncarlier/feedpushr/commit/f43d4c5))
* use URL to declare filters and outputs ([a77d74d](https://github.com/ncarlier/feedpushr/commit/a77d74d))



