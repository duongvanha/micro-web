#!/bin/sh

cp -ru /build-dir/node_modules/ /usr/src/app/

exec npm run serve
