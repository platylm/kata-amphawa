package sum

import "testing"

func Test_Summation_Input_2_3_4_5_Should_Be_14(t *testing.T) {
	expected := 14
	input := []int{2, 3, 4, 5}

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_Summation_Input_28_35_42_Should_Be_105(t *testing.T) {
	expected := 105
	input := []int{28, 35, 42}

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_Summation_Input_sum_Should_Be_474872753(t *testing.T) {
	expected := 474872753
	input := []int{9388053,
		9957604,
		7816221,
		3581878,
		1054154,
		7600924,
		5887511,
		1557430,
		2270,
		5014122,
		4806041,
		5188700,
		9178740,
		430827,
		2132412,
		2734167,
		9100584,
		8834925,
		4791907,
		4884140,
		5670781,
		5369137,
		4908478,
		5673518,
		9298945,
		871252,
		7679157,
		5972648,
		4938563,
		8427722,
		750335,
		5472266,
		3677089,
		4545035,
		4251796,
		170379,
		3561536,
		8567021,
		6225714,
		7931231,
		5452413,
		2145831,
		2657139,
		3507260,
		1046564,
		9704289,
		8288480,
		4795964,
		6306314,
		1518930,
		1485435,
		874336,
		1791207,
		1522134,
		3806839,
		7182804,
		8864304,
		522620,
		4155659,
		5496211,
		4887587,
		1946361,
		6806144,
		8780260,
		3803815,
		2833526,
		9214299,
		831221,
		5834346,
		2644171,
		396096,
		1020184,
		4996411,
		5516647,
		7825539,
		8802776,
		1342015,
		1138164,
		469291,
		1168329,
		5932613,
		2993371,
		3736850,
		554876,
		1451661,
		6532348,
		4260290,
		8172615,
		665706,
		3917625,
		3962113,
		6507229,
		6243422,
		5343617,
		1466969,
		2485945,
		7388936,
		6731505,
		6986394,
		7294009,
		912426,
		814490,
		4674511,
		6556104}

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}
