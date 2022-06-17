# Git Hotspot

Git Hotspot is a git extension that you can use to identify which files in a git repository are modified the most.

For example:

```
$ git clone https://github.com/grafana/tempo.git
$ cd tempo
$ git hotspot

Files                                                # Modifications
operations/jsonnet/microservices/config.libsonnet    51
operations/jsonnet/microservices/ingester.libsonnet  23
modules/querier/querier_test.go                      21
modules/overrides/overrides.go                       17
tempodb/encoding/common/types.go                     16
tempodb/encoding/finder_paged.go                     11
tempodb/pool/pool.go                                 11
.goreleaser.yml                                      7
GOVERNANCE.md                                        7
example/tk/jsonnetfile.json                          6
tempodb/search/data_combiner.go                      4
.drone/drone.jsonnet                                 4
tempodb/encoding/iterator_multiblock_test.go         4
tempodb/search/searchable_block.go                   3
integration/bench/stress_test_write_path.js          3
```

### Installation

Requires:
- git
- Go version 1.18

Clone this repo then run the following in your CLI:

```
$ make build  
$ make install
```

Now navigate to a git repo and run `git hotspot`.

