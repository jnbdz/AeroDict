# Makefile

include config.mk

install:
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 aerodict $(DESTDIR)$(BINDIR)
	install -d $(DESTDIR)$(MANDIR)
	install -m 644 aerodict.1 $(DESTDIR)$(MANDIR)
	install -d $(DESTDIR)$(DATADIR)
	cp -r config $(DESTDIR)$(DATADIR)

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/aerodict
	rm -f $(DESTDIR)$(MANDIR)/aerodict.1
	rm -rf $(DESTDIR)$(DATADIR)

.PHONY: install uninstall
