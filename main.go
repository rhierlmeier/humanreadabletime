/**
 * Copyright (c) 2023 rhierlmeier
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache Public License 2.0 which is available at
 *  http://www.apache.org/licenses/
 *
 * SPDX-License-Identifier: APSL-2.0
 */
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func getFactorFromUnit(unit string) (float64, error) {
	switch unit {
	case "":
		return 1, nil
	case "ms":
		return 1, nil
	case "s":
		return 1000, nil
	case "m":
		return 1000 * 60, nil
	case "h":
		return 1000 * 60 * 60, nil
	case "d":
		return 1000 * 60 * 60 * 24, nil
	default:
		return 0, errors.New("Unknown unit[" + unit + "], supported units are s, m, h or d")
	}
}

func toHumanReadableString(ms float64) string {

	x := ms / 1000
	seconds := int64(math.Mod(x, 60))
	x /= 60
	minutes := int64(math.Mod(x, 60))
	x /= 60
	hours := int64(math.Floor(math.Mod(x, 24)))
	days := int64(math.Floor(x / 24))

	ret := ""
	if days > 0 {
		ret += strconv.FormatInt(days, 10)
		if days > 1 {
			ret += " Tage"
		} else {
			ret += " Tag"
		}
	}
	if hours > 0 {
		if len(ret) > 0 {
			ret += " "
		}
		ret += strconv.FormatInt(hours, 10)
		ret += " Std."
	}
	if minutes > 0 {
		if len(ret) > 0 {
			ret += " "
		}
		ret += strconv.FormatInt(minutes, 10)
		ret += " Min."
	}

	if minutes > 0 {
		if len(ret) > 0 {
			ret += " "
		}
		ret += strconv.FormatInt(seconds, 10)
		ret += "s"
	}
	return ret
}

func main() {

	re := regexp.MustCompile(`(\d+(?:\.\d+)?)(ms|s|m|h|d)?`)

	parts := re.FindStringSubmatch(os.Args[1])

	unit := ""
	numStr := ""
	switch len(parts) {
	case 3:
		{
			numStr = parts[1]
			unit = parts[2]
		}
	default:
		{
			fmt.Println("Error: Invalid arguments")
			os.Exit(1)
		}
	}

	factor, err := getFactorFromUnit(unit)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Error: Missing argument")
		os.Exit(1)
	}

	t := float64(num) * float64(factor)

	fmt.Println(toHumanReadableString(t))
}
