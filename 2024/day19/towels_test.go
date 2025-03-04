package day19

import "testing"

func TestCountPossible(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"r, wr, b, g, bwu, rb, gb, br",
				"",
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			expected: 6,
		},
		{
			inputs: []string{
				"gwurrw, guuu, bgguuu, gwrg, rwrg, gugbw, gubwgb, bubrugr, ggugww, urr, wgub, bbbgbuuu, wuwrru, gububbr, rrr, bbubrub, urbgwb, wub, gwguu, urbrw, rgwrwg, uuuw, ggrg, ggbbwg, gbug, gruw, uuugw, wrg, urg, ggbr, rbwr, wbwb, brrwwu, guw, ruu, gwrgru, uwbg, gbrrg, ugrb, wgr, bgbrg, bgw, brg, wrgr, wbwg, uwrru, grgr, bgu, uwwurgw, ubbu, g, rbburugu, www, gwgbbbrw, grgugu, bgug, burur, wuu, uwurbg, rrug, wwwburg, wwgbr, ugrr, gurwrbrb, gr, rwru, burrg, rb, wgggw, rru, urubgbr, bb, rrbuuuw, wbruruu, rgu, ggb, u, urw, ubu, grurrru, uwr, gwbuuu, bwwg, gbgg, uuwbrrg, gbr, bbub, bru, uru, wrwrg, gbbbubwg, gurbuu, bbbuwbr, uuwu, uggrur, guu, rwwgr, ggu, gbu, rrwrgrb, ww, buu, gru, urggwbb, wrggbug, gugg, wubb, uwrbb, gguurb, rgb, w, rgbr, bbb, grbbuw, ruguugg, bub, ugu, rwuur, rrbg, wruburu, wwubu, ugwrbbub, bbggw, wrrb, wbuu, wbr, bwwug, rbgwr, buuw, urb, rgr, ug, ubru, buurr, bbw, rgbb, bbr, gwbw, rbgg, uu, grwrgbwr, wubw, gruwrg, bgrw, wgg, uwb, urguwbb, bwr, rbuugw, wbg, wrrr, uggwu, brb, rguw, bubwuw, rbg, ubw, bugwu, ruurr, ggr, bwuu, rrbr, bg, rrubr, ruw, wru, ugg, gurur, rbug, rwrgw, uwub, rgrgg, gw, gwurr, wug, ubrbu, gbru, grr, bgbbrb, brgw, rrrg, bubr, wbu, bwu, wbrrgw, ruurgbgg, rwgwu, wwur, wwwb, rrg, brbgg, rbb, rrruguwu, rw, bwru, gg, wwr, rwgwwgb, ubur, wwrwug, gruurrw, wgbugbu, grwwgg, rbu, gww, ggg, ubb, bwrub, rrw, gwgr, rurgwrub, ubbb, uuww, bgr, bguwbw, brw, burw, guuuu, wu, rwg, ubbrr, wgwrwub, wguuw, wbb, uub, gwbbb, ugb, gurug, wgwbw, ubggw, rbr, urwww, br, uww, wubub, guwub, grgg, gbbrg, bugwrrg, bur, gbbbbr, wrr, ur, wbbbr, bbrrrww, rrwbg, ubg, wugg, rwgbu, gbw, gggwbwu, rwb, bbu, rgw, wwgrwr, rrurrbr, buw, gub, gwr, urrb, rguuw, rrrbwwg, uuuu, bgrr, uwu, wbw, gwg, bwgwbr, guug, rrru, brr, bwuuwrr, wwg, uwg, ru, rbubuur, wuug, wwbburgr, guwuru, wwb, rguu, bwrg, uwwgwgwg, gububwbu, rggu, gggwru, rurrbw, wgrg, gggbb, bbwbu, wrrgu, rbugu, gu, rggb, bbuu, gbguw, uwugw, ruuwb, wgrwr, gwu, ugww, urrru, rgbbw, wgw, rgwbbgw, bu, uuuuuwgw, wrgbuuwb, ubruwgu, wwwrwb, uugwb, rrgrru, rug, wrbbwg, uguuwww, wgb, rwwu, wgu, guguu, gwrb, rbwgr, rrub, wbrrgbg, wbrurr, gbww, gubb, wbbuw, rwbw, grgbg, rrrgggr, rubruw, guggw, bwg, wggguw, ggwww, rbw, gbb, ugr, wuw, rww, gur, ubbbgr, bgg, rrbguu, bugr, wrw, ub, ugrubur, uurbb, rggub, r, rwbubw, bwb, gwbwu, rwr, wbbrwu, uwubwg, bggr, ggw, uurbr, wrb, ugbw, ubgggw, grb, wubruw, rggggwb, bbuugb, rr, urubb, gb, gbwr, rbww, rur, bwbrrb, gbuurr, gbwwgbg, rgbg, wgwb, uw, uur, rub, rgubgr, rbgggg, bbuwu, wrwbbg, ubr, brgugu, wbbwwuuw, wuuggg, gbg, gurb, rbbggw, uwwb, grbwww, bbbggbr, gugww, wb, uwrww, bgb, uug, wrbbrw, uwgu, bgbbb, grw, bwur, ruug, ugw, grg, uggbrw, bruwr, rgrg, bbuwubb, bug, wrbg, ruww, gwbubb, gug, rgbuw, rg, bbuwrgu, grgrrr, urwr, bgurbb, ugbuu, gwb, grwb, wwu, rggruwu, wurgg, uuw, uwrg, wg, wbrbr, uubgg, wwgw, ruwbr, rwgr, ugwg, rggwuwbg, urgurguw, uurw, rgrb, bbg, rurwug, ggbwgugg, grbwuwgw, rgg",
				"",
				"buwwugwwurbgrrbgrbrubwu",
			},
			expected: 1,
		},
	}
	for _, c := range cases {
		result := CountPossible(c.inputs)
		if result != c.expected {
			t.Errorf("CountPossible(%q) == %d, expected %d",
				c.inputs,
				result,
				c.expected,
			)
		}
	}
}

func TestSumCombinations(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"r, wr, b, g, bwu, rb, gb, br",
				"",
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			expected: 16,
		},
		{
			inputs: []string{
				"gwurrw, guuu, bgguuu, gwrg, rwrg, gugbw, gubwgb, bubrugr, ggugww, urr, wgub, bbbgbuuu, wuwrru, gububbr, rrr, bbubrub, urbgwb, wub, gwguu, urbrw, rgwrwg, uuuw, ggrg, ggbbwg, gbug, gruw, uuugw, wrg, urg, ggbr, rbwr, wbwb, brrwwu, guw, ruu, gwrgru, uwbg, gbrrg, ugrb, wgr, bgbrg, bgw, brg, wrgr, wbwg, uwrru, grgr, bgu, uwwurgw, ubbu, g, rbburugu, www, gwgbbbrw, grgugu, bgug, burur, wuu, uwurbg, rrug, wwwburg, wwgbr, ugrr, gurwrbrb, gr, rwru, burrg, rb, wgggw, rru, urubgbr, bb, rrbuuuw, wbruruu, rgu, ggb, u, urw, ubu, grurrru, uwr, gwbuuu, bwwg, gbgg, uuwbrrg, gbr, bbub, bru, uru, wrwrg, gbbbubwg, gurbuu, bbbuwbr, uuwu, uggrur, guu, rwwgr, ggu, gbu, rrwrgrb, ww, buu, gru, urggwbb, wrggbug, gugg, wubb, uwrbb, gguurb, rgb, w, rgbr, bbb, grbbuw, ruguugg, bub, ugu, rwuur, rrbg, wruburu, wwubu, ugwrbbub, bbggw, wrrb, wbuu, wbr, bwwug, rbgwr, buuw, urb, rgr, ug, ubru, buurr, bbw, rgbb, bbr, gwbw, rbgg, uu, grwrgbwr, wubw, gruwrg, bgrw, wgg, uwb, urguwbb, bwr, rbuugw, wbg, wrrr, uggwu, brb, rguw, bubwuw, rbg, ubw, bugwu, ruurr, ggr, bwuu, rrbr, bg, rrubr, ruw, wru, ugg, gurur, rbug, rwrgw, uwub, rgrgg, gw, gwurr, wug, ubrbu, gbru, grr, bgbbrb, brgw, rrrg, bubr, wbu, bwu, wbrrgw, ruurgbgg, rwgwu, wwur, wwwb, rrg, brbgg, rbb, rrruguwu, rw, bwru, gg, wwr, rwgwwgb, ubur, wwrwug, gruurrw, wgbugbu, grwwgg, rbu, gww, ggg, ubb, bwrub, rrw, gwgr, rurgwrub, ubbb, uuww, bgr, bguwbw, brw, burw, guuuu, wu, rwg, ubbrr, wgwrwub, wguuw, wbb, uub, gwbbb, ugb, gurug, wgwbw, ubggw, rbr, urwww, br, uww, wubub, guwub, grgg, gbbrg, bugwrrg, bur, gbbbbr, wrr, ur, wbbbr, bbrrrww, rrwbg, ubg, wugg, rwgbu, gbw, gggwbwu, rwb, bbu, rgw, wwgrwr, rrurrbr, buw, gub, gwr, urrb, rguuw, rrrbwwg, uuuu, bgrr, uwu, wbw, gwg, bwgwbr, guug, rrru, brr, bwuuwrr, wwg, uwg, ru, rbubuur, wuug, wwbburgr, guwuru, wwb, rguu, bwrg, uwwgwgwg, gububwbu, rggu, gggwru, rurrbw, wgrg, gggbb, bbwbu, wrrgu, rbugu, gu, rggb, bbuu, gbguw, uwugw, ruuwb, wgrwr, gwu, ugww, urrru, rgbbw, wgw, rgwbbgw, bu, uuuuuwgw, wrgbuuwb, ubruwgu, wwwrwb, uugwb, rrgrru, rug, wrbbwg, uguuwww, wgb, rwwu, wgu, guguu, gwrb, rbwgr, rrub, wbrrgbg, wbrurr, gbww, gubb, wbbuw, rwbw, grgbg, rrrgggr, rubruw, guggw, bwg, wggguw, ggwww, rbw, gbb, ugr, wuw, rww, gur, ubbbgr, bgg, rrbguu, bugr, wrw, ub, ugrubur, uurbb, rggub, r, rwbubw, bwb, gwbwu, rwr, wbbrwu, uwubwg, bggr, ggw, uurbr, wrb, ugbw, ubgggw, grb, wubruw, rggggwb, bbuugb, rr, urubb, gb, gbwr, rbww, rur, bwbrrb, gbuurr, gbwwgbg, rgbg, wgwb, uw, uur, rub, rgubgr, rbgggg, bbuwu, wrwbbg, ubr, brgugu, wbbwwuuw, wuuggg, gbg, gurb, rbbggw, uwwb, grbwww, bbbggbr, gugww, wb, uwrww, bgb, uug, wrbbrw, uwgu, bgbbb, grw, bwur, ruug, ugw, grg, uggbrw, bruwr, rgrg, bbuwubb, bug, wrbg, ruww, gwbubb, gug, rgbuw, rg, bbuwrgu, grgrrr, urwr, bgurbb, ugbuu, gwb, grwb, wwu, rggruwu, wurgg, uuw, uwrg, wg, wbrbr, uubgg, wwgw, ruwbr, rwgr, ugwg, rggwuwbg, urgurguw, uurw, rgrb, bbg, rurwug, ggbwgugg, grbwuwgw, rgg",
				"",
				"buwwugwwurbgrrbgrbrubwu",
			},
			expected: 51046,
		},
	}
	for _, c := range cases {
		result := SumCombinations(c.inputs)
		if result != c.expected {
			t.Errorf("SumCombinations(%q) == %d, expected %d",
				c.inputs,
				result,
				c.expected,
			)
		}
	}
}
