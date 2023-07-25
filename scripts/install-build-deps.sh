#!/usr/bin/env bash
dpkg --add-architecture amd64 armhf
apt-get update -qq && apt-get install -y libudev-dev libudev-dev:amd64 libudev-dev:armhf
dpkg -S libudev.h
