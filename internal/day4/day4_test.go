package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

const passports = `byr:1937 iyr:2017 eyr:2020 hgt:183cm
hcl:#fffffd ecl:gry pid:860033327 cid:147

byr:1931 iyr:2013 eyr:2024 hgt:179cm 
hcl:#ae17e1 ecl:brn pid:760753108

byr:abc iyr:2018 eyr:2024 hgt:172cm
hcl:#602927 ecl:blu pid:839621424

byr:1919 iyr:2017 eyr:2028 hgt:67in
hcl:#fffffd ecl:gry pid:459927933

byr:2003 iyr:2019 eyr:2030 hgt:163cm
hcl:#a97842 ecl:amb pid:472877556 cid:322

byr:2000 iyr:abc eyr:2022 hgt:180cm
hcl:#733820 ecl:brn pid:751634349 cid:320

byr:1921 iyr:2009 eyr:2030 hgt:161cm
hcl:#4dd6d4 ecl:grn pid:258660154 cid:217

byr:1978 iyr:2021 eyr:2022 hgt:170cm
hcl:#602927 ecl:blu pid:399347273 cid:109

byr:1980 iyr:2020 eyr:abc hgt:183cm
hcl:#18171d ecl:gry pid:172106685 cid:289

byr:1955 iyr:2014 eyr:2019 hgt:187cm
hcl:#623a2f ecl:brn pid:008305281 cid:74

byr:1960 iyr:2016 eyr:2031 hgt:167cm
hcl:#623a2f ecl:grn pid:428624233

byr:1957 iyr:2018 eyr:2027 hgt:abccm
hcl:#ceb3a1 ecl:oth pid:358876826 cid:314

byr:1988 iyr:2015 eyr:2025 hgt:149cm
hcl:#6b5442 ecl:brn pid:899417027 cid:268

byr:1995 iyr:2020 eyr:2028 hgt:194cm
hcl:#b6652a ecl:hzl pid:594197202

byr:1930 iyr:2010 eyr:2020 hgt:abcin
hcl:#a97842 ecl:brn pid:010268954

byr:1935 iyr:2016 eyr:2028 hgt:58in
hcl:#6b5442 ecl:blu pid:187679418

byr:1930 iyr:2019 eyr:2025 hgt:77in
hcl:#888785 ecl:oth pid:704379775

byr:1945 iyr:2014 eyr:2022 hgt:123
hcl:#7d3b0c ecl:hzl pid:011589584

byr:1950 iyr:2010 eyr:2028 hgt:176cm
hcl:18171d ecl:grn pid:685748669

byr:1989 iyr:2020 eyr:2020 hgt:163cm
hcl:#18171d ecl:abc pid:721397788 cid:308

byr:1951 iyr:2012 eyr:2029 hgt:186cm
hcl:#fffffd ecl:amb pid:abc

byr:1923 iyr:2019 eyr:2025 hgt:167cm
hcl:#341e13 ecl:amb pid:12345678 cid:102`

func TestSolve(t *testing.T) {
	stream := input.FromLiteral(passports)

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 22, solution1, "Solution 1 should be 22")
	assert.Equal(t, 2, solution2, "Solution 2 should be 2")
	assert.Nil(t, err, "Error should be nil")
}
