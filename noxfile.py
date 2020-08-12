if __name__ == "__main__":
    __import__("sys").exit(
        "Do not execute this file directly. Use nox instead, it will know how to handle this file"
    )

import os
import pathlib
import shutil

import nox
from nox.sessions import Session


NODE_VERSION = "12.18.0"


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

        session.run("nodeenv", "--node", NODE_VERSION, str(nodeenv_dir))

        session.run(
            str(bindir.joinpath("npm")),
            "install",
            "-g",
            "ganache-cli@6.9.1",
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

        session.run("nodeenv", "--node", NODE_VERSION, str(nodeenv_dir))

        for pkg in ["prettier@2.0.5", "prettier-plugin-solidity@1.0.0-alpha.54"]:
            session.run(
                str(bindir.joinpath("npm")), "install", "-g", pkg, silent=True, external=True,
            )


@nox.session
def update_requirements(session: Session) -> None:
    session.install("pip-tools")
    session.run("pip-compile", "requirements.in")


@nox.session
def upgrade_requirements(session: Session) -> None:
    session.install("pip-tools")
    session.run("pip-compile", "-U", "requirements.in")


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
    session.run("prettier", "--check", "contracts/contracts", external=True)


@nox.session
def test_contracts(session: Session) -> None:
    session.install("-r", "requirements.txt")
    install_ganache(session)
    session.chdir("contracts")

    session.run("brownie", "test")
