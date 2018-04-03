# carrot

### Example

``` go
func main() {
	m := carrot.GetWeights()
	m.Add(12, "BonusGame")
	m.Add(25, "FreeGame")
	m.Add(600, "BaseGame")

	c := carrot.New()
	c.Add("gamemode", m)
	c.Run("gamemode")
	// Run Count: 1,000,000 BaseGame: 94.197%, FreeGame: 3.907%, BonusGame: 1.89%
}
```