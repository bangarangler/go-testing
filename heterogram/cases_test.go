package heterogram

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "test empty string",
		input:       "",
		expected:    true,
	},
	{
		description: "test double hyphen",
		input:       "six-year-old",
		expected:    true,
	},
	{
		description: "test uppercase letter",
		input:       "Alpha",
		expected:    false,
	},
	{
		description: "test repeating letter",
		input:       "toto",
		expected:    false,
	},
}
