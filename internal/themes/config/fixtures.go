package theme_config

var ThemeConfigYamlGood01 = []byte(`
name: Test01
slug: test01
`)

var ThemeConfigYamlGood02 = []byte(`
name: Test02
slug: test02
description: Test02
author: Test02
`)

var ThemeConfigModelGood01 = ThemeConfig{
	Name: "Test01",
	Slug: "test01",
}

var ThemeConfigModelGood02 = ThemeConfig{
	Name:        "Test02",
	Slug:        "test02",
	Description: "Test02",
	Author:      "Test02",
}
