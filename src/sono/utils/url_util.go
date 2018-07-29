package utils

import (
    "regexp"
    "sono/log"
)

func IsNormal(url string) bool {
    isNormalReg, err := regexp.MatchString("/sono/(\\w{8}(-\\w{4}){3}-\\w{12}?)/normal$", url)
    if err != nil {
        log.Error("match normal url error")
        return false
    }
    return isNormalReg
}

func IsTest(url string) bool {
    isTestReg, err := regexp.MatchString("/sono/(\\w{8}(-\\w{4}){3}-\\w{12}?)/test", url)
    if err != nil {
        log.Error("match test url error")
        return false
    }
    return isTestReg
}

func ExtractUuid(url string) string {
    reg := regexp.MustCompile("(\\w{8}(-\\w{4}){3}-\\w{12}?)")
    res := reg.Find([]byte(url))
    return string(res)
}
