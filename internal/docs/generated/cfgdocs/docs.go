// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package cfgdocs

var READMEShort = `View and Modify Resource Configuration.`
var READMELong = `
Programmatically print and modify raw json or yaml Resource Configuration

Commands: [annotate], [cat], [count], [create-setter], [fmt], [grep], [list-setters],
[merge], [merge3], [set], [tree]

| Command        | Description                                   |
|----------------|-----------------------------------------------|
| [annotate]     | set ` + "`" + `metadata.annotation` + "`" + `s on Resources       |
| [cat]          | print Resources in a package                  |
| [count]        | print Resource counts by type                 |
| [create-setter]| create or modify a custom field-setter        |
| [fmt]          | format Resource yaml                          |
| [grep]         | filter Resources configuration                |
| [list-setters] | list setters                                  |
| [merge]        | merge Resources in one directory into another |
| [merge3]       | perform 3-way merge between directories       |
| [set]          | set one or more fields programmatically       |
| [tree]         | print Resources using a tree structure        |
`
var READMEExamples = `
    # print the raw package contents
    $ kpt cfg cat helloworld

    # print the package using tree based structure
    $ kpt cfg tree helloworld --name --image --replicas
    helloworld
    ├── [deploy.yaml]  Deployment helloworld-gke
    │   ├── spec.replicas: 5
    │   └── spec.template.spec.containers
    │       └── 0
    │           ├── name: helloworld-gke
    │           └── image: gcr.io/kpt-dev/helloworld-gke:0.1.0
    └── [service.yaml]  Service helloworld-gke

    # only print Services
    $ kpt cfg grep "kind=Service" helloworld | kpt cfg tree --name --image --replicas
    .
    └── [service.yaml]  Service helloworld-gke

    #  list available setters
    $ kpt cfg list-setters helloworld replicas
        NAME          DESCRIPTION        VALUE    TYPE     COUNT   SETBY
      replicas   'helloworld replicas'   5       integer   1

    # set a high-level knob
    $ kpt cfg set helloworld replicas 3
    set 1 fields
`

var AnnotateShort = `Set an annotation on one or more Resources`
var AnnotateLong = `
  DIR:
    Path to local directory.
`
var AnnotateExamples = `
    # set an annotation on all Resources: 'key: value'
    kpt cfg annotate DIR --kv key=value

    # set an annotation on all Service Resource
    kpt cfg annotate DIR --kv key=value --kind Service

    # set an annotation on the foo Service Resource
    kpt cfg annotate DIR --kv key=value --kind Service --name foo

    # set multiple annotations
    kpt cfg annotate DIR --kv key1=value1 --kv key2=value2`

var CatShort = `Print Resource Config from a local directory.`
var CatLong = `
    kpt cfg cat DIR

  DIR:
    Path to local directory.
`
var CatExamples = `
    # print Resource config from a directory
    kpt cfg cat my-dir/`

var CountShort = `Count Resources Config from a local directory.`
var CountLong = `
    kpt cfg count DIR

  DIR:
    Path to local directory.
`
var CountExamples = `
    # print Resource counts from a directory
    kpt cfg count my-dir/`

var CreateSetterShort = `Create a custom setter for a Resource field`
var CreateSetterLong = `
    kpt cfg create-setter DIR NAME VALUE

  DIR

    A directory containing Resource configuration.

  NAME

    The name of the setter to create.

  VALUE

    The current value of the field, or a substring within the field.
`
var CreateSetterExamples = `
    # create a setter for port fields matching "8080"
    kpt cfg create-setter DIR/ port 8080 --type "integer" --field port \
         --description "default port used by the app"

    # create a setter for a substring of a field rather than the full field -- e.g. only the
    # image tag, not the full image
    kpt cfg create-setter DIR/ image-tag v1.0.1 --type "string" \
        --field image --description "current stable release"

[demo]: https://storage.googleapis.com/kpt-dev/docs/cfg-create-setter.gif "kpt cfg create-setter"`

var FmtShort = `Format yaml configuration files.`
var FmtLong = `
Fmt will format input by ordering fields and unordered list items in Kubernetes
objects.  Inputs may be directories, files or stdin, and their contents must
include both apiVersion and kind fields.

- Stdin inputs are formatted and written to stdout
- File inputs (args) are formatted and written back to the file
- Directory inputs (args) are walked, each encountered .yaml and .yml file
  acts as an input

For inputs which contain multiple yaml documents separated by \n---\n,
each document will be formatted and written back to the file in the original
order.

Field ordering roughly follows the ordering defined in the source Kubernetes
resource definitions (i.e. go structures), falling back on lexicographical
sorting for unrecognized fields.

Unordered list item ordering is defined for specific Resource types and
field paths.

- .spec.template.spec.containers (by element name)
- .webhooks.rules.operations (by element value)
`
var FmtExamples = `
	# format file1.yaml and file2.yml
	kpt cfg fmt file1.yaml file2.yml

	# format all *.yaml and *.yml recursively traversing directories
	kpt cfg fmt my-dir/

	# format kubectl output
	kubectl get -o yaml deployments | kpt cfg fmt

	# format kustomize output
	kustomize build | kpt cfg fmt`

var GrepShort = `Search for matching Resources in a directory or from stdin`
var GrepLong = `
    kpt cfg grep QUERY DIR

  QUERY:
    Query to match expressed as 'path.to.field=value'.
    Maps and fields are matched as '.field-name' or '.map-key'
    List elements are matched as '[list-elem-field=field-value]'
    The value to match is expressed as '=value'
    '.' as part of a key or value can be escaped as '\.'

  DIR:
    Path to local directory.
`
var GrepExamples = `
    # find Deployment Resources
    kpt cfg grep "kind=Deployment" my-dir/

    # find Resources named nginx
    kpt cfg grep "metadata.name=nginx" my-dir/

    # use tree to display matching Resources
    kpt cfg grep "metadata.name=nginx" my-dir/ | kpt cfg tree

    # look for Resources matching a specific container image
    kpt cfg grep "spec.template.spec.containers[name=nginx].image=nginx:1\.7\.9" my-dir/ | kpt cfg tree`

var ListSettersShort = `List setters for Resources.`
var ListSettersLong = `
    kpt cfg list-setters DIR [NAME]

  DIR

    A directory containing Resource configuration.

  NAME

    Optional.  The name of the setter to display.
`
var ListSettersExamples = `
  Show setters:

    $ kpt cfg list-setters DIR/
        NAME      DESCRIPTION   VALUE     TYPE     COUNT   SETBY  
    name-prefix   ''            PREFIX    string   2

[demo]: https://storage.googleapis.com/kpt-dev/docs/cfg-set.gif "kpt cfg set"`

var MergeShort = `Merge Resource configuration files`
var MergeLong = `
Merge reads Kubernetes Resource yaml configuration files from stdin or sources packages and write
the result to stdout or a destination package.

Resources are merged using the Resource [apiVersion, kind, name, namespace] as the key.  If any of
these are missing, merge will default the missing values to empty.

Resources specified later are high-precedence (the source) and Resources specified
earlier are lower-precedence (the destination).

For information on merge rules, run:

	kpt cfg docs merge
`
var MergeExamples = `
    cat resources_and_patches.yaml | kpt cfg merge > merged_resources.yaml`

var Merge3Short = `Merge diff of Resource configuration files into a destination (3-way)`
var Merge3Long = `
Merge3 performs a 3-way merge by applying the diff between 2 sets of Resources to a 3rd set.

Merge3 may be for rebasing changes to a forked set of configuration -- e.g. compute the difference between the original
set of Resources that was forked and an updated set of those Resources, then apply that difference to the fork.

If a field value differs between the ORIGINAL_DIR and UPDATED_DIR, the value from the UPDATED_DIR is taken and applied
to the Resource in the DEST_DIR.

For information on merge rules, run:

	kpt cfg docs-merge3
`
var Merge3Examples = `
    kpt cfg merge3 --ancestor a/ --from b/ --to c/`

var SetShort = `Set values on Resources fields values.`
var SetLong = `
May set either the complete or partial field value.

` + "`" + `set` + "`" + ` identifies setters using field metadata published as OpenAPI extensions.
` + "`" + `set` + "`" + ` parses both the Kubernetes OpenAPI, as well OpenAPI published inline in
the configuration as comments.

` + "`" + `set` + "`" + ` maybe be used to:

- edit configuration programmatically from the cli
- create reusable bundles of configuration with custom setters

  DIR

    A directory containing Resource configuration.

  NAME

    Optional.  The name of the setter to perform or display.

  VALUE

    Optional.  The value to set on the field.


To print the possible setters for the Resources in a directory, run ` + "`" + `set` + "`" + ` on
a directory -- e.g. ` + "`" + `kpt cfg set DIR/` + "`" + `.

#### Tips

- A description of the value may be specified with ` + "`" + `--description` + "`" + `.
- The last setter for the field's value may be defined with ` + "`" + `--set-by` + "`" + `.
- Create custom setters on Resources, Kustomization.yaml's, patches, etc

The description and setBy fields are left unmodified unless specified with flags.

To create a custom setter for a field see: ` + "`" + `kustomize help config create-setter` + "`" + `
`
var SetExamples = `
  Resource YAML: Name Prefix Setter

    # DIR/resources.yaml
    ...
    metadata:
        name: PREFIX-app1 # {"type":"string","x-kustomize":{"partialFieldSetters":[{"name":"name-prefix","value":"PREFIX"}]}}
    ...
    ---
    ...
    metadata:
        name: PREFIX-app2 # {"type":"string","x-kustomize":{"partialFieldSetters":[{"name":"name-prefix","value":"PREFIX"}]}}
    ...

  List setters: Show the possible setters

    $ config set DIR/
        NAME      DESCRIPTION   VALUE     TYPE     COUNT   SETBY  
    name-prefix   ''            PREFIX    string   2

  Perform set: set a new value, owner and description

    $ kpt cfg set DIR/ name-prefix "test" --description "test environment" --set-by "dev"
    set 2 values

  List setters: Show the new values

    $ config set DIR/
        NAME      DESCRIPTION         VALUE     TYPE     COUNT     SETBY 
    name-prefix   'test environment'   test     string   2          dev

  New Resource YAML:

    # DIR/resources.yaml
    ...
    metadata:
        name: test-app1 # {"description":"test environment","type":"string","x-kustomize":{"setBy":"dev","partialFieldSetters":[{"name":"name-prefix","value":"test"}]}}
    ...
    ---
    ...
    metadata:
        name: test-app2 # {"description":"test environment","type":"string","x-kustomize":{"setBy":"dev","partialFieldSetters":[{"name":"name-prefix","value":"test"}]}}
    ...

[demo]: https://storage.googleapis.com/kpt-dev/docs/cfg-set.gif "kpt cfg set"`

var TreeShort = `Display Resource structure from a directory or stdin.`
var TreeLong = `
kpt cfg tree may be used to print Resources in a directory or cluster, preserving structure

Args:

  DIR:
    Path to local directory directory.

Resource fields may be printed as part of the Resources by specifying the fields as flags.

kpt cfg tree has build-in support for printing common fields, such as replicas, container images,
container names, etc.

kpt cfg tree supports printing arbitrary fields using the '--field' flag.

By default, kpt cfg tree uses Resource graph structure if any relationships between resources (ownerReferences)
are detected, as is typically the case when printing from a cluster. Otherwise, directory graph structure is used. The
graph structure can also be selected explicitly using the '--graph-structure' flag.
`
var TreeExamples = `
    # print Resources using directory structure
    kpt cfg tree my-dir/

    # print replicas, container name, and container image and fields for Resources
    kpt cfg tree my-dir --replicas --image --name

    # print all common Resource fields
    kpt cfg tree my-dir/ --all

    # print the "foo"" annotation
    kpt cfg tree my-dir/ --field "metadata.annotations.foo"

    # print the "foo"" annotation
    kubectl get all -o yaml | kpt cfg tree \
      --field="status.conditions[type=Completed].status"

    # print live Resources from a cluster using owners for graph structure
    kubectl get all -o yaml | kpt cfg tree --replicas --name --image

    # print live Resources with status condition fields
    kubectl get all -o yaml | kpt cfg tree \
      --name --image --replicas \
      --field="status.conditions[type=Completed].status" \
      --field="status.conditions[type=Complete].status" \
      --field="status.conditions[type=Ready].status" \
      --field="status.conditions[type=ContainersReady].status"

###

[demo]: https://storage.googleapis.com/kpt-dev/docs/cfg-tree.gif "kpt cfg tree"`
