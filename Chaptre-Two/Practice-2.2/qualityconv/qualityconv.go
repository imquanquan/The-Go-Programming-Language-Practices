package qualityconv

import "fmt"

type Pound float64
type Kilogram float64

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

func (p Pound) String() string { return fmt.Sprintf("%g lb", p) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.4536) }

func KToP(k Kilogram) Pound { return Pound(k / 0.4536) }