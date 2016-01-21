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
			if lastVal.Field(i).Interface() != currVal.Field(i).Interface() {
				diff.Elem().Field(i).Set(currVal.Field(i))
			}
		}
	}
	return
}

func areIdenticalScoreboard(one, two Scoreboard) bool {
	return reflect.DeepEqual(one, two)
}

func diffSide(last, current Side) (out Side) {
	diff := reflect.ValueOf(&out)
	lastVal, currVal := reflect.ValueOf(last), reflect.ValueOf(current)
	for i := 0; i < lastVal.NumField(); i++ {
		switch lastVal.Field(i).Kind() {
		case reflect.Array, reflect.Slice:
			for j := 0; j < lastVal.Field(i).Len(); j++ {
				diff.Elem().Field(i).Set(
					reflect.Append(diff.Elem().Field(i),
						reflect.ValueOf(diffPlayer(
							lastVal.Field(i).Index(j).Interface().(Player),
							currVal.Field(i).Index(j).Interface().(Player)))))
			}
		default:
			if lastVal.Field(i).Interface() != currVal.Field(i).Interface() {
				diff.Elem().Field(i).Set(currVal.Field(i))
			}
		}
	}
	return
}

func diffPlayer(last, current Player) (out Player) {
	diff := reflect.ValueOf(&out)
	lastVal, currVal := reflect.ValueOf(last), reflect.ValueOf(current)
	for i := 0; i < lastVal.NumField(); i++ {
		if lastVal.Field(i).Interface() != currVal.Field(i).Interface() {
			diff.Elem().Field(i).Set(currVal.Field(i))
		}
	}
	return out
}
