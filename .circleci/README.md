# CircleCI dynamic configuration

This directory contains a CircleCI dynamic configuration [1]. `config,yml` calls into a short
clojure script, which generates a CircleCI configuration to be executed, based on which files have
changed.

## Running the script locally

You can run the script locally. You need to first install the clojure CLI tools [2]. Then run

```
clojure -X gen/gen
```

You may also pass additional arguments:

```
clojure -X gen/gen :base '"origin/schmir/dev"' :head '"HEAD"'
```

The script lives inside this directory in the file `gen.clj`. The `get-yaml-files` function chooses the yaml files to be merged.

[1] https://circleci.com/docs/2.0/dynamic-config/

[2] https://clojure.org/guides/getting_started#_installation_on_linux
