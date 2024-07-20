package views

import "embed"

//go:embed home/*
//go:embed landing/*
//go:embed *
var FS embed.FS
