if __name__ == "__main__":
    __import__("sys").exit(
        "Do not execute this file directly. Use nox instead, it will know how to handle this file"
    )

import os
import pathlib
import shutil

import nox
from nox.sessions import Session


nox.options.sessions = ["black", "flake8", "mypy", "test_contracts"]

python_paths = [
    "contracts/tests/",
    "noxfile.py",
]
requirements_as_constraints = ["-c", "requirements.txt"]

nox.options.error_on_external_run = True
nox.options.reuse_existing_virtualenvs = True


def install_ganache(session: Session) -> None:
    """install ganache-cli"""
    session.install("nodeenv")
    assert session.bin is not None
    nodeenv_dir = pathlib.Path(session.bin).parent.joinpath("node")
    bindir = nodeenv_dir.joinpath("bin").absolute()

    ganache_cli = bindir.joinpath("ganache-cli")
    os.environ["PATH"] = str(bindir) + os.pathsep + os.environ["PATH"]
    session.env["PATH"] = str(bindir) + os.pathsep + session.env["PATH"]

    if not ganache_cli.exists():

        if nodeenv_dir.exists():
            shutil.rmtree(nodeenv_dir)

        session.run("nodeenv", str(nodeenv_dir))

        session.run(
            str(bindir.joinpath("npm")),
            "install",
            "-g",
            "ganache-cli@6.12.2",
            silent=True,
            external=True,
        )


def install_prettier(session: Session) -> None:
    """install prettier"""
    session.install("nodeenv")
    assert session.bin is not None
    nodeenv_dir = pathlib.Path(session.bin).parent.joinpath("node")
    bindir = nodeenv_dir.joinpath("bin").absolute()

    prettier_cli = bindir.joinpath("prettier")
    os.environ["PATH"] = str(bindir) + os.pathsep + os.environ["PATH"]
    session.env["PATH"] = str(bindir) + os.pathsep + session.env["PATH"]

    if not prettier_cli.exists():
        if nodeenv_dir.exists():
            shutil.rmtree(nodeenv_dir)

        session.run("nodeenv", str(nodeenv_dir))

        for pkg in ["prettier@2.3.1", "prettier-plugin-solidity@1.0.0-beta.13"]:
            session.run(
                str(bindir.joinpath("npm")), "install", "-g", pkg, silent=True, external=True
            )


def fix_requirements(session: Session) -> None:
    """remove extra requirements, because otherwise pip can't use this file as a constraints file
    anymore"""
    session.run("sed", "-i", r"s/\[.*\]//", "requirements.txt", external=True)


@nox.session
def update_requirements(session: Session) -> None:
    session.install("pip-tools")
    session.run("pip-compile", "requirements.in")
    fix_requirements(session)


@nox.session
def upgrade_requirements(session: Session) -> None:
    session.install("pip-tools")
    session.run("pip-compile", "-U", "requirements.in")
    fix_requirements(session)


@nox.session
def black(session: Session) -> None:
    session.install("black", *requirements_as_constraints)
    session.run("black", "--check", "--diff", *python_paths)


@nox.session
def flake8(session: Session) -> None:
    session.install("flake8", "flake8-import-order", *requirements_as_constraints)
    session.run("flake8", *python_paths)


@nox.session
def mypy(session: Session) -> None:
    session.install("mypy", *requirements_as_constraints)
    session.install("-r", "requirements.txt")
    session.run("mypy", *python_paths)


@nox.session
def prettier(session: Session) -> None:
    install_prettier(session)
    session.run("prettier", "--check", ".", external=True)


@nox.session
def test_contracts(session: Session) -> None:
    session.install("-r", "requirements.txt")
    install_ganache(session)
    session.chdir("contracts")
    session.run("brownie", "compile")
    session.run("brownie", "test")


@nox.session
def build_dapp(session: Session) -> None:
    session.install("nodeenv")
    assert session.bin is not None
    nodeenv_dir = pathlib.Path(session.bin).parent.joinpath("node")
    bindir = nodeenv_dir.joinpath("bin").absolute()

    os.environ["PATH"] = str(bindir) + os.pathsep + os.environ["PATH"]
    session.env["PATH"] = str(bindir) + os.pathsep + session.env["PATH"]
    if not nodeenv_dir.exists():
        session.run("nodeenv", str(nodeenv_dir))

    session.cd("example/")

    npm = str(bindir.joinpath("npm"))
    session.run(npm, "install", external=True, silent=True)
    session.run(npm, "run", "build", external=True, silent=True)
    session.cd("..")

    session.env["BINDIR"] = pathlib.Path("example/dist").absolute()
    session.cd("shuttermint")
    session.run("make", "wasm", external=True, silent=False)
    session.cd("..")

    # Alas, the following doesn't work
    # session.run(npm, "install", "-g", "serve", external=True, silent=True)
    # session.run("serve", "-s", "example/dist/", external=True)
