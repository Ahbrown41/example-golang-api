#
# (c)  2021 – Wawa, Inc.
# All Rights Reserved.
#
# NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
# All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
# Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
#

export VERSION=`cat ./VERSION | xargs`

echo "Building Version: ${VERSION}"

go build -ldflags "-X main.Version=${VERSION}" .

echo "Create Tag: ${VERSION}"
git tag "${VERSION}"
git push origin --tags

zip -r "api-reference-golang-${VERSION}.zip" ./api-reference-golang.exe ./config.yaml ./README.md ./docs

echo "Created Package: api-reference-golang-${VERSION}.zip"