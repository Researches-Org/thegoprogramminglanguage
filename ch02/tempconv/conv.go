package tempconv

// CToF converts a Celsius temperatue to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
  return Fahrenheit(c*9/5 + 32)
}

// FtoC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
  return Celsius((f - 32) * 5 / 9)
}

// KToC converts a Kelvin temperatute to Celsius. 
func KToC(k Kelvin) Celsius {
  return Celsius(k - 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
  return CToF(KToC(k))
}

