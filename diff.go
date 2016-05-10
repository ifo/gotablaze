package main

import (
	"reflect"
)

func diffScoreboard(last, current Scoreboard) (out Scoreboard) {
	diff := reflect.ValueOf(&out)
	lastVal, currVal := reflect.ValueOf(last), reflect.ValueOf(current)
	for i := 0; i < lastVal.NumField(); i++ {
		switch lastVal.Field(i).Kind() {
		case reflect.Struct:
			if lastVal.Field(i).Type() == reflect.TypeOf(Side{}) {
				diff.Elem().Field(i).Set(
					reflect.ValueOf(diffSide(
						lastVal.Field(i).Interface().(Side),
						currVal.Field(i).Interface().(Side))))
			}
		default:
			if lastVal.Field(i) != currVal.Field(i) {
				diff.Elem().Field(i).Set(currVal.Field(i))
			}
		}
	}
	return
}

func areIdenticalScoreboard(one, two Scoreboard) bool {
	return reflect.DeepEqual(one, two)
}

func diffSide(last, current Side) Side {
	out := Side{}
	diff := reflect.ValueOf(&out)
	lastVal, currVal := reflect.ValueOf(last), reflect.ValueOf(current)
	for i := 0; i < lastVal.NumField(); i++ {
		switch lastVal.Field(i).Kind() {
		case reflect.Array, reflect.Slice:
			// sometimes the api gives 0 players, then adds them into the game later
			// in that case, just return all of the players
			if lastVal.Field(i).Len() == 0 || lastVal.Field(i).Len() != currVal.Field(i).Len() {
				diff.Elem().Field(i).Set(currVal.Field(i))
				continue
			}
			switch lastVal.Field(i).Index(0).Interface().(type) {
			case HeroID:
				// when length is 5, all heroes have been picked or banned
				if lastVal.Field(i).Len() != 5 {
					diff.Elem().Field(i).Set(currVal.Field(i))
				}
			case Player:
				for j := 0; j < lastVal.Field(i).Len(); j++ {
					diff.Elem().Field(i).Set(
						reflect.Append(diff.Elem().Field(i),
							reflect.ValueOf(diffPlayer(
								lastVal.Field(i).Index(j).Interface().(Player),
								currVal.Field(i).Index(j).Interface().(Player)))))
				}
			case Ability:
				for j := 0; j < lastVal.Field(i).Len(); j++ {
					diff.Elem().Field(i).Set(
						reflect.Append(diff.Elem().Field(i),
							reflect.ValueOf(diffAbility(
								lastVal.Field(i).Index(j).Interface().(Ability),
								currVal.Field(i).Index(j).Interface().(Ability)))))
				}
			}
		default:
			if lastVal.Field(i) != currVal.Field(i) {
				diff.Elem().Field(i).Set(currVal.Field(i))
			}
		}
	}
	return out
}

func diffPlayer(last, current Player) (out Player) {
	diff := reflect.ValueOf(&out)
	lastVal, currVal := reflect.ValueOf(last), reflect.ValueOf(current)
	for i := 0; i < lastVal.NumField(); i++ {
		if lastVal.Field(i) != currVal.Field(i) {
			diff.Elem().Field(i).Set(currVal.Field(i))
		}
	}
	return
}

func diffAbility(last, current Ability) (out Ability) {
	diff := reflect.ValueOf(&out)
	lastVal, currVal := reflect.ValueOf(last), reflect.ValueOf(current)
	for i := 0; i < lastVal.NumField(); i++ {
		if lastVal.Field(i) != currVal.Field(i) {
			diff.Elem().Field(i).Set(currVal.Field(i))
		}
	}
	return
}
