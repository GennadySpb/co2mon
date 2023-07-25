#!/usr/bin/env bash
dpkg --add-architecture amd64
dpkg --add-architecture arm64
dpkg --add-architecture armhf
apt-get update -qq && apt-get install -y libudev-dev:amd64 libudev-dev:arm64 libudev-dev:armhf
dpkg -S libudev.h
