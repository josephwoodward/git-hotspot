# Git Hotspot Extension

Git Hotspot is a git extension that identifies 'hot' files (tracked files in a git repository that are modified the most).

Why is this useful? For a few reasons:

- When starting to learn a new codebase, I often find it helpful to begin by reviewing files which are edited most often as it's likely you'll spend some time in them yourself.
- Files subject to a large amount of churn are often good candidates to refactor as it often an indicator it's doing too much. This is especially true in languages that favor one type/class per file.

## Usage 

```bash
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

## Installation

Requirements:
- git
- Go version 1.18 or above

Clone this repo then run the following in your CLI:

```bash
$ make build  
$ make install
```

Now navigate to a git repository and run `git hotspot`.

