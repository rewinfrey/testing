package test

func Test_unmarshal(t *testing.T) {
	for _, testCase := range unmarshalCases {
	}
	for i, testCase := range unmarshalCases {
		t.Run(fmt.Sprintf("[%v]%s", i, testCase.input), func(t *testing.T) {
		})
	}
}
