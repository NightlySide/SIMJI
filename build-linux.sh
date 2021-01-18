#!/bin/sh

APP=simji
APPDIR=dist/${APP}_1.0.0

mkdir -p $APPDIR/usr/bin
mkdir -p $APPDIR/usr/share/applications
mkdir -p $APPDIR/usr/share/icons/hicolor/1024x1024/apps
mkdir -p $APPDIR/usr/share/icons/hicolor/256x256/apps
mkdir -p $APPDIR/DEBIAN

CGO_ENABLED=0 go build -o $APPDIR/usr/bin/$APP

cp static/favicon.png $APPDIR/usr/share/icons/hicolor/1024x1024/apps/${APP}.png
cp static/favicon.png $APPDIR/usr/share/icons/hicolor/256x256/apps/${APP}.png

cat > $APPDIR/usr/share/applications/${APP}.desktop << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=$APP
Exec=$APP
Icon=$APP
Terminal=false
StartupWMClass=Lorca
EOF

cat > $APPDIR/DEBIAN/control << EOF
Package: ${APP}
Version: 1.0-0
Section: base
Priority: optional
Architecture: amd64
Maintainer: Alexandre FROEHLICH <nightlyside@gmail.com>
Description: Une application SIMJI ENSTA Bretagne
EOF

dpkg-deb --build $APPDIR
