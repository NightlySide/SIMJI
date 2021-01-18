APP=simji
APPDIR=dist/$(APP)_1.1.0

.PHONY: makefolders copyicons packagefiles gobuild debianbuild clean postclean

.ONESHELL:
default: clean makefolders copyicons gobuild debianbuild postclean

makefolders:
	@echo "-- Making folders"
	mkdir -p $(APPDIR)/usr/bin
	mkdir -p $(APPDIR)/usr/share/applications
	mkdir -p $(APPDIR)/usr/share/icons/hicolor/1024x1024/apps
	mkdir -p $(APPDIR)/usr/share/icons/hicolor/256x256/apps
	mkdir -p $(APPDIR)/DEBIAN

copyicons:
	@echo "-- Copying icons"
	cp static/favicon.png $(APPDIR)/usr/share/icons/hicolor/1024x1024/apps/$(APP).png
	cp static/favicon.png $(APPDIR)/usr/share/icons/hicolor/256x256/apps/$(APP).png

packagefiles:
	@echo "-- Packaging files"
	go run github.com/markbates/pkger/cmd/pkger

gobuild:
	@echo "-- Building binary with GO"
	CGO_ENABLED=0 GOOS=linux go build -v -o $(APPDIR)/usr/bin/$(APP)
	ln -s $(APPDIR)/usr/bin/$(APP) ./$(APP)

debianbuild:
	@echo "-- Creating DEB package"
	cat > $(APPDIR)/usr/share/applications/$(APP).desktop << EOF
	[Desktop Entry]
	Version=1.0
	Type=Application
	Name=$(APP)
	Exec=$(APP)
	Icon=$(APP)
	Terminal=false
	StartupWMClass=Simji
	EOF

	cat > $(APPDIR)/DEBIAN/control << EOF
	Package: $(APP)
	Version: 1.0-0
	Section: base
	Priority: optional
	Architecture: amd64
	Maintainer: Alexandre FROEHLICH <nightlyside@gmail.com>
	Description: Une application SIMJI ENSTA Bretagne
	EOF

	dpkg-deb --build $(APPDIR)

clean:
	-
	@echo "-- Cleaning previous installs"
	rm -R $(APPDIR)
	rm $(APP)
	
postclean:
	-
	@echo "-- Cleaning post install"
	#rm pkger.go