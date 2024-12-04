# -*- coding: utf-8 -*-
from setuptools import setup
from pathlib import Path

setup_kwargs = {
    'name': 'unitsnet-generator',
    'license': 'MIT',
    'author': 'Haim Kastner',
    'author_email': 'hello@haim-kastner.com',
    'maintainer': 'Haim Kastner',
    'maintainer_email': 'hello@haim-kastner.com',
    'url': 'https://github.com/haimkastner/unitsnet-go',
    'python_requires': '>=3.8,<4.0',
    'install_requires': [],
}

setup(**setup_kwargs)
