#! /usr/bin/env python3

"""create a 'fingerprint' file for this python installation.
This is used on CircleCI as part of the cache key. The problem we've
seen there was that a micro version upgrade broke the virtualenv's in
.nox (which we don't rebuild in order to have better caching)
This program print the python version as well as the sha256 sum of the
python executable and the path to the executable. We use CircleCI's
checksum of this file as part of the cache key.
"""

import hashlib
import pathlib
import sys


def main():
    h = hashlib.sha256()
    executable_path = pathlib.Path(sys.executable).resolve()
    with open(executable_path, "rb") as f:
        h.update(f.read())

    print(sys.version)
    print(f"{h.hexdigest()}  {executable_path}")


if __name__ == "__main__":
    main()
