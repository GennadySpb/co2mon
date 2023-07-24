#!/usr/bin/env bash
apt-get update -qq && apt-get install -y libudev-dev libhidapi-dev
dpkg -S libudev.h
