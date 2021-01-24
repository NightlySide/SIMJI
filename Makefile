APP=simji
APPDIR=pkg/$(APP)_1.2.0

.PHONY: test coverage makefolders copyicons packagefiles gobuild debianbuild clean postclean

.ONESHELL:
default: clean makefolders copyicons packagefiles gobuild debianbuild postclean

makefolders:
	@echo "-- Making folders"
	mkdir -p $(APPDIR)/usr/bin
	mkdir -p $(APPDIR)/usr/share/applications
	mkdir -p $(APPDIR)/usr/share/icons/hicolor/1024x1024/apps
	mkdir -p $(APPDIR)/usr/share/icons/hicolor/256x256/apps
	mkdir -p $(APPDIR)/DEBIAN

copyicons:
	@echo "-- Copying icons"
	cp internal/static/favicon.png $(APPDIR)/usr/share/icons/hicolor/1024x1024/apps/$(APP).png
	cp internal/static/favicon.png $(APPDIR)/usr/share/icons/hicolor/256x256/apps/$(APP).png

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
	rm pkged.go

test:
	@echo "-- Running tests"
	go run github.com/rakyll/gotest simji/internal/assembler
	go run github.com/rakyll/gotest simji/internal/vm
	go run github.com/rakyll/gotest simji/internal/log

coverage:
	@echo "-- Running test coverage"
	go run github.com/rakyll/gotest --cover simji/internal/assembler
	go run github.com/rakyll/gotest --cover simji/internal/vm
	go run github.com/rakyll/gotest --cover simji/internal/log

coverage-display:
	@echo "-- Running test coverage and display results"
	go run github.com/rakyll/gotest -coverprofile=cov.out ./...
	go tool cover -html cov.out
	rm cov.out