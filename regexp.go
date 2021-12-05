package fsspgoparser

import "regexp"

var (
	personName = regexp.MustCompile(`^([А-ЯЁ]+) ([А-ЯЁ]+) ([А-ЯЁ]+) *([\d.]*) *([А-Я.,\s-]*?)$`)
	legalName  = regexp.MustCompile(`^(.+?),(.+)$`)
	production = regexp.MustCompile(`^(.+-ИП) от (\d{2}\.\d{2}\.\d{4}) *(.*)$`)
	document   = regexp.MustCompile(`^(.+) от (\d{2}\.\d{2}\.\d{4}) № ([А-Я ]{0,3}[№ ]*[\d\-\/А-Я,]+) (.+)$`)
	subject    = regexp.MustCompile(`(.+?): ([\d.]+) руб\.`)
	department = regexp.MustCompile(`^(.+?) \d`)
	bailiff    = regexp.MustCompile(`([А-ЯЁ]+) ([А-ЯЁ]\.) ([А-ЯЁ]\.)<br>(\+\d+)`)
	ipEnd      = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}), (\d{2}), (\d), (\d)`)
)
