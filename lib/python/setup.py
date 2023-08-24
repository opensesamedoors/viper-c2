from setuptools import setup, find_packages

VERSION = "0.0.3" 
DESCRIPTION = "Viper API libary"

setup(
    name="viper", 
    version=VERSION,
    author="ngn",
    description=DESCRIPTION,
    packages=find_packages(),
    install_requires=["requests"] 
)
