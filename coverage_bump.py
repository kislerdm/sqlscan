#! /usr/local/bin/python3

# Copyright 2021 dkisler.com Dmitry Kisler

"""Tool to patch the README.md with the code coderage badde
https://shields.io/ is used to generate badges
"""

import logging
import os
import re
from pathlib import Path
from typing import Union


def read(path: Union[Path, str]) -> str:
    """Read file."""
    with open(path, "r") as fread:
        return fread.read()


def write(obj: str, path: Union[Path, str]) -> None:
    """Read file."""
    with open(path, "w") as fwrite:
        fwrite.write(obj)


def execute_cmd(cmd: str) -> None:
    """Run system command."""
    os.system(cmd)


def run_gocover(path: Path) -> None:
    """Run gocover."""
    execute_cmd(
        f"""go test -tags test -coverprofile=/tmp/go-cover.tmp ./... > /dev/null
go tool cover -func /tmp/go-cover.tmp -o {path} && rm /tmp/go-cover.tmp"""
    )


def extract_total_coverage(raw: str) -> int:
    """Extract total coverage."""
    tail_line = raw.splitlines()[-1]
    return int(float(tail_line.split("\t")[-1][:-1]))


def generate_url(coverage_pct: float) -> str:
    """Generate badge source URL."""
    color = "yellow"
    if coverage_pct == 100:
        color = "brightgreen"
    elif coverage_pct > 90:
        color = "green"
    elif coverage_pct > 70:
        color = "yellowgreen"
    elif coverage_pct > 50:
        color = "yellow"
    else:
        color = "orange"

    return f"https://img.shields.io/badge/coverage-{coverage_pct}%25-{color}"


def main() -> None:
    """Run."""
    root = Path(__file__).absolute().parents[1]
    path_readme = root / "README.md"
    path_coverage = root / "COVERAGE"
    placeholder_tag = "Code Coverage"
    regexp_pattern = rf"\[\!\[{placeholder_tag}\]\(.*\)\]\(.*\)"

    run_gocover(path_coverage)

    coverage = read(path_coverage)

    coverage_pct = extract_total_coverage(coverage)

    badge_url = generate_url(coverage_pct)

    inpt = read(path_readme)

    search = re.findall(regexp_pattern, inpt)

    if not search:
        raise Exception(f"No placeholder found in README.md. Add '[![{placeholder_tag}]()]()'.")

    placeholder_inject = f"[![{placeholder_tag}]({badge_url})]({badge_url})"

    out = re.sub(regexp_pattern, placeholder_inject, inpt)

    write(out, path_readme)


if __name__ == "__main__":
    log = logging.getLogger("coverage-bump")

    try:
        main()
    except Exception as ex:  # pylint: disable=broad-except
        log.error(ex)
