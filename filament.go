package gormfilament

import (
	"gopkg.in/go-playground/colors.v1"
	"gorm.io/gorm"
)

type Filament struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Material string `gorm:"not null"`
	Color    string `gorm:"not null"`
	Weight   uint   `gorm:"not null"`
	Slot     uint   `gorm:"unique;not null;check:slot > 0"`
	Price    uint   `gorm:"check:price >= 0"`
}

func CreateFilamentRGBRaw(name string, material string, r uint8, g uint8, b uint8, weight uint, slot uint, price uint) Filament {
	var colorRGB = colors.RGBColor{R: r, G: g, B: b}
	var spoolColor string = colorRGB.ToHEX().String()

	Filament := Filament{Name: name, Material: material, Color: spoolColor, Weight: weight, Slot: slot}

	return Filament
}

func CreateFilamentRGB(name string, material string, color colors.RGBColor, weight uint, slot uint, price uint) Filament {
	var spoolColor string = color.ToHEX().String()
	Filament := Filament{Name: name, Material: material, Color: spoolColor, Weight: weight, Slot: slot}

	return Filament
}

func CreateFilamentRGBHex(name string, material string, color string, weight uint, slot uint, price uint) Filament {
	spoolColor, err := colors.ParseHEX(color)
	if err != nil {
		spoolColor, _ = colors.ParseHEX("#000000") //Black
	}
	Filament := Filament{Name: name, Material: material, Color: spoolColor.String(), Weight: weight, Slot: slot}

	return Filament
}
