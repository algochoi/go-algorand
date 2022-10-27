window.BENCHMARK_DATA = {
  "lastUpdate": 1666882088183,
  "repoUrl": "https://github.com/algochoi/go-algorand",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "86622919+algochoi@users.noreply.github.com",
            "name": "algochoi",
            "username": "algochoi"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a8b32bc42c669ba31e26abbf1a374f2261b73f15",
          "message": "Bench extract (#21)\n\n* Update the Version, BuildNumber, genesistimestamp.data\r\n\r\n* testing: fix TestRekeyUpgrade when round=0 (#2188)\r\n\r\nThe previous fix for this test (#2178) was incomplete. It did not included proper handling for the case of round 0.\r\nThis PR addresses this corner case.\r\n\r\n* Improve fresh node startup time (#2185)\r\n\r\n## Summary\r\n\r\nThis PR contains three different changes that would allow the node to start up faster:\r\n1. It disables the block 0 fix we previously deployed. This fix was intended to address catching up archival nodes, where we would calculate the block 0 hash incorrectly. The fix was long deployed, and the hash calculation is also fixed since.\r\n2. The tracker and blocks databases are now being opened concurrently.\r\n3. The accounts tracker database schema upgrade would now skip some of the steps that aren't required for a fresh database.\r\n\r\nCombining all the above, the ledger startup of a new node is now about 2-3 times faster.\r\n\r\n## Test Plan\r\n\r\nThis change has no functional enduser change. The existing tests provide sufficient coverage.\r\n\r\n* PeerSelector and Catchup service tests (#2184)\r\n\r\nUnit tests to improve the test code coverage in peerSelector and catchup/service\r\npeerSelector.go 100% coverage\r\ncatchup/service.go ~69%\r\n\r\n* Adding missing test for universalFetcher.go (#2175)\r\n\r\nAdding missing test for universalFetcher.go\r\n\r\nCode coverage increased from 74%->96%\r\n\r\n* testing: fix data race accessing messagesOfInterest during network shutdown (#2181)\r\n\r\n<!--\r\nThanks for submitting a pull request! We appreciate the time and effort you spent to get this far.\r\n\r\nIf you haven't already, please make sure that you've reviewed the CONTRIBUTING guide:\r\nhttps://github.com/algorand/go-algorand/blob/master/CONTRIBUTING.md#code-guidelines\r\n\r\nIn particular ensure that you've run the following:\r\n* make generate\r\n* make sanity (which runs make fmt, make lint, make fix and make vet)\r\n\r\nIt is also a good idea to run tests:\r\n* make test\r\n* make integration\r\n-->\r\n\r\n## Summary\r\n\r\nFix race [reported here](https://github.com/algorand/go-algorand-internal/issues/1268):\r\n```\r\nRead at 0x00c00675ed48 by goroutine 210:\r\n  github.com/algorand/go-algorand/network.(*WebsocketNetwork).ServeHTTP()\r\n      /home/travis/gopath/src/github.com/algorand/go-algorand/network/wsNetwork.go:1114 +0x19ad\r\n  github.com/gorilla/mux.(*Router).ServeHTTP()\r\n      /home/travis/gopath/pkg/mod/github.com/gorilla/mux@v1.6.2/mux.go:162 +0x193\r\n  github.com/algorand/go-algorand/network.(*RequestTracker).ServeHTTP()\r\n      /home/travis/gopath/src/github.com/algorand/go-algorand/network/requestTracker.go:474 +0x77b\r\n  net/http.serverHandler.ServeHTTP()\r\n      /home/travis/.gimme/versions/go1.14.7.linux.amd64/src/net/http/server.go:2836 +0xce\r\n  net/http.(*conn).serve()\r\n      /home/travis/.gimme/versions/go1.14.7.linux.amd64/src/net/http/server.go:1924 +0x837\r\nPrevious write at 0x00c00675ed48 by goroutine 79:\r\n  github.com/algorand/go-algorand/network.(*WebsocketNetwork).Stop()\r\n      /home/travis/gopath/src/github.com/algorand/go-algorand/network/wsNetwork.go:855 +0x265\r\n  github.com/algorand/go-algorand/network.TestGetPeers()\r\n      /home/travis/gopath/src/github.com/algorand/go-algorand/network/wsNetwork_test.go:977 +0xe1d\r\n  testing.tRunner()\r\n      /home/travis/.gimme/versions/go1.14.7.linux.amd64/src/testing/testing.go:1039 +0x1eb\r\n\r\n```\r\n\r\n## Test Plan\r\n\r\nThe CI should run this test. I don't think the flaky test runs on pull requests right now, but it should run during the nightly builds. \r\n\r\nThe race is not reproducible on my system. Running this before and after the test does not report any errors (until a 10 minute timeout):\r\n```\r\n$ gotestsum --format testname -- -tags \"sqlite_unlock_notify sqlite_omit_load_extension osusergo netgo static_build\" -race github.com/algorand/go-algorand/network -run TestGetPeers -count 1000000000\r\n```\r\n\r\n* Add mention of install_buildtools.sh to README (#2195)\r\n\r\nRecently the build tooling was updated in #2108 which split out the installation of Go tools golint, stringer, swagger, msgp into a separate script install_buildtools.sh. This updates the README to reflect that in the environment setup instructions.\r\n\r\n* Remove buildnumber.dat and genesistimestamp.dat files.\r\n\r\n* run some independent commands in parallel, total test 230s -> 170s (#2197)\r\n\r\nrun some independent commands in parallel, total test 230s -> 170s\r\n`20210526_115456 :179 finished e2e_subs/sectok-app.sh OK in 235.471021 seconds`\r\ndown to\r\n`20210526_122739 :179 finished e2e_subs/sectok-app.sh OK in 169.771514 seconds`\r\n\r\n* Update version.go\r\n\r\n* testing: e2e sub assets-app speedup (#2198)\r\n\r\nOptimize assets-app test\r\n\r\n* Clean up `goal wallet new` output (#2183)\r\n\r\nWhen you run goal wallet new there is a confusing warning (\"One or more non-printable characters ...\") presented just before the backup phrase, which is displayed between some sanitized ANSI codes. This is due to security improvements in #1585 to prevent goal from printing control characters. This PR removes the ANSI color formatting altogether from the infoBackupPhrase (already done for Windows in #1942), which is one simple way to clean up this output and remove the warning message.\r\n\r\n* testing: fix bug in TestApplicationsUpgradeOverREST (#2196)\r\n\r\nThe TestApplicationsUpgradeOverREST e2e had a bug in the testing code -\r\nif the upgrade takes place before the broadcast takes place, the broadcast result could be non-err.\r\nThis PR ensures that if we get a non-error, we have already upgraded.\r\n\r\n* debug tools: improve carpenter by adding timestamp (#2207)\r\n\r\ndebug tools: improve carpenter by adding timestamp\r\n\r\n* Improve testing predicate to provide better info in case of a fail case. (#2206)\r\n\r\nThe e2e test TestRewardUnitThreshold was failing on this statement:\r\n```golang\r\nr.Truef(latestBalanceNewAccount.PendingRewards >= (initialBalanceNewAccount+amountRichAccountPokesWith)/rewardUnit, \"new account should have pending rewards (e2e)\")\r\n```\r\n\r\nI replaces the `r.Truef` with `r.GreaterOrEqualf` to get the values off when it fails.\r\n( I wan't able to reproduce it, so I figured getting some extra info for the next time would be helpful )\r\n\r\n* Re-enable gofmt check during travis build (#2212)\r\n\r\nWhile talking to @tsachiherman we noticed stray fmt'd files in my repo after running make sanity, and he suggested submitting a fix PR. I discovered this shell line generated a list of gofiles that is always empty, probably because the way we set up builds for branches (and vendoring) has changed since it was written.\r\n\r\n* A test recommendended by Pavel, an imporvment rec'd by Brian (#2216)\r\n\r\nTwo small improvements recommended during CR of teal4-math\r\n\r\n* \"Rent\" larger program size (#2157)\r\n\r\nAllow paying for more program space.\r\n\r\n* Access values from a previous app call's scratch space (#2158)\r\n\r\nThis change allows contract to contract composability by allowing future app call transactions to read the scratch space of previous transactions. This allows smart contracts to expose side effects for future application calls, such as a price oracle returning the exchange rate of a particular asset pair.\r\n\r\nChanges include:\r\n\r\nAdding PastSideEffects fields to EvalParams.\r\ngload and gloads docs.\r\n\r\n* Refresh wallet handle after waiting for series of rounds, as handle might expire. (#2203)\r\n\r\nRefresh wallet handle after waiting for series of rounds, as handle might expire\r\n\r\n* Fix proposal propagation (#2079)\r\n\r\nWhen a relay receives a proposal for a future round, the proposal fails to be relayed. This change fixes that by sending it when the agreement advances to the corresponding round.\r\n\r\n* Correct path to genesis file in nightly perf testing (#2218)\r\n\r\nFix issue cp: cannot create regular file 'gen/devnet/genesis.json': No such file or directory by creating a directory if it doesn't exist to store the genesis file. Changed the reference genesis.json file to installer/genesis/ path.\r\n\r\n* pool fees in a txgroup (#2173)\r\n\r\nA fair amount of complexity is introduced in transaction groups when the goal is to let some entity perform an action at the expense of another. For example, a contract account might be willing to perform an exchange, but expects the caller to compensate it to replace the fee that the contract account must pay.\r\n\r\nThis changes fee accounting to simplify these situations. Rather than check that each txn in a txn group meets the min fee, the txgroup is checked as a whole to ensure the total fee exceeds n*min_fee for an n member txgroup.\r\n\r\n* Fixing a bug in eval and TestOnSwitchToUnSupportedProtocol (#2220)\r\n\r\nestOnSwitchToUnSupportedProtocol had multiple segments. This test is separated into multiple tests.\r\n\r\nThis PR fixes two issues:\r\n\r\nFix for a bug in eval: in the event of an error and early termination of the eval, the loading of accounts in a separate go routine was not terminated. loadAccounts will be terminated when eval returns.\r\nTestOnSwitchToUnSupportedProtocol1 had a bug in setting the blocks for protocol switch. The first block was not getting the NextProtocolSwitchOn set to 1. Despite the bug, the test was passing most of the time, and failing some of the time.\r\nTestOnSwitchToUnSupportedProtocol3 had a bug in setting the blocks for protocol switch. AddBlocks starts adding from the next round of the passed block. NextProtocolSwitchOn was getting set one round late.\r\n\r\n* Extend catchpoint wait for round timeout from 10 to 60 milliseconds. (#2231)\r\n\r\nExtend catchpoint test next round timeout from 10 to 60 seconds to try to avoid test failure.\r\n\r\n* bandwidth stats for cluster tests (#2159)\r\n\r\nheapWatch.py gains the ability to get a list of hosts from terraform-inventory from algonet tool. Can have --token and --admin-token set for talking to those cluster algod instances. (Cluster tool is separately updated to set one token across all cluster algod.)\r\n\r\nmetrics_delta.py processes captured /metrics data into stats for each algod followed and also produces one summary line of all data processed, e.g.:\r\nsummary: 1091.99 TPS, 26341397 tx B/s, 3490256 rx B/s\r\n\r\n* Check setbyte length properly. (#2236)\r\n\r\nWe we checking the against the length with > instead of >=.  That's\r\nwrong, so we could cause a panic instead of clean error when trying to\r\nset the byte 1 past the array length.  Fortunately, both have the same\r\neffect: the txn fails. But we should do this properly.  Thanks to\r\n@amityadav0 for the report in #2221.\r\n\r\n* Use minimal travis container since we install go ourselves. (#2232)\r\n\r\nSince we manage go ourselves I want to see if our build works without running the travis go setup step.\r\n\r\n* Modify build_release to match build_pr and remove ARM Deploy. (#2237)\r\n\r\nThis pulls in the cleanups we've made to other testing into the rel/\r\ntests. Additionally, the ARM Deploy builder has not successfully built\r\nin a while, so we'll remove it altogether.\r\n\r\n* Fix rel nightly test by extending catchpoint round timeout (#2240)\r\n\r\nFix catchpoint unit test by extending timeout from 10 to 60 seconds.\r\n\r\n* Fix the two docker files - avoid make deps step, which is no longer needed. (#2241)\r\n\r\nFix the two docker files - avoid make deps step, which is no longer needed.\r\nThe compilation of go-algorand no longer requires the installation ( and therefore the validation ) of the dependencies.\r\n\r\n* Use gotestsum instead of logfilter for test formatting. (#2161)\r\n\r\nUse `gotestsum` to format go test output, include duration messages in a few spots that take up some time, simplify some small things here and there.\r\n\r\n* Ensures that the catchup service have the network library load the DNS records before attempting the first sync. (#2248)\r\n\r\nEnsures that the catchup service have the network library load the DNS records before attempting the first sync -\r\n\r\nCurrent code was attempting to sync without any DNS records available, failing, and trying again later on. Since the DNS records were being refreshed every 60 seconds, the third or fourth attempt would be successful. This PR attempt to give the catchup service a good chance of succeeding on the first attempt.\r\n\r\n* Combine Clear and Approval Program size limits (#2225)\r\n\r\nApps have been able to use 1k space for each of their programs since apps were introduced in v24. But clear state programs are quite small. It is more useful to provide 2k of space, divided however the app prefers. This PR does that, including giving 2k extra space for each unit of \"extra pages\" requested at app creation time.\r\n\r\nTests considering before and after consensus updates.\r\n\r\n* Remove ci-deps and update GOPROXY in Dockerfiles (#2247)\r\n\r\nRecent refactoring has changed dependency installation. ci-deps did not appear to be used in a meaningful way, so it was removed from the Makefile and Dockerfiles. Additionally, downloading from go has become unstable, so updating the GOPROXY options should help.\r\n\r\n* codecov integration (#2228)\r\n\r\nIntegrate with codecov.io to attach coverage reports to PRs\r\n\r\n* Disable compact certs and auction tests (#2254)\r\n\r\nThey need work and shouldn't block our releases since they features aren't currently in use/enabled.\r\n\r\n* If the branch is rel/nightly set channel to nightly (#2252)\r\n\r\nThe way that the new pipeline scripts computed channel did not account for the rel/nightly branch, which should set the channel to nightly. This was preventing rpm and deb binaries from being built. This change identifies if the TRAVIS_BRANCH env is set to rel/nightly, and if so, will set the channel to nightly.\r\n\r\n* Optimize constant assembly (#2215)\r\n\r\nWith this PR, constants introduced by pseudo-ops (int, byte, addr) are now placed into constant blocks or loaded with pushint/pushbytes in the optimal way to save program space, starting with TEAL v4.\r\n\r\nThis optimization effectively reorders the intcblock and bytecblock that the assembler creates for pseudo-op constants such that the most frequently referenced constants are first (and can thereby take advantage of the space-saving intc_X/bytec_X opcodes). Additionally, any constants referenced only once are taken out of the constant block and instead loaded with pushint or pushbytes to further save space.\r\n\r\nOther changes:\r\n\r\nRenamed OpStream's noIntcBlock and noBytecBlock to hasIntcBlock and hasBytecBlock to clarify this flag is true only if the input code defines a constant block.\r\npushbytes now disassembles with a comment containing a parsed representation of its contents, like bytec.\r\n\r\n* combine app state key/value size limits (#2172)\r\n\r\nCombine the limits enforced on keys and values in Teal such that there is a 128 byte limit on a key/value pair, rather than a 64 byte limit on each. The previous method meant app creators/users were essentially paying for (\"renting\") space that might be used by keys, but rarely was, and even if it was, could serve little purpose.\r\n\r\n* Catchupaccessor coverage (#2235)\r\n\r\ntesting: improve Catchupaccessor code coverage\r\n\r\n* Expose creatable IDs to TEAL code within the same group (#2243)\r\n\r\nThis PR adds gaid and gaids opcodes (similar to gtxn and gtxns), which will allow smart contracts to access asset IDs of assets/apps which were created earlier in the same transaction group.\r\n\r\nPreviously, accessing creatable IDs required subsequent app calls with asset/application IDs as app arguments.\r\n\r\n* Allow AVM code to access a max number of foreign refs. (#2263)\r\n\r\nAlong the lines of similar \"combine\" PRs, this allows AVM code to access a total number of \"foreign\" references, as opposed to limiting each one individually.\r\n\r\n* deprecate auction code (#2261)\r\n\r\nRIP Auction code. We should create a tag when this is merged. It is also archived in https://github.com/algorand/auction-tools\r\n\r\n* wrote tests\r\n\r\n* Temporarily disabled TestBasicCatchpointCatchup and reset test timeouts from 60 to 10 seconds. (#2266)\r\n\r\n* Add support for RHEL Centos 8 RPM  (#2190)\r\n\r\n* lint\r\n\r\n* Disable TestPeriodicSync test (#2269)\r\n\r\nTemporarily disable TestPeriodicSync test since it is broken and blocking the release pipeline.\r\n\r\n* Fix incorrect `gaid` and `gaids` docs (#2275)\r\n\r\nThe gaid and gaids docs refer to the deprecated CreatableID transaction field that was used by the txn opcode. This PR fixes the documentation and adds additional information about only being able to access IDs of previous transactions in the current group.\r\n\r\n* Regularize access to \"foreign\" references. (#2264)\r\n\r\nAll opcodes that take accounts, asas, or apps can use the thing itself (an address, or asa/app id) or an \"indirect\" reference through the \"foreign\" arrays of the app transaction. In all cases, the \"thing\" must appear in those foreign arrays, it's just more convenient sometimes to use the item in code, rather than indirect. (Note this is a new requirement for some opcodes that previously allowed access to any number of local state objects of an account asset_holding_get asset_opted_in and app_local_*)\r\n\r\n* fix random failure (#2280)\r\n\r\nThe unit test had a corner case that would not ensure that all the rounds were flushed to disk.\r\n\r\n* Make account endpoint produce deterministic output (#2276)\r\n\r\nThe JSON return value of the /v2/accounts/{addr} endpoint represents an account with several arrays which are populated from Go maps (such as held assets, created apps, etc.). Currently these arrays are unordered, meaning every time you query this endpoint, you'll likely receive a response with a different ordering of these arrays. This makes SDK testing more difficult than it needs to be, since different responses cannot be directly compared to each other.\r\n\r\nThis PR sorts the slices in generated.Account so that the /v2/accounts/{addr} endpoint will always produce arrays with deterministic order.\r\n\r\n* In this chage, fixes to peer selector and the test.\r\n\r\npeerSelector.go: various bug fixes\r\n- introduce peerSelectorPeer to wrap the network.Peer and add peerClass information, to be able to distinguish between peers of the same address but different classes.\r\n- keep track of download failures to be able to increase the cost of each failure when failing more than succeeding. This is to evict the peer faster when constantly failing to download.\r\n- initialize rankSum and rankSamples to initialRank of the class. Otherwise, the peer rank will have a very long warmup time before relfecting the correct rank.\r\n- let resetRequestPenalty bound the rank within the class bounds. Otherwise, the penalty calculation pushes the rank out of the class bounds (bug).\r\n- getNextPeer is local to the package\r\n- getNextPeer, PeerDownloadDurationToRank and RankPeer use peerSelectorPeer instead of network.Peer\r\n- refreshAvailablePeers distinguishes between peers with the same address but of different peer class\r\n- findPeer returns the peer given the address and the peer class (instead of just the address)\r\n\r\ncatchpointCatchup_test.go:\r\n- Remove comment about giving the second node all the stake, since it is not the case here.\r\n- Use the round from the catchpoint instead of guessing the round as 36. In case the following catchpoint was obtained due to race conditions, checking for round 37 will be trivial, since it will also be obtained from the catchpoint.\r\n\r\ncatchpointService.go and service.go:\r\n- Update the code to use peerSelectorPeer instead of network.Peer with peerSelector\r\n\r\npeerSelector_test.go:\r\n- Update the tests to use peerSelectorPeer instead of network.Peer with peerSelector\r\n- Cleanup debugging printouts.\r\n\r\n* write tests\r\n\r\n* fix import\r\n\r\n* rename test\r\n\r\n* Added tests, exponential increase of download failure impact, local functions.\r\n\r\n* test for error first\r\n\r\n* Fix the test name so it runs\r\n\r\n* final v28 foundation spec (#2286)\r\n\r\nChanges to readmes and specs to reflect the matching foundation spec.\r\n\r\n* fix tests\r\n\r\n* report substring missing immediates properly (#2287)\r\n\r\nFixes crash that should be a clean error report for using substring wrong.\r\n\r\nUnit tests added to confirm and prevent regression.\r\n\r\n* add check\r\n\r\n* Fix darwin-arm64 builds\r\n\r\n* address comments\r\n\r\n* Enable unit tests on mac build script.\r\n\r\n* More changes.\r\n\r\n* Introduce V28 consensus version (#2255)\r\n\r\n* TEAL v4\r\n* Larger programs\r\n* Larger app/asset lookup limits\r\n* Longer asset URL\r\n* Fee pooling within a group\r\n* Keyreg txn additional checks\r\n\r\nRemoved InitialRewardsRateCalculation and PaysetCommit from vFuture since it is already in v26\r\n\r\nFixed some tests after exposing zero fees and strict keyreg as current consensus\r\n\r\n* Forgot to run .md generation (#2292)\r\n\r\n* add another case\r\n\r\n* Allow fee to be below minfee, if given explicitly. (#2295)\r\n\r\nThis makes other txns accept explicitly low fees.  Sorry for the code\r\nduplication, but I did not want to change libgoal's existing behavior\r\nwhere it increases fee to minfee.\r\n\r\n* Expose extra program pages to API (#2294)\r\n\r\nAdd the AppsTotalExtraPages account field to the response returned by the /v2/accounts/{addr} endpoint.\r\n\r\n* use enums for message event permutations, add require trace helpers\r\n\r\n* add playerPermutation enums to permutation test\r\n\r\n* rename requireTraceContainsAction to requireTraceContains\r\n\r\n* testing: fix random failure in TestAppEmptyAccountsLocal (#2302)\r\n\r\nThe test had two unrelated bugs:\r\n1. We need to call `WaitForCommit` before `reloadLedger` to ensure the block is being written to disk before the blockQ getting reinitialized ( and loose its content ).\r\n2. The calculation of the total rewards unit in `makeNewEmptyBlock` was wrong. I corrected it. For tests that run only one or two rounds, this might be good enough, but for long-running tests, it would start fail pretty quickly.\r\n\r\n* Update the Version, BuildNumber, genesistimestamp.data\r\n\r\n* make linter happy by removing underscores from enum names\r\n\r\n* Remove commented code in TestPlayerPermutation\r\n\r\n* Move extra page test to e2e_subs\r\n\r\n* Specify truncated division is used in TEAL\r\n\r\nSpecify that the truncated division is used (https://en.wikipedia.org/wiki/Modulo_operation#Variants_of_the_definition)\r\n\r\n* typos, grammar, inconsistencies\r\n\r\nmostly casing issues, ie, algorand > Algorand || sqlite > SQLite\r\n\r\n* Spec updates to go with division explanation.\r\n\r\n* added benchmark to lruAccounts write function\r\n\r\n* lruaccounts benchmark: filling the acounts with data before benchmarking\r\n\r\n* fix: amount of accounts generated in lruAccounts write benchmark\r\n\r\n* lruAccounts write fix: benchmark used too much memory\r\n\r\n* lruAccounts benchmark: a more controlled distribution between accounts in the benchmark\r\n\r\n* REST API: make extra-program-pages and apps-total-extra-pages optional\r\n\r\n* This also helps in not exposing them before the protocol switch\r\n\r\n* Merge pull request #2313 from algorandskiy/pavel/extra-pages-api\r\n\r\nREST API: make extra-program-pages and apps-total-extra-pages optional\r\n\r\n* Bump buildnumber.dat\r\n\r\n* Run misspell linter with -w flag (#2320)\r\n\r\nThis runs the popular misspell linter with the -w flag to automatically correct spelling mistakes in go-algorand.\r\n\r\n* Reduce unneeded contention around checking the peers connectivity. (#2319)\r\n\r\nExisting `messageHandlerThread` was checking that all the connected peers are properly communicating by examining their recent message timings. This implementation served us well - however, it was executed redundantly. How much redundantly ? 19 times too many every 3 minutes ( and all of them at the **exact** same time.. ).\r\n\r\nThis PR ensures that all the `messageHandlerThread` shares the same ticker for testing the `checkPeersConnectivity`. This is expected to reduce the pressure on the internal `peersLock`.\r\n\r\n* Improve TestMetricSegment test realibility (#2322)\r\n\r\nmprove TestMetricSegment test reliability by repeating the test with incrementing time delays.\r\nThis implementation would allow faster execution on faster platforms, and allow fallback for slower platforms.\r\n\r\n* Restore TestPeriodicSync and TestBasicCatchpointCatchup tests. (#2315)\r\n\r\n* testing: enable previously disabled TestConfigMigrate unit test (#2326)\r\n\r\ntesting: enable previously disabled TestConfigMigrate unit test\r\n\r\n* testing: fix gotestsum install (#2328)\r\n\r\ntesting: fix gotestsum install\r\n\r\navoid installing swagger when not needed.\r\n\r\n* testing: move travis_retry to skip rebuilding (#2324)\r\n\r\nRemove the top-level retry command. This will allow build / lint failures to terminate the build with no retry, and may speed up test failure retries.\r\n\r\n* testing: fix telemetry unit tests (#2321)\r\n\r\nFix and re-enable async telemetry unit test TestAsyncTelemetryHook_CloseDrop.\r\n\r\n* Add Fedora support into install_linux_deps.sh (#2331)\r\n\r\nAdd Fedora deps into install_linux_deps.sh\r\n\r\n* Add missing txna allowed args to doc and langspec (#2336)\r\n\r\nWe exposed Assets and Applications in TEAL v3 but did not update doc/langspec. This commit fixes it.\r\n\r\n* Add initial devMode support (#2334)\r\n\r\nAdd new local private network mode - devmode. devMode allows the developer to deploy a single node network, where every transaction being sent to the node automatically generates a new block.\r\n\r\nThis feature is focused primarily around 3rd parties that want to test their solution on the Algorand platform, without waiting for the network to make progress.\r\n\r\n* Implement hello circleci yaml (#2417)\r\n\r\nThis adds a hello world circleci yaml. This can be used to run circleci jobs on PRs until the pipeline is fully implemented.\r\n\r\n* catchup: fix potential nil dereferencing (#2420)\r\n\r\nThe handing for returned peer was incorrect in case the getPeerErr is non-nil, as we would attempt to dereference the nil pointer.\r\n\r\n* testing: ensure deploy_linux_version generates a linux/amd64 binaries for algonet usage (#2422)\r\n\r\nThe deploy_linux_version.sh script currently creates a linux version of the current project tree with the same architecture as the hosting environment.\r\n\r\nWhen we attempt to builds on M1 Macs, it default to arm64 base images and as a result - makes a linux/arm64 compatible binaries. These resulting binaries, however, would not work correctly on an algonet deployed network, since the hosts there are amd64.\r\n\r\nTo rectify that situation, we'll be changing the docker file to ensure linux/amd64 binaries are generated. This would ensure that the existing functionality works as intended. In the future, when we would add arm64 support for algonet, we could accompany that by creating corresponding images on docker as well.\r\n\r\n* Specs and unit test to avoid forgetting in the future (#2418)\r\n\r\nSpecs and unit test to avoid forgetting in the future\r\n\r\n* Fix and enable TestNewAccountCanGoOnlineAndParticipate (#2238)\r\n\r\n* TestNewAccountCanGoOnlineAndParticipate was failing because the test was\r\nnot waiting enough to get to the round where the newly funded account's\r\nfunds will be considered for proposing purposes.\r\n\r\nIt was miscalculating the round that it should wait form.\r\n\r\nMoreover, the rounds considered to when the account is funded was prone to\r\nrace conditions.\r\n\r\nIn addition, the test was using WaitForRoundWithTimeout which may be\r\nvery slower if the current round is already ahead. Instead, now it is\r\nusing ClientWaitForRound, which does not care about individual rounds delayed.\r\n\r\n* Addressing review commnets:\r\n- fixing a typo\r\n- getting exact transaction round for funding the account\r\n- testing exact blocks for the proposer\r\n- using a single node network instead of two nodes\r\n- waiting for exactly one round for the new account to propose and checking that\r\n- sending the funds and closing the rich account so there will be no possiblity of that proposing a block\r\n\r\n* Fix docker repo update (#2342)\r\n\r\nThe scripts to update the docker repository do not do a full rebuild, and the submitted Dockerfile does not change. This means that if you run the docker image/repo update, it'll re-use the old image. To fix this, we call docker build with --no-cache. Other refactoring helps simplify the code.\r\n\r\nWe add a new --cached flag to the docker/releases/build_releases.sh script to explicitly cache. Otherwise the testnet update would issue a full rebuild. We additionally automatically handle the 'latest' tagging in build_releases.sh as well.\r\n\r\n* Bump Version, Remove buildnumber.dat and genesistimestamp.dat files.\r\n\r\n* Initialize past side effects with correct length during dryrun requests (#2448)\r\n\r\nThis PR fixes a typo in the dryrun code that caused certain dryrun requests containing multiple transactions to fail.\r\n\r\n* optimize txtail memory consumption (#2413)\r\n\r\nThe changes in this PR are as follows:\r\n\r\nThe unused method Ledger. GetRoundTxIds was removed. As a result, the txTail. getRoundTxIds can be removed as well. This makes the txids map stored in the roundTxMembers structure redundant.\r\nIn the cow.go, avoid adding empty leases to the cb.mods.Txleases map. Since we already not testing for empty leases, we can safely avoid storing them.\r\nOptimize the txTail. loadFromDisk to generate optimal lastValid map sizes.\r\nOptimize the txTail. loadFromDisk to avoid storing empty leases in the txleases map.\r\n\r\n* enable TestPartkeyOnlyRewards on macos (#2429)\r\n\r\nTestPartkeyOnlyRewards was disabled on darwin.\r\nIt is no longer failing on darwin (5/5 runs passed).\r\n\r\n* Fix 10 minute timeout in travis. (Revert parts of #2324) (#2494)\r\n\r\nWe started seeing the 10-minute timeout error on travis after the recent changes to where travis_retry gets called.\r\n\r\nThis means travis_wait is still needed. travis_wait and travis_retry don't play well together on the ephemeral build machine, so we basically need to rollback the entire change in #2324\r\n\r\nNote: while making this change I noticed that we don't use travis_retry for all of the build_test.sh / integration_test.sh entries. Not sure why but I left them as they were before #2324\r\n\r\n* allow to parse ipv6 localhost \"[::]:4601\" (#2430)\r\n\r\nAllow parsing of ipv6 localhost -colon- port address.\r\n\"[::]:4601\" was failing to parse. This and other forms pass ParseHostOrURL() now.\r\n\r\n* disable TestAgreementSynchronous10 (#2503)\r\n\r\nThe test TestAgreementSynchronous10 is failing. We should be fixing it, but until we do that, I'm going to disable it so it won't mask other issues.\r\n\r\n* testing: avoid division by zero during TestBasicCatchpointWriter (#2502)\r\n\r\nThe `randomFullAccountData` method was dividing by `lastCreatableID` which could be zero.\r\nThe probability for that is pretty slim, but given that it was found during a travis run, we should fix it.\r\n\r\n* selector pseudo-op in support of ABI (#2358)\r\n\r\nAdds a Teal pseudo-op `selector` that assembles as if it were the `byte` pseudo-op, but stores 4 bytes of hash.\r\n\r\nThis allows Teal, like:\r\n```\r\ntxn ApplicationArgs 0\r\nselector \"add(uint64,uint64)uint128\"\r\n==\r\nbnz add\r\n```\r\n\r\nand avoid the need to embed the actual hash in a .teal file (which would also require calculating it)\r\n\r\nUnlike `byte`, the argument *must* be a quoted string - no base64 or hex allowed, as the argument should be a method signature.\r\n\r\n* Creator access and app_params_get (#2301)\r\n\r\nAdds the ability to get the creator of apps and assets in teal.\r\n\r\nThis is a simple new field for `asset_params_get`, but it introduces a new opcode `app_params_get` to enable access for apps.  Meanwhile `app_params_get` allows access to other global parameters about an app - the schema sizes, extra pages, and the programs themselves.\r\n\r\nThis also begins work on LogicVersion=5, which should be AVM 1.0\r\n\r\n* fix regex for matching charset of valid DNS hosts, add test (#2505)\r\n\r\nA regex was trying to match valid DNS names but missed the '-' char. Fix.\r\n\r\n* Fix random failure in TestWebsocketNetworkPrioLimit (#2509)\r\n\r\nThe peers array is modified when adding/removing entries from it. When that does happen, we increase the peersChangeCounter, so that the broadcast method would know that it's peers list need to be refreshed.\r\nThe said update was missing from prioTracker.setPriority, which was causing the issue.\r\n\r\n* Make TestPeersDownloadFailed and TestHistoricData predictable (#2516)\r\n\r\nBoth the `TestPeersDownloadFailed` as well as `TestHistoricData` were randomly failing, as they were depending on a \"genuine\" random distribution of the random function in order to succeed. When that doesn't happen, they were failing.\r\n\r\nThis PR make sure to \"bake-in\" the random seed so that the tests are repeatable and a success is repeatable as well.\r\n\r\n* Improve testing of alloc bounded slices. (#2515)\r\n\r\nThis change is needed in order to support codec types such as the following:\r\n```golang\r\n//msgp:allocbound typeA 16\r\ntype typeA []typeB\r\n```\r\nsince we want the codec object randomizer to recognize that `typeA` has a allocbound defined, and use that bound.\r\n( this would be used in the feature/txnsync branch, but the change doesn't really related to any of the other changes in the feature/txnsync branch and would be a good change regardless )\r\n\r\n* Save Logging Information in Data Directory If Provided (#2415)\r\n\r\nOur logic should be as follows:\r\n\r\nWhen Loading:\r\n\r\nWe first look inside the provided data-directory. If a config file is there, load it and return it\r\nOtherwise, look in the global directory. If a config file is there, load it and return it.\r\nWhen Saving:\r\n\r\nIf a data-directory was provided then save the config file there.\r\nOtherwise, save the config file in the global directory\r\n\r\n* Lruaccts write speedup (#2329)\r\n\r\nOptimized lruAccounts write function. gains: from ~260ns to ~200ns per write op\r\n\r\n* Use golangci-lint (#2523)\r\n\r\nThis adds a golangci-lint configuration file that runs an initial minimal set of linters: golint, govet, and misspell. This makes it easier to integrate with editors and IDEs and also opens the door to using from CI.\r\n\r\n* update codec tester (#2527)\r\n\r\nThis PR improves the checkBoundsLimitingTag method of the message pack random generated object tester.\r\nWith this change, we're no longer ignoring data types that has no struct tags, but rather looking for corresponding msgp directives.\r\n\r\nThis is the same change applied to the feature/txnsync branch. This change is expected to be nop on the master branch, but would be required on the feature branch.\r\n\r\n* Fix random failures in TestPeriodicSync (#2535)\r\n\r\nThe waiting period wasn't long enough. Allow much longer wait period before giving up.\r\n\r\nAlso - omit unneeded log entries.\r\n\r\n* Testing: use periodicSyncLogger for cleaner test runs (#2539)\r\n\r\nUse periodicSyncLogger in a test to avoid unnecessary log outputs.\r\n\r\n* Add travis wait to compilation step on travis. (#2544)\r\n\r\nTravis nightly build failed due to no-output for over 10m from build process.\r\n\r\n* test conditional slack alert (#2537)\r\n\r\nTravis is currently unable to only notify on failures of a particular branch, however, it can notify on us on the build of a particular branch (both success and failures). Since we are also considering CircleCI, I think it would be best to just add the slack notification into the travis file now for the rel/nightly branch and notify the #devops-jenkins channel (where most of our jobs status activity goes) when it both succeeds/fails. The \"successes\" will also help us determine whether or not the test ran at all.\r\n\r\n* Add Extract opcodes  (#2521)\r\n\r\nAdds extract opcodes that allow a substring to be extracted given a start index and a length (#2347).\r\n\r\nThe specs for the extract and extract3 opcodes are the same as the substring and substring3 opcodes, but takes a length rather than an end index. This also adds extract16bits, extract32bits, extract64bits which extracts 2, 4, and 8 byte strings respectively and converts them into integers.\r\n\r\n* Benchmark\r\n\r\n* Modify target\r\n\r\nCo-authored-by: DevOps Service <devops-service@algorand.com>\r\nCo-authored-by: John Lee <64482439+algojohnlee@users.noreply.github.com>\r\nCo-authored-by: Tsachi Herman <tsachi.herman@algorand.com>\r\nCo-authored-by: algonautshant <55754073+algonautshant@users.noreply.github.com>\r\nCo-authored-by: Will Winder <wwinder.unh@gmail.com>\r\nCo-authored-by: chris erway <51567+cce@users.noreply.github.com>\r\nCo-authored-by: Brian Olson <brianolson@users.noreply.github.com>\r\nCo-authored-by: algobarb <78746954+algobarb@users.noreply.github.com>\r\nCo-authored-by: John Jannotti <john.jannotti@algorand.com>\r\nCo-authored-by: shiqizng <80276844+shiqizng@users.noreply.github.com>\r\nCo-authored-by: Jacob Daitzman <jdtzmn@gmail.com>\r\nCo-authored-by: nicholasguoalgorand <67928479+nicholasguoalgorand@users.noreply.github.com>\r\nCo-authored-by: egieseke <eric_gieseke@yahoo.com>\r\nCo-authored-by: John Lee <john.lee@algorand.com>\r\nCo-authored-by: bricerisingalgorand <60147418+bricerisingalgorand@users.noreply.github.com>\r\nCo-authored-by: Jason Paulos <jasonpaulos@users.noreply.github.com>\r\nCo-authored-by: Jacob Daitzman <jdtzmn@users.noreply.github.com>\r\nCo-authored-by: Nicholas Guo <nicholas.guo@algorand.com>\r\nCo-authored-by: algonautshant <shant@algorand.com>\r\nCo-authored-by: Rakshith Gopala Krishna <rakshith.gopalakrishna@algorand.com>\r\nCo-authored-by: Pavel Zbitskiy <65323360+algorandskiy@users.noreply.github.com>\r\nCo-authored-by: chris erway <chris.erway@algorand.com>\r\nCo-authored-by: Pavel Zbitskiy <pavel@algorand.com>\r\nCo-authored-by: Fabrice Benhamouda <fabrice.benhamouda@normalesup.org>\r\nCo-authored-by: figurestudios <64747030+figurestudios@users.noreply.github.com>\r\nCo-authored-by: John Jannotti <jj@cs.brown.edu>\r\nCo-authored-by: Jonathan Weiss <jonathan@Algo-Weiss-MBP.local>\r\nCo-authored-by: Jonathan Weiss <jonathan.weiss@algorand.com>\r\nCo-authored-by: algonathan <85506383+algonathan@users.noreply.github.com>\r\nCo-authored-by: pzbitskiy <pavel.zbitskiy@gmail.com>\r\nCo-authored-by: AlgoStephenAkiki <85183435+AlgoStephenAkiki@users.noreply.github.com>",
          "timestamp": "2022-05-27T21:06:53-04:00",
          "tree_id": "6cc085936be26e4a5c8960a4e53695cdbe479d2a",
          "url": "https://github.com/algochoi/go-algorand/commit/a8b32bc42c669ba31e26abbf1a374f2261b73f15"
        },
        "date": 1653700461904,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUintMath/pop1",
            "value": 55.9,
            "unit": "ns/op\t        15.0 waste/op",
            "extra": "21415906 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop",
            "value": 105,
            "unit": "ns/op\t        45.0 waste/op",
            "extra": "11560170 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/add",
            "value": 110,
            "unit": "ns/op\t        45.0 waste/op",
            "extra": "10762449 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sub",
            "value": 112,
            "unit": "ns/op\t        45.0 waste/op",
            "extra": "10954622 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mul",
            "value": 118,
            "unit": "ns/op\t        45.0 waste/op",
            "extra": "10234029 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/div",
            "value": 118,
            "unit": "ns/op\t        45.0 waste/op",
            "extra": "10343334 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divmodw",
            "value": 1004,
            "unit": "ns/op\t       120 waste/op",
            "extra": "1219371 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sqrt",
            "value": 122,
            "unit": "ns/op\t        30.0 waste/op",
            "extra": "9890746 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/exp",
            "value": 151,
            "unit": "ns/op\t        45.0 waste/op",
            "extra": "7988914 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/expw",
            "value": 550,
            "unit": "ns/op\t        60.0 waste/op",
            "extra": "2166118 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michaeldiamant@users.noreply.github.com",
            "name": "Michael Diamant",
            "username": "michaeldiamant"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c03e3d3bc320c52a847c0e7850b84fe3efda8948",
          "message": "Bump py-algorand-sdk to v1.17.0 (#4530)",
          "timestamp": "2022-09-12T21:03:07-04:00",
          "tree_id": "fd86d9fbea1229d0bd7be61427205834f1a4ff08",
          "url": "https://github.com/algochoi/go-algorand/commit/c03e3d3bc320c52a847c0e7850b84fe3efda8948"
        },
        "date": 1663097300054,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUintMath/dup",
            "value": 53.52,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "22539340 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop1",
            "value": 51.11,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "24744525 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop",
            "value": 96.13,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "12616016 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/add",
            "value": 105.7,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11478013 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/addw",
            "value": 125.8,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "8445790 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sub",
            "value": 101,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "12197786 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mul",
            "value": 101.7,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11721882 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mulw",
            "value": 123.6,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9982303 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/div",
            "value": 112.6,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10865293 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divw",
            "value": 164.6,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "7161152 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divmodw",
            "value": 1160,
            "unit": "ns/op\t         8.000 extra/op\t     311 B/op\t      11 allocs/op",
            "extra": "904822 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sqrt",
            "value": 117.6,
            "unit": "ns/op\t         2.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10258072 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/exp",
            "value": 147.2,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "8173064 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/expw",
            "value": 558.7,
            "unit": "ns/op\t         4.000 extra/op\t     111 B/op\t       5 allocs/op",
            "extra": "2105636 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "wwinder.unh@gmail.com",
            "name": "Will Winder",
            "username": "winder"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "bb7c59f3982eb107a294e8be0e5b7b13962f33a3",
          "message": "Remove dead link from kmd README. (#4608)",
          "timestamp": "2022-10-12T10:06:45-04:00",
          "tree_id": "41dfde3b6c4437a95e75b75aee1364649dbe6b63",
          "url": "https://github.com/algochoi/go-algorand/commit/bb7c59f3982eb107a294e8be0e5b7b13962f33a3"
        },
        "date": 1665599140819,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUintMath/dup",
            "value": 56.01,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "20835853 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop1",
            "value": 53.54,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "19665284 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop",
            "value": 98.35,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "12167847 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/add",
            "value": 107.8,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11832912 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/addw",
            "value": 130.9,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9211603 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sub",
            "value": 107.2,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11811451 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mul",
            "value": 106.5,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11007292 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mulw",
            "value": 130.8,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9312451 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/div",
            "value": 114.3,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10864789 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divw",
            "value": 170.1,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "7086981 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divmodw",
            "value": 1100,
            "unit": "ns/op\t         8.000 extra/op\t     311 B/op\t      11 allocs/op",
            "extra": "1093377 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sqrt",
            "value": 115.7,
            "unit": "ns/op\t         2.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10595743 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/exp",
            "value": 150.9,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "8009139 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/expw",
            "value": 561.9,
            "unit": "ns/op\t         4.000 extra/op\t     111 B/op\t       5 allocs/op",
            "extra": "2189418 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "brianolson@users.noreply.github.com",
            "name": "Brian Olson",
            "username": "brianolson"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": false,
          "id": "36fc7f9b0a0831d2658353ecfaf3702ea6784c6b",
          "message": "algod: turn off cadaver trace file by default (#4676)",
          "timestamp": "2022-10-20T14:56:43-04:00",
          "tree_id": "f653e1a6765743b479be743df9413a8fdd88b38a",
          "url": "https://github.com/algochoi/go-algorand/commit/36fc7f9b0a0831d2658353ecfaf3702ea6784c6b"
        },
        "date": 1666303969287,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUintMath/dup",
            "value": 56.74,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "20838852 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop1",
            "value": 59.21,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "21670654 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop",
            "value": 108.5,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11127402 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/add",
            "value": 115.1,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10741228 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/addw",
            "value": 144.6,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "8736566 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sub",
            "value": 115.6,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9009483 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mul",
            "value": 116.2,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10136781 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mulw",
            "value": 142.5,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "8453287 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/div",
            "value": 122.6,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9933957 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divw",
            "value": 173.4,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "6818972 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divmodw",
            "value": 1185,
            "unit": "ns/op\t         8.000 extra/op\t     311 B/op\t      11 allocs/op",
            "extra": "896308 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sqrt",
            "value": 120.5,
            "unit": "ns/op\t         2.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10111742 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/exp",
            "value": 152.4,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "7789822 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/expw",
            "value": 626.8,
            "unit": "ns/op\t         4.000 extra/op\t     111 B/op\t       5 allocs/op",
            "extra": "1892395 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "91566643+algoidurovic@users.noreply.github.com",
            "name": "algoidurovic",
            "username": "algoidurovic"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": false,
          "id": "ab87a8a94c7ab9c3ae5d49bde30576f709344ad7",
          "message": "AVM: match, pushints, and pushbytess opcodes (#4645)\n\nadd match opcode along with assembler and eval unit tests\r\nimplementation of match, pushints, and pushbytess\r\nadd docs for pushints and pushbytess",
          "timestamp": "2022-10-26T15:36:09-04:00",
          "tree_id": "ff44f2e8f556675099b2e93e6248cc58d43c074c",
          "url": "https://github.com/algochoi/go-algorand/commit/ab87a8a94c7ab9c3ae5d49bde30576f709344ad7"
        },
        "date": 1666882086175,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkUintMath/dup",
            "value": 52.44,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "23352050 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop1",
            "value": 48.58,
            "unit": "ns/op\t         1.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "24674104 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/pop",
            "value": 93.42,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "12701194 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/add",
            "value": 98.41,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "12580393 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/addw",
            "value": 124.7,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9997398 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sub",
            "value": 97.76,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10884740 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mul",
            "value": 99.43,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "11811596 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/mulw",
            "value": 123.5,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "9924478 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/div",
            "value": 108.6,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10973589 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divw",
            "value": 162,
            "unit": "ns/op\t         4.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "7021035 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/divmodw",
            "value": 1159,
            "unit": "ns/op\t         8.000 extra/op\t     311 B/op\t      11 allocs/op",
            "extra": "964878 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/sqrt",
            "value": 113.2,
            "unit": "ns/op\t         2.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "10797085 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/exp",
            "value": 143.8,
            "unit": "ns/op\t         3.000 extra/op\t       7 B/op\t       0 allocs/op",
            "extra": "8369952 times\n2 procs"
          },
          {
            "name": "BenchmarkUintMath/expw",
            "value": 557.3,
            "unit": "ns/op\t         4.000 extra/op\t     111 B/op\t       5 allocs/op",
            "extra": "2135839 times\n2 procs"
          }
        ]
      }
    ]
  }
}