package utils

import "strconv"

// StringToInt converts a string to an integer with a default value
func StringToInt(value string, defaultValue int) int {
    if value == "" {
        return defaultValue
    }
    
    result, err := strconv.Atoi(value)
    if err != nil {
        return defaultValue
    }
    
    return result
}