package main

import (
	r "reflect"
)

func (s Scoreboard) diff(last Scoreboard) Scoreboard {
	var out Scoreboard
	d := r.ValueOf(&out).Elem()
	c := r.ValueOf(s)
	l := r.ValueOf(last)

	for i := 0; i < c.NumField(); i++ {
		switch c.Field(i).Kind() {
		case r.Struct:
			cs := c.Field(i).Interface().(Side)
			ls := l.Field(i).Interface().(Side)
			d.Field(i).Set(r.ValueOf(cs.diff(ls)))
		default:
			if c.Field(i) != l.Field(i) {
				d.Field(i).Set(c.Field(i))
			}
		}
	}

	return out
}

func (s Scoreboard) identical(other Scoreboard) bool {
	return r.DeepEqual(s, other)
}

func (s Side) diff(last Side) Side {
	var out Side
	d := r.ValueOf(&out).Elem()
	c := r.ValueOf(s)
	l := r.ValueOf(last)

	for i := 0; i < c.NumField(); i++ {
		df := d.Field(i)
		cf := c.Field(i)
		lf := l.Field(i)

		switch cf.Kind() {
		case r.Array, r.Slice:
			// sometimes the api gives 0 players, then adds them into the game later
			// in that case, just return all of the players
			if cf.Len() == 0 || cf.Len() != lf.Len() {
				df.Set(cf)
				continue
			}
			switch cf.Index(0).Interface().(type) {
			case HeroID:
				// when length is 5, all heroes have been picked or banned
				if cf.Len() != 5 {
					df.Set(cf)
				}
			case Player:
				for j := 0; j < cf.Len(); j++ {
					cpl := cf.Index(j).Interface().(Player)
					lpl := lf.Index(j).Interface().(Player)
					df.Set(r.Append(df, r.ValueOf(cpl.diff(lpl))))
				}
			case Ability:
				for j := 0; j < cf.Len(); j++ {
					cab := cf.Index(j).Interface().(Ability)
					lab := lf.Index(j).Interface().(Ability)
					df.Set(r.Append(df, r.ValueOf(cab.diff(lab))))
				}
			}
		default:
			if cf != lf {
				df.Set(cf)
			}
		}
	}
	return out
}

func (p Player) diff(last Player) Player {
	var out Player
	d := r.ValueOf(&out).Elem()
	c := r.ValueOf(p)
	l := r.ValueOf(last)

	for i := 0; i < c.NumField(); i++ {
		if c.Field(i) != l.Field(i) {
			d.Field(i).Set(c.Field(i))
		}
	}
	return out
}

func (a Ability) diff(last Ability) Ability {
	var out Ability
	d := r.ValueOf(&out).Elem()
	c := r.ValueOf(a)
	l := r.ValueOf(last)

	for i := 0; i < c.NumField(); i++ {
		if c.Field(i) != l.Field(i) {
			d.Field(i).Set(c.Field(i))
		}
	}
	return out
}
