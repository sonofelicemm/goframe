package utils

import (
    "testing"
)

func TestIsNormal(t *testing.T) {
    isInput := IsNormal("/sono/8a8eb3a3-ead2-4aa6-be4c-f60c4278ed7c/normal")
    if !isInput {
        t.Error("wrong")
    }
}

func TestIsTest(t *testing.T) {
    isTest := IsTest("/sono/081a2b14-deef-43cb-bcb7-342be55a9b6/test")
    if !isTest {
        t.Error("wrong")
    }
}

func TestUrlUtilExtractUuid(t *testing.T) {
    uuid := ExtractUuid("/sono/081a2b14-deef-43cb-bcb7-342be55a9b6c/test")
    if uuid != "081a2b14-deef-43cb-bcb7-342be55a9b6c" {
        t.Error("wrong")
    }
}
