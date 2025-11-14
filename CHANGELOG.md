# Changelog

## 0.4.0-alpha.1 (2025-11-14)

Full Changelog: [v0.1.0-alpha.1...v0.4.0-alpha.1](https://github.com/llamastack/llama-stack-client-go/compare/v0.1.0-alpha.1...v0.4.0-alpha.1)

### ⚠ BREAKING CHANGES

* **api:** /v1/inspect only lists v1 apis by default
* **api:** /v1/inspect only lists v1 apis by default
* **api:** use input_schema instead of parameters for tools
* **api:** fixes to remove deprecated inference resources

### Features

* add new API filter for all non-deprecated APIs ([d74e63b](https://github.com/llamastack/llama-stack-client-go/commit/d74e63bc707992772addee329a21676f3f4aaabc))
* Adding option to return embeddings and metadata from `/vector_stores/*/files/*/content` and UI updates ([492a412](https://github.com/llamastack/llama-stack-client-go/commit/492a412c44448a6d76a2e8a04a5bd85c92deb5d0))
* **api:** Adding prompts API to stainless config ([9d436bb](https://github.com/llamastack/llama-stack-client-go/commit/9d436bb1701c851bbedc124471c8caefddd38334))
* **api:** expires_after changes for /files ([222bb4e](https://github.com/llamastack/llama-stack-client-go/commit/222bb4ed27e9a4afd247aacea50d1645326a6b54))
* **api:** fix completion response breakage perhaps? ([3e9c39f](https://github.com/llamastack/llama-stack-client-go/commit/3e9c39f1bf837daf2487ec2c465a9522e6b5befe))
* **api:** fix file batches SDK to list_files ([c9da417](https://github.com/llamastack/llama-stack-client-go/commit/c9da41734d782e758971ca2f5b9821669dbd7331))
* **api:** fixes to remove deprecated inference resources ([9f926b2](https://github.com/llamastack/llama-stack-client-go/commit/9f926b25335010398c825844a60927709ccfc07f))
* **api:** fixes to URLs ([6c9752f](https://github.com/llamastack/llama-stack-client-go/commit/6c9752f49b98e57c4ce06c6310f33d19207b3ae1))
* **api:** manual updates ([114833d](https://github.com/llamastack/llama-stack-client-go/commit/114833d8ba92478231c41c85c11ac25b16a39c96))
* **api:** manual updates ([1573312](https://github.com/llamastack/llama-stack-client-go/commit/15733129bb01716b4a1adcddc68f31c634e07f80))
* **api:** manual updates ([5444c7c](https://github.com/llamastack/llama-stack-client-go/commit/5444c7cf065438ce85e0648a3a21b1642421b97f))
* **api:** manual updates ([78a1b1a](https://github.com/llamastack/llama-stack-client-go/commit/78a1b1ad7d5c7c254b9c88f72316ba684077033a))
* **api:** manual updates ([eb137af](https://github.com/llamastack/llama-stack-client-go/commit/eb137afa61bf5621159b871561e3a11ca3589b00))
* **api:** manual updates ([9de5708](https://github.com/llamastack/llama-stack-client-go/commit/9de5708ab28e06c58731c3ac57c3e4c4fbc12aca))
* **api:** manual updates??! ([9dd7521](https://github.com/llamastack/llama-stack-client-go/commit/9dd7521061ccd921680da4d179375ae19ece3591))
* **api:** move datasets to beta, vector_db -&gt; vector_store ([df1a9e5](https://github.com/llamastack/llama-stack-client-go/commit/df1a9e5dd98267cffab0d79cc9c3c4124c882040))
* **api:** move post_training and eval under alpha namespace ([42bdca7](https://github.com/llamastack/llama-stack-client-go/commit/42bdca705b93c1c8c3330a4075bd8f5bde977f65))
* **api:** moving { rerank, agents } to `client.alpha.` ([f12fecf](https://github.com/llamastack/llama-stack-client-go/commit/f12fecf6d892348ebee06652255f6ac66f38e0bb))
* **api:** point models.list() to /v1/openai/v1/models ([010542c](https://github.com/llamastack/llama-stack-client-go/commit/010542c2ce776abdcb2ecbaa93cb4b754b10bfd8))
* **api:** query_metrics, batches, changes ([06e03be](https://github.com/llamastack/llama-stack-client-go/commit/06e03be805d93736fcf4f848c5f9888e2871c911))
* **api:** remove agents types ([56acfc3](https://github.com/llamastack/llama-stack-client-go/commit/56acfc3fe739820d08f74c109bda07e799c43724))
* **api:** remove openai/v1 endpoints ([754e8a4](https://github.com/llamastack/llama-stack-client-go/commit/754e8a415667396f5f84e84ab014dbc7a42a1e8a))
* **api:** removing openai/v1 ([b8635d7](https://github.com/llamastack/llama-stack-client-go/commit/b8635d7781c593fc1fb4bda7311189428f5bc128))
* **api:** SDKs for vector store file batches ([2060878](https://github.com/llamastack/llama-stack-client-go/commit/2060878c2b6b81c76b65f56dab6d699df12fb7d0))
* **api:** SDKs for vector store file batches apis ([e5f679f](https://github.com/llamastack/llama-stack-client-go/commit/e5f679f8f193fcdf76aba82288130854e9b86819))
* **api:** several updates including Conversations, Responses changes, etc. ([978c65f](https://github.com/llamastack/llama-stack-client-go/commit/978c65fe02f0e044843b278ac27e0d9cc38f1739))
* **api:** some updates to query metrics ([8243d43](https://github.com/llamastack/llama-stack-client-go/commit/8243d43dfa43bb9ec92d1edbed36f114240a265e))
* **api:** sync ([d3f850b](https://github.com/llamastack/llama-stack-client-go/commit/d3f850bfd124b84d4d49e384961fd99c3d90be02))
* **api:** tool api (input_schema, etc.) changes ([837277d](https://github.com/llamastack/llama-stack-client-go/commit/837277d4e6b4ffb57ce0136072b80400011da79a))
* **api:** updates to vector_store, etc. ([837fd3a](https://github.com/llamastack/llama-stack-client-go/commit/837fd3a7e19b103f5ea122e02fe5005afa5e8f62))
* **api:** updating post /v1/files to have correct multipart/form-data ([a3d6051](https://github.com/llamastack/llama-stack-client-go/commit/a3d6051547ce2b9cbd5af96b6b802515b737e7fb))
* **api:** use input_schema instead of parameters for tools ([a16eaef](https://github.com/llamastack/llama-stack-client-go/commit/a16eaef870f6ec94ae6adf36eed0d65bfa9fd3b8))
* **api:** vector_db_id -&gt; vector_store_id ([6c8dbbd](https://github.com/llamastack/llama-stack-client-go/commit/6c8dbbd143f4fb8b7631a6597b40110612cee87b))
* Implement the 'max_tool_calls' parameter for the Responses API ([ba81b90](https://github.com/llamastack/llama-stack-client-go/commit/ba81b90d0c4d49bdf100ce235353ba4c22212a40))


### Bug Fixes

* **api:** another fix to capture correct responses.create() params ([f3a9ee7](https://github.com/llamastack/llama-stack-client-go/commit/f3a9ee7303c890444802c76412d5d245a1420bdb))
* **api:** ensure openapi spec has deprecated routes ([fd39305](https://github.com/llamastack/llama-stack-client-go/commit/fd393057fb65ee6303391debfa09ab7578f9e1c1))
* **api:** fix the ToolDefParam updates ([65cef22](https://github.com/llamastack/llama-stack-client-go/commit/65cef2268480297f4233dd1c4c817aa03943f18e))
* bugfix for setting JSON keys with special characters ([ceb15f3](https://github.com/llamastack/llama-stack-client-go/commit/ceb15f300fdf9b7e1b2615c14c352878bcfc082b))
* **client:** fix circular dependencies and offset pagination ([0b95836](https://github.com/llamastack/llama-stack-client-go/commit/0b95836016ca0d089d3f7c07456ff5f55989011f))
* close body before retrying ([66adbea](https://github.com/llamastack/llama-stack-client-go/commit/66adbea266032b1198c76c8f590808d61a3d145a))
* fix stream event model reference ([d8b42f6](https://github.com/llamastack/llama-stack-client-go/commit/d8b42f67eefb216968989a10d68b2ff0e3e65a62))
* **internal:** unmarshal correctly when there are multiple discriminators ([d76c69c](https://github.com/llamastack/llama-stack-client-go/commit/d76c69c30d1402e13178448691d8202e6f2b5d82))
* MCP authorization parameter implementation ([e4d35c8](https://github.com/llamastack/llama-stack-client-go/commit/e4d35c89e989590b1df003d6b769b14d3d97b2b8))
* use slices.Concat instead of sometimes modifying r.Options ([15dfa47](https://github.com/llamastack/llama-stack-client-go/commit/15dfa47636cc1cd0ccb6b089ae363a7e70a5f56c))


### Chores

* **api:** /v1/inspect only lists v1 apis by default ([7ce9eb3](https://github.com/llamastack/llama-stack-client-go/commit/7ce9eb3a6695e8f6f6fdf8675f6199730728a348))
* **api:** /v1/inspect only lists v1 apis by default ([81ef91f](https://github.com/llamastack/llama-stack-client-go/commit/81ef91f4b156e5376f4857df41abd804d9a70f9f))
* bump minimum go version to 1.22 ([7b5d227](https://github.com/llamastack/llama-stack-client-go/commit/7b5d227df87389479dc2f6954ba59147b5d1a0fc))
* do not install brew dependencies in ./scripts/bootstrap by default ([062a46b](https://github.com/llamastack/llama-stack-client-go/commit/062a46b117baaf537ef9a0edef4222d7a1b3a839))
* fix readme example ([d639e08](https://github.com/llamastack/llama-stack-client-go/commit/d639e08566eb3986bae58061db0eca193fe4c407))
* fix readme examples ([7f95573](https://github.com/llamastack/llama-stack-client-go/commit/7f95573f7148f160c1054dd8656787f8c67ae603))
* **internal:** clean up ([7ba4b62](https://github.com/llamastack/llama-stack-client-go/commit/7ba4b6268aff1e3e1ab45a811e39b4acbe738121))
* **internal:** codegen related update ([30d522e](https://github.com/llamastack/llama-stack-client-go/commit/30d522ee9034d5ef35599aa5c6c6bcb2c512bfe3))
* **internal:** codegen related update ([a3cccf1](https://github.com/llamastack/llama-stack-client-go/commit/a3cccf10d30121514bb6b07a6416c589a1881763))
* **stainless:** add config for file header ([50d5a83](https://github.com/llamastack/llama-stack-client-go/commit/50d5a8365f3e0b68e313fd09566330c9a570c6db))
* update more docs for 1.22 ([67c0b00](https://github.com/llamastack/llama-stack-client-go/commit/67c0b0067523c93560b6d6467b81e3e8c2ecb61e))


### Documentation

* update examples ([245c643](https://github.com/llamastack/llama-stack-client-go/commit/245c643bb01b573243c31bea5f66761ef7e3fba1))

## 0.1.0-alpha.1 (2025-08-21)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/llamastack/llama-stack-client-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** manual updates ([85d523a](https://github.com/llamastack/llama-stack-client-go/commit/85d523a7a88beeccf8b9f0276a7269c53a16cfe2))
* **api:** update via SDK Studio ([3cf935f](https://github.com/llamastack/llama-stack-client-go/commit/3cf935feca91ef5cdb122ea6b927cb53517d09a6))
* **api:** update via SDK Studio ([d46e9ec](https://github.com/llamastack/llama-stack-client-go/commit/d46e9ec1644163b32129f9553074741320da1785))
* **api:** update via SDK Studio ([bb4e68b](https://github.com/llamastack/llama-stack-client-go/commit/bb4e68bf227b5071f5cd10102311f93d8695fb24))
* **api:** update via SDK Studio ([a6e0745](https://github.com/llamastack/llama-stack-client-go/commit/a6e0745e5ceea8ed6653a13b817d42d7e2eaa857))
* **api:** update via SDK Studio ([8a76c0a](https://github.com/llamastack/llama-stack-client-go/commit/8a76c0a4645a8794d2a12ff06d25e89448602a6f))
* **api:** update via SDK Studio ([f7bb530](https://github.com/llamastack/llama-stack-client-go/commit/f7bb5304af2a1cd0d90f7e309580dce74be0fc65))
* **api:** update via SDK Studio ([b4a019a](https://github.com/llamastack/llama-stack-client-go/commit/b4a019a51b93ea06fc1e0ef37a7b4062836763c7))
* **api:** update via SDK Studio ([bbeecc7](https://github.com/llamastack/llama-stack-client-go/commit/bbeecc71fde53f87959c7aff72061da75ac2d5bf))
* **api:** update via SDK Studio ([4903a50](https://github.com/llamastack/llama-stack-client-go/commit/4903a50fd3493eef613cd7647b7f020e66302260))
* **api:** update via SDK Studio ([eb6a7e7](https://github.com/llamastack/llama-stack-client-go/commit/eb6a7e756466b2167f524ed615cc023998108b8f))
* **api:** update via SDK Studio ([1b3b90a](https://github.com/llamastack/llama-stack-client-go/commit/1b3b90a5faa3205827a8afcfef16744f54fc2e25))
* **api:** update via SDK Studio ([4bf99e2](https://github.com/llamastack/llama-stack-client-go/commit/4bf99e20b4f2f217aead81c10e998a8efa2bda63))
* **api:** update via SDK Studio ([80aac92](https://github.com/llamastack/llama-stack-client-go/commit/80aac927dcbe6a6f432729ca118ca0adb3256871))
* **api:** update via SDK Studio ([491612b](https://github.com/llamastack/llama-stack-client-go/commit/491612b3b9faac4bed7e0c8fc3704400a9bae72b))
* **api:** update via SDK Studio ([1aabd16](https://github.com/llamastack/llama-stack-client-go/commit/1aabd16a7a623b23c76bdd0e170c9cbe1bade3d1))
* **api:** update via SDK Studio ([4d82937](https://github.com/llamastack/llama-stack-client-go/commit/4d82937d18f2f2edc5712856c4392ac1f7577f64))
* **api:** update via SDK Studio ([9551913](https://github.com/llamastack/llama-stack-client-go/commit/9551913aeae9e18c28c39e30d4a258ee927550b3))
* **api:** update via SDK Studio ([309e275](https://github.com/llamastack/llama-stack-client-go/commit/309e2756b27dd666c6b46599340936fea8f8db09))
* **api:** update via SDK Studio ([dcc41ff](https://github.com/llamastack/llama-stack-client-go/commit/dcc41ffcbdec14a476985db477ce85e205e598ba))
* **client:** support optional json html escaping ([12fa3b8](https://github.com/llamastack/llama-stack-client-go/commit/12fa3b8b833d3c0476e7b9d050c9d788776e96ca))


### Bug Fixes

* **client:** process custom base url ahead of time ([34759ab](https://github.com/llamastack/llama-stack-client-go/commit/34759ab215115e3603963256d9b9526a260f44d4))
* Fix LICENSE to MIT ([#5](https://github.com/llamastack/llama-stack-client-go/issues/5)) ([d37397c](https://github.com/llamastack/llama-stack-client-go/commit/d37397ce499dddeb925b3469296efe8ae578b5ff))


### Chores

* configure new SDK language ([dc6ea4a](https://github.com/llamastack/llama-stack-client-go/commit/dc6ea4a0e9dcf92a9615bb736e2665e332f34e04))
* **internal:** update comment in script ([9d93e2e](https://github.com/llamastack/llama-stack-client-go/commit/9d93e2e728123bc5b702e6d356fe2a29a0907fee))
* update @stainless-api/prism-cli to v5.15.0 ([f61fee4](https://github.com/llamastack/llama-stack-client-go/commit/f61fee4f6499288c3b907b7e9fc05c9c52df8eeb))
* update SDK settings ([95d2d64](https://github.com/llamastack/llama-stack-client-go/commit/95d2d649fd287ab4c4bf325bb9f325e3e062c4d3))
